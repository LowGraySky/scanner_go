package service

import (
	"fmt"
	"sync"
)

var (
    parsedSlotMap sync.Map
)

func Analyse() {
	slot := GetSlot()
	if slot.isAlreadyRead() {
		fmt.Printf("Slot with number %q already processed, SKIP", slot.Result)
	} else {
		slotNumber := slot.Result
		fmt.Printf("Begin analyse slot with number: %q", slotNumber)
		block := GetBlock(slotNumber)
		
	}
}

func (s GetSlotResponseBody) isAlreadyRead() bool {
	number := s.Result
	_, exists := parsedSlotMap.Load(number)
	return exists
}