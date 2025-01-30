package model

type InstructionData struct {
	CycleFrequency string
	InAmount string
	InAmountPerCycle string
}

type TransactionData struct {
	Token string
	User string
	InstructionData InstructionData
}