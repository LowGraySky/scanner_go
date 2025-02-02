package service

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"strconv"
	"sync"
	"time"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const dateTimeLayout = "02 Jan 2006 15:04:05"

var (
	parsedSlotMap sync.Map
)

type RealProcessor struct {
	Analyser       Analyser
	Serialiser     Serialiser
	SolanaCaller   SolanaCaller
	TelegramCaller TelegramCaller
}

func (r *RealProcessor) Process(bot gotgbot.Bot) {
	slot, _ := r.SolanaCaller.GetSlot()
	if slot.Error.Code != 0 && slot.Error.Message != "" {
		config.Log.Infof("Error when get slot number, error: %q", slot.Error)
		return
	}
	slotNumber := slot.Result
	if isAlreadyRead(slotNumber) {
		config.Log.Infof("Slot with number %d already processed, SKIP", slotNumber)
	} else {
		config.Log.Infof("Begin analyse slot with number: %d", slotNumber)
		block, _ := r.SolanaCaller.GetBlock(slotNumber)
		if block.Error.Code != 0 && block.Error.Message != "" {
			config.Log.Errorf("Error when get block information by slot with number: %d, error: %s", slotNumber, block.Error)
			return
		}
		parsedSlotMap.Store(slotNumber, nil)
		orders := r.Analyser.Analyse(slotNumber, block.Result.Transactions)
		if len(orders) == 0 {
			config.Log.Infof("Slot with number %d not exists DCA orders!", slotNumber)
		} else {
			txData := r.Serialiser.Serialize(slotNumber, orders)
			for _, d := range txData {
				msg := constructTelegramMessage(d)
				r.TelegramCaller.SendMessage(bot, msg.String())
			}
		}
	}
}

func isAlreadyRead(number uint) bool {
	_, exists := parsedSlotMap.Load(number)
	return exists
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
	}
}

func round(val string) int {
	r, _ := strconv.Atoi(val)
	res := r/100000
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
