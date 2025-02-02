package service

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"web3.kz/solscan/model"
)

type Processor interface {
	Process(bot gotgbot.Bot)
}

type Analyser interface {
	Analyse(slotNumber uint, transactions []model.Transaction) []model.Transaction
}

type Serialiser interface {
	Serialize(slotNumber uint, orders []model.Transaction) []model.TransactionData
}

type SolanaCaller interface {
	GetSlot() (model.GetSlotResponseBody, error)
	GetBlock(slotNumber uint) (model.GetBlockResponseBody, error)
}

type TelegramCaller interface {
	StartBot() (*gotgbot.Bot, error)
	SendMessage(bot gotgbot.Bot, message string)
}

type JupiterCaller interface {
	GetToken(address string) (model.TokenInfo, error)
}
