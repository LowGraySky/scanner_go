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
}

func (tm TelegramDCAOrderMessage) String() string {
	futures := ""
	if tm.MexcFutures {
		futures = "Futures: MEXC"
	}
	return fmt.Sprintf(`
%d %s %s

Frequency: %d every 60 seconds (%d cycles)
ETA: %dm
Potential price change: %d %
%s
CA: %s

User: %s
Period: %s - %s`,
		tm.InAmount, tm.Operation, tm.Symbol,
		tm.InAmountPerCycle, tm.Eta,
		tm.Eta,
		tm.PotencialPriceChange,
		futures,
		tm.TokenCA,
		tm.UserAddress,
		tm.PeriodStart, tm.PeriodEnd,
	)
}

//$70.62K selling HOOD 🟥
//
//Frequency: $706.19 every 60 seconds (101 cycles)
//ETA: 1h, 41m
//Scores: 🤔
//Potential price change: 10.387% (0.113% per cycle)
//
//MC: $33.57M → LQ: $1.25M
//Holders: 64,245
//V24h: $49.79M → V1h: $1.68M → VI1h: 3372.31792%
//Price: $0.00075
//
//Futures: MEXC
//
//Trade bots: BLX - PHO - PEP - STB - TRO - BLO - BNK
//
//CA: h5NciPdMZ5QCB5BYETJMYBMpVx9ZuitR6HcVjyBhood
//#HcVjyBhood
//
//User: HvDf4Cxd2evdYueLhK5LoaiEvDXFXgb1uRrkoYPdvHfH
//#rkoYPdvHfH
//
//Period: 02 Feb 2025 15:45:49 - 02 Feb 2025 17:26:49 GMT
