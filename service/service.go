package service

import "web3.kz/solscan/model"

type Processor interface {
	Process()
}

type Analyser interface {
	Analyse(slotNumber uint, transactions []model.Transaction) []model.Transaction
}

type Serialiser interface {
	Serialize(slotNumber uint, orders []model.Transaction) []model.InstructionData
}

type SolanaCaller interface {
	GetSlot() (model.GetSlotResponseBody, error)
	GetBlock(slotNumber uint) (model.GetBlockResponseBody, error)
}