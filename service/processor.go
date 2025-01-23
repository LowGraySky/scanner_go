package service

import (
    "sync"
    "web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

var (
	parsedSlotMap sync.Map
)

func Process() {
	slot, _ := GetSlot()
	if slot.Error.Code != 0 && slot.Error.Message != "" {
		config.Log.Infof("Error when get slot number, error: %q", slot.Error)
		return
	}
	slotNumber := slot.Result
	if isAlreadyRead(slotNumber) {
		config.Log.Infof("Slot with number %d already processed, SKIP", slotNumber)
	} else {
		config.Log.Infof("Begin analyse slot with number: %d", slotNumber)
		block, _ := GetBlock(slotNumber)
		if block.Error.Code != 0 && block.Error.Message != "" {
			config.Log.Infof("Error when get block information by slot with number: %d, error: %q", slotNumber, slot.Error)
			return
		}
		parsedSlotMap.Store(slotNumber, nil)
		channel := make(chan model.Transaction)
		go Analyse(slotNumber, block.Result.Transactions, channel)
		tx := <- channel
		Serialize(tx)
	}
}

func isAlreadyRead(number uint) bool {
	_, exists := parsedSlotMap.Load(number)
	return exists
}