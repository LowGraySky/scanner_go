package model

import "fmt"

const (
	BUY  OrderOperation = "BUY"
	SELL OrderOperation = "SELL"
)

type OrderOperation string

func (o OrderOperation) String() string {
	return fmt.Sprintf("%s", string(o));
}

type InstructionData struct {
	CycleFrequency   string
	InAmount         string
	InAmountPerCycle string
}

type TransactionData struct {
	Token           string
	TokenSymbol     string
	User            string
	Operation       OrderOperation
	InstructionData InstructionData
}
