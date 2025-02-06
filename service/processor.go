package service

import (
	"context"
	"strconv"
	"time"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const (
	dateTimeLayout = "02 Jan 2006 15:04:05"
	dcaClosedByUserMesssage = "DCA closed by user"
)

var ctx = context.Background()

type RealProcessor struct {
	Analyser       Analyser
	Serialiser     Serialiser
	SolanaCaller   SolanaCaller
	RedisCaller    RedisCaller[string, int64]
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
	block, _ := r.SolanaCaller.GetBlock(slotNumber)
	if block.Error.Code != 0 && block.Error.Message != "" {
		config.Log.Errorf("Error when get block information by slot with number: %d, error: %s", slotNumber, block.Error)
		return
	}
	orders := r.Analyser.Analyse(slotNumber, block.Result.Transactions)
	if len(orders) == 0 {
		config.Log.Infof("Slot with number %d not exists DCA orders!", slotNumber)
	} else {
		r.processTransactions(slotNumber, orders)
	}
}

func (r *RealProcessor) processTransactions(slotNumber uint, orders []model.Transaction) {
	for _, order := range orders {
		if order.Meta.IsOpenDca() {
			err := r.processOpenOrder(slotNumber, order)
			if err != nil {
				config.Log.Errorf("")
				// config.Log.Errorf("Message %s hasn't delivered!, err: %q", msg.String(), err.Error()) <-- change
			}
		} else if order.Meta.IsCloseDca() {

		} else {
			config.Log.Warnf("")
		}
	}
}

func (r *RealProcessor) processOpenOrder(slotNumber uint, order model.Transaction) error {
	data, err := r.Serialiser.Serialize(slotNumber, order)
	if err != nil {
		return err
	}
	msg := constructTelegramMessage(data)
	tgMessage, err := r.TelegramCaller.SendMessage(msg.String())
	if err != nil {
		config.Log.Errorf("Error when send message %s  from slot %data to telegram, error: ", msg.String(), slotNumber, err.Error())
		return err
	}
	config.Log.Infof("Succuess send telegram message: %s", msg.String())
	dcaKey :=  order.TransactionDetails.GetDcaKeyOpen()
	err1 := r.RedisCaller.Set(ctx, dcaKey, tgMessage.MessageId, calculateExpirationTime(data))
	if err1 != nil {
		config.Log.Errorf("Error when UPLOAD DCA key %s from slot: %data to redis, error: %q", dcaKey, slotNumber, err1.Error())
		return err1
	}
	return nil
}

func (r *RealProcessor) processCloseOrder(order model.Transaction) {
	messageId, err := r.RedisCaller.Get(ctx, order.TransactionDetails.GetDcaKeyClose())
	if err != nil {
		config.Log.Error("Error when GET ")
	}
	err1 := r.TelegramCaller.SendReplyMessage(dcaClosedByUserMesssage, messageId)
	if err1 != nil {
		config.Log.Errorf("Error when reply ")
	}
}

func calculateExpirationTime(data model.TransactionData) time.Duration {
	end := eta(data.InstructionData)
	return time.Duration(end + 15) * time.Minute
}

func constructTelegramMessage(transactionData model.TransactionData) model.TelegramDCAOrderMessage {
	start := time.Now()
	eta := eta(transactionData.InstructionData)
	end := start.Add(time.Duration(eta) * time.Minute)
	return model.TelegramDCAOrderMessage{
		Symbol:               transactionData.TokenSymbol,
		Operation:            transactionData.Operation.String(),
		Eta:                  eta,
		PotencialPriceChange: calculatePriceChange(transactionData.InstructionData),
		TokenCA:              transactionData.Token,
		UserAddress:          transactionData.User,
		InAmount:             round(transactionData.InstructionData.InAmount),
		InAmountPerCycle:     round(transactionData.InstructionData.InAmountPerCycle),
		PeriodStart:          start.UTC().Format(dateTimeLayout),
		PeriodEnd:            end.UTC().Format(dateTimeLayout),
		MexcFutures:          true, // TODO <--
		Signature:            transactionData.Signature,
	}
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
