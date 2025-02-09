package model

import (
	"fmt"
	"strings"
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
	Futures              []string
	Signature            string
	MexcFutures          bool
	GateFuture           bool
	BitgetFurutes        bool
}

func (tm TelegramDCAOrderMessage) String() string {
	futures := make([]string, 0)
	if tm.MexcFutures {
		futures = append(futures, "MEXC")
	}
	if tm.BitgetFurutes {
		futures = append(futures, "BITGET")
	}
	if tm.GateFuture {
		futures = append(futures, "GATE")
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
		strings.Join(futures[:], " "),
		tm.TokenCA,
		tm.UserAddress,
		tm.PeriodStart, tm.PeriodEnd,
		tm.Signature,
	)
}
