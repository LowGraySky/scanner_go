package service

import (
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"time"
	"web3.kz/solscan/model"
)

type Processor interface {
	Process()
}

type Analyser interface {
	Analyse(slotNumber uint, transactions []model.Transaction) []model.Transaction
}

type Serialiser interface {
	Serialize(slotNumber uint, orders model.Transaction) (model.TransactionData, error)
}

type SolanaCaller interface {
	GetSlot() (model.GetSlotResponseBody, error)
	GetBlock(slotNumber uint) (model.GetBlockResponseBody, error)
}

type TelegramCaller interface {
	SendMessage(message string) (*gotgbot.Message, error)
	SendReplyMessage(message string, messageId int64) error
}

type JupiterCaller interface {
	GetToken(address string) (model.TokenInfo, error)
}

type RedisCaller[T any, R any] interface {
	Get(ctx context.Context, key T) (R, error)
	Set(ctx context.Context, key T, value R, expiration time.Duration) error
}

type TokenFetcher interface {
	GetTokenInfo(address string) (model.TokenInfo, error)
	IsExistsOnMexc(symbol string) bool
	IsExistsOnGate(symbol string) bool
	IsExistsOnBitget(symbol string) bool
}

type MexcCaller interface {
	GetToken(symbol string) (model.MexcTokenInfoResponse, error)
}

type GateCaller interface {
	GetToken(symbol string) (model.GateTokenInfoResponse, error)
}
