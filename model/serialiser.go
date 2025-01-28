package model

type InstructionData struct {
	CycleFrequency string
	InAmount string
	InAmountPerCycle string
}

type DcaOrderCoreInformation struct {
	Amount float32
	AmountPerCycle string
	CycleFrequency string
}