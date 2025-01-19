package service

import (
    "sync"
    "web3.kz/solscan/config"
)

var (
	parsedSlotMap sync.Map
)

func Process() {
	slot, _ := GetSlot()
	if slot.Error.Code != 0 && slot.Error.Message != "" {
		config.Log.Printf("Error when get slot number, error: %q", slot.Error)
		return
	}
	slotNumber := slot.Result
	config.Log.Printf("Begin analyse slot with number: %q\n", slotNumber)
	if isAlreadyRead(slotNumber) {
		config.Log.Printf("Slot with number %q already processed, SKIP\n", slotNumber)
	} else {
		config.Log.Printf("Begin analyse slot with number: %q\n", slotNumber)
		block, _ := GetBlock(slotNumber)
		if block.Error.Code != 0 && block.Error.Message != "" {
			config.Log.Printf("Error when get block information by slot with number: %q, error: %q", slotNumber, slot.Error)
			return
		}
		parsedSlotMap.Store(slotNumber, nil)
		Analyse(block.Transactions)
	}
}

func isAlreadyRead(number int) bool {
	_, exists := parsedSlotMap.Load(number)
	return exists
}