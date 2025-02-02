package model

import (
	"encoding/json"
	"web3.kz/solscan/config"
)

type TelegramDCAOrderMessage struct {
	Symbol string `json:"symbol"`
	Operation string `json:"operation"`
	Eta uint `json:"eta"`
	PotencialPriceChange float32 `json:"potencial_price_change"`
	TokenCA string `json:"token_ca"`
	UserAddress string `json:"user_address"`
	InAmount string `json:"in_amount"`
	PeriodStart string `json:"period_start"`
	PeriodEnd string `json:"period_end"`
}

func (tm TelegramDCAOrderMessage) String() string {
	str, err := json.Marshal(tm)
	if err != nil {
		config.Log.Errorf("Error when convert data to json, error: %q", err.Error())
	}
	return string(str)
}