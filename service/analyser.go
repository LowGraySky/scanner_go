package service

import (
	"web3.kz/solscan/config"
)

const (
	jupiterDcaAddress  = "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M"
)

func Analyse(transactions []Transaction) {
	if len(transactions) == 0 {
		config.Log.Println("Skip analyse block, transaction count is 0!")
		return
	}
	
}