package model

import "time"

type TransactionCoreInfo struct {
	Capitalisation uint64
	Liquidity uint64
	Eta uint
	PotencialPriceChange uint
	TokenContract string //
	UserAddress string // // result -> transactions -> transaction -> message -> accountKeys -> pubkey (where signer = true)
	Amount uint64  //
	FrequencyAmount uint
	StartPeriod time.Duration
	EndPeriod time.Duration
}



//MCAP: $17.91M
//Liquidity: $597.17K
//ETA: 1 h
//Potential price change: 19.96%
//CA: Fgr6ZejzV1nguSdkt3ugh5AdzucdHPaMsR2dCF2cpump
//User: 6tRZkbqwy99hWuExk8TnKYpUHHcQ6jb4FjnNwjh4RkHY
//Amount: 3,926,545.20 BARRON
//Frequency: $1.17K every 60 seconds
//Min sell price:  0.000410 SOL per BARRON
//Period: 19 Jan 2025 05:31:34 GMT - 19 Jan 2025 06:31:34 GMT




// meta -> logMessages -> contains "Program log: Instruction: OpenDcaV2"
// result -> transactions -> transaction -> message -> accountKeys (where signer = true)
// result -> transactions -> transaction -> signatures <-------- подпись