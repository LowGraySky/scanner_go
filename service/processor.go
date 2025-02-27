package service

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const (
	dateTimeLayout          = "02 Jan 2006 15:04:05"
	dcaClosedByUserMesssage = "DCA closed by user"
)

var ctx = context.Background()
var slotMap sync.Map

type RealProcessor struct {
	Analyser       Analyser
	Serialiser     Serialiser
	SolanaCaller   SolanaCaller
	RedisCaller    RedisCaller[string, int64]
	TokenFetcher   TokenFetcher
	TelegramCaller TelegramCaller
}

func (r *RealProcessor) Process() {
	slot, _ := r.SolanaCaller.GetSlot()
	if slot.Error.Code != 0 && slot.Error.Message != "" {
		config.Log.Infof("Error when get slot number, error: %q", slot.Error)
		return
	}
	slotNumber := slot.Result
	config.Log.Infof("Begin analyse slot with number: %d", slotNumber)
	_, exists := slotMap.Load(slotNumber)
	if exists == true {
		config.Log.Infof("Slot %d already processed yerlier, skip", slotNumber)
		return
	}
	block, err := r.getBlockWithRetryIfNotAvailbale(slotNumber)
	if err != nil {
		config.Log.Errorf("Error when find block by number: %d, error: %q", slotNumber, err.Error())
		return
	}
	slotMap.Store(slotNumber, nil)
	config.Log.Debugf("Append slot with number: %d to slot map", slotNumber)
	orders := r.Analyser.Analyse(slotNumber, block.Result.Transactions)
	if len(orders) == 0 {
		config.Log.Infof("Slot with number %d not exists DCA orders!", slotNumber)
	} else {
		r.processTransactions(slotNumber, orders)
	}
	slotMap.Delete(slotNumber)
	config.Log.Debugf("Finish procees slot: %d, delet from slot map", slotNumber)
}

func (r *RealProcessor) getBlockWithRetryIfNotAvailbale(slotNumber uint) (model.GetBlockResponseBody, error) {
	block, _ := r.SolanaCaller.GetBlock(slotNumber)
	if block.Error.Code == -32004 {
		config.Log.Infof("Block not available for slot: %d, retry request", slotNumber)
		return r.SolanaCaller.GetBlock(slotNumber)
	}
	if block.Error.Code != 0 && block.Error.Message != "" {
		config.Log.Errorf("Error when get block information by slot with number: %d, error: %s", slotNumber, block.Error)
		return model.GetBlockResponseBody{}, errors.New("unsuccess request")
	}
	return block, nil
}

func (r *RealProcessor) processTransactions(slotNumber uint, orders []model.Transaction) {
	for _, order := range orders {
		if order.Meta.IsOpenDca() {
			err := r.processOpenOrder(slotNumber, order)
			if err != nil {
				config.Log.Errorf("Error when process OPEN DCA order from slot %d, error: %q", slotNumber, err.Error())
			}
		} else if order.Meta.IsCloseDca() {
			err1 := r.processCloseOrder(slotNumber, order)
			if err1 != nil {
				config.Log.Errorf("Error when process CLOSE DCA order from slot %d, error: %q", slotNumber, err1.Error())
			}
		} else {
			config.Log.Warnf("Unknown operation: order NOT CLOSE and NOT OPEN! <--")
		}
	}
}

func (r *RealProcessor) processOpenOrder(slotNumber uint, order model.Transaction) error {
	data, err := r.Serialiser.Serialize(slotNumber, order)
	if err != nil {
		return err
	}
	msg := r.constructTelegramMessage(data)
	tgMessage, err := r.TelegramCaller.SendMessage(msg.String())
	if err != nil {
		config.Log.Errorf("Error when send message %s  from slot %data to telegram, error: ", msg.String(), slotNumber, err.Error())
		return err
	}
	config.Log.Infof("Succuess send telegram message: %s", msg.String())
	dcaKey := order.TransactionDetails.GetDcaKeyOpen()
	err1 := r.RedisCaller.Set(ctx, dcaKey, tgMessage.MessageId, calculateExpirationTime(data))
	if err1 != nil {
		config.Log.Errorf("Error when UPLOAD DCA key %s from slot: %d, error: %q", dcaKey, slotNumber, err1.Error())
		return err1
	}
	config.Log.Infof("Upload message id by DCA key: %s to storage", dcaKey)
	return nil
}

func (r *RealProcessor) processCloseOrder(slotNumber uint, order model.Transaction) error {
	dcaKey := order.TransactionDetails.GetDcaKeyClose()
	messageId, err := r.RedisCaller.Get(ctx, dcaKey)
	if err != nil {
		config.Log.Errorf("Error when get DCA key %s from slot: %d, error: %q", dcaKey, slotNumber, err.Error())
		return err
	}
	config.Log.Infof("Get message id %d by DCA key: %s from redis", messageId, dcaKey)
	err1 := r.TelegramCaller.SendReplyMessage(dcaClosedByUserMesssage, messageId)
	if err1 != nil {
		config.Log.Errorf("Error when reply message with id: %d, error: %q", messageId, err1.Error())
		return err1
	}
	config.Log.Infof("Success send reply to message id: %d", messageId)
	return nil
}

func (r *RealProcessor) constructTelegramMessage(transactionData model.TransactionData) model.TelegramDCAOrderMessage {
	start := time.Now()
	eta := eta(transactionData.InstructionData)
	end := start.Add(time.Duration(eta) * time.Minute)
	symbol := transactionData.TokenSymbol
	tokenInfo := r.TokenFetcher.ExchangeTokenInfo(symbol)
	return model.TelegramDCAOrderMessage{
		Symbol:               symbol,
		Operation:            transactionData.Operation.String(),
		Eta:                  eta,
		PotencialPriceChange: calculatePriceChange(transactionData.InstructionData),
		TokenCA:              transactionData.Token,
		UserAddress:          transactionData.User,
		InAmount:             round(transactionData.InstructionData.InAmount),
		InAmountPerCycle:     round(transactionData.InstructionData.InAmountPerCycle),
		PeriodStart:          start.UTC().Format(dateTimeLayout),
		PeriodEnd:            end.UTC().Format(dateTimeLayout),
		Signature:            transactionData.Signature,
		MexcFutures:          tokenInfo.IsExistsMexc.Bool,
		BitgetFurutes:        tokenInfo.IsExistsBitget.Bool,
		GateFuture:           tokenInfo.IsExistsGate.Bool,
	}
}

func uintToString(val uint) string {
	return strconv.FormatUint(uint64(val), 10)
}

func calculateExpirationTime(data model.TransactionData) time.Duration {
	end := eta(data.InstructionData)
	return time.Duration(end+15) * time.Minute
}

func round(val string) int {
	r, _ := strconv.Atoi(val)
	res := r / 100000
	return res
}

func eta(data model.InstructionData) uint {
	v1, _ := strconv.Atoi(data.InAmount)
	v2, _ := strconv.Atoi(data.InAmountPerCycle)
	return uint(v1 / v2)
}

func calculatePriceChange(instruction model.InstructionData) float32 {
	// TODO <--
	return 1.0
}
