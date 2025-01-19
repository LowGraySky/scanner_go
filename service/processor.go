package service

import (
	"log/slog"
	"sync"
	"web3.kz/solscan/config"
)

const (
	jupiterDcaAddress  = "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M"
)

var (
	parsedSlotMap sync.Map
)

func Process() {
	slot, _ := GetSlot()
	if slot.Error.Code != 0 && slot.Error.Message != "" {
		config.Log.Info("Error when get slot number, error: %q", slot.Error)
		return
	}
	slotNumber := slot.Result
	config.Log.Printf("Begin analyse slot with number: %q\n", slotNumber)
	if isAlreadyRead(slotNumber) {
		config.Log.Printf("Slot with number %q already processed, SKIP\n", slotNumber)
	} else {
		config.Log.Printf("Begin analyse slot with number: %q\n", slotNumber)
		block := GetBlock(slotNumber)
	}
}



func isAlreadyRead(number int) bool {
	_, exists := parsedSlotMap.Load(number)
	return exists
}