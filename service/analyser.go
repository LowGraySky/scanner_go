package service

import (
	"sync"
	"web3.kz/solscan/config"
)

const (
	jupiterDcaAddress  = "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M"
)

var (
	parsedSlotMap sync.Map
)

func Analyse() {
	//slot, _ := GetSlot()
//	slotNumber := slot.Result

	slotNumber, _ := GetSlot()
	config.Log.Printf("Begin analyse slot with number: %q\n", slotNumber.Result)

//	if isAlreadyRead(slotNumber) {
//		config.Log.Printf("Slot with number %q already processed, SKIP\n", slotNumber)
//	} else {
//		config.Log.Printf("Begin analyse slot with number: %q\n", slotNumber)
//		block := GetBlock(slotNumber)
	//}
}

func isAlreadyRead(number int64) bool {
	_, exists := parsedSlotMap.Load(number)
	return exists
}