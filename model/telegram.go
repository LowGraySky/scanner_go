package model

import (
	"fmt"
)

type TelegramDCAOrderMessage struct {
	Symbol               string
	Operation            string
	Eta                  uint
	PotencialPriceChange float32
	TokenCA              string
	UserAddress          string
	InAmount             int
	InAmountPerCycle     int
	PeriodStart          string
	PeriodEnd            string
	MexcFutures          bool
	Signature            string
}

func (tm TelegramDCAOrderMessage) String() string {
	futures := ""
	if tm.MexcFutures {
		futures = "MEXC"
	}
	operationSymbol := "🟩"
	if tm.Operation == "SELL" {
		operationSymbol = "🟥"
	}
	return fmt.Sprintf(`
%d %s %s %s

<b>Frequency</b>: %d every 60 seconds (%d cycles)
<b>ETA</b>: %dm
<b>Potential price change</b>: %f %%
<b>Futures</b>: %s
<b>CA</b>: %s

<b>User</b>: %s
<b>Period</b>: %s - %s GMT

<a href="https://solscan.io/tx/%s">Solscan</a>`,
		tm.InAmount, tm.Operation, tm.Symbol, string(operationSymbol),
		tm.InAmountPerCycle, tm.Eta,
		tm.Eta,
		tm.PotencialPriceChange,
		string(futures),
		tm.TokenCA,
		tm.UserAddress,
		tm.PeriodStart, tm.PeriodEnd,
		tm.Signature,
	)
}
