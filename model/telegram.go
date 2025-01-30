package model

type TelegramDCAOrderMessage struct {
	Eta uint `json:"eta"`
	PotencialPriceChange float32 `json:"potencial_price_change"`
	TokenCA string `json:"token_ca"`
	UserAddress string `json:"user_address"`
	InAmount string `json:"in_amount"`
	PeriodStart string `json:"period_start"`
	PeriodEnd string `json:"period_end"`
}

//MCAP: $17.91M
//Liquidity: $597.17K
//ETA: 1 h
//Potential price change: 19.96%
//
//Holders: 14,726
//Vol 24h: $75.71M
//
//CA: Fgr6ZejzV1nguSdkt3ugh5AdzucdHPaMsR2dCF2cpump
//#2dCF2cpump
//
//User: 6tRZkbqwy99hWuExk8TnKYpUHHcQ6jb4FjnNwjh4RkHY
//#nNwjh4RkHY
//
//Amount: 3,926,545.20 BARRON
//Frequency: $1.17K every 60 seconds
//
//Min sell price:  0.000410 SOL per BARRON
//
//Period: 19 Jan 2025 05:31:34 GMT - 19 Jan 2025 06:31:34 GMT
