package service

import (
	"fmt"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const (
	openDcaV2Comment = "Program log: Instruction: OpenDcaV2"
)

type RealAnalyser struct {}

func (a *RealAnalyser) Analyse(
	slotNumber uint,
	transactions []model.Transaction) []model.Transaction {
	if len(transactions) == 0 {
		config.Log.Infof("Skip analyse block with slot number %q, transaction count is 0!", slotNumber)
		return make([]model.Transaction, 0)
	} else {
		var orders []model.Transaction
		config.Log.Infof("Slot with number %d got transaction, start search", slotNumber)
		for _, tx := range transactions {
			meta := tx.Meta
			if isLogMesssagesExists(meta) && isOpenDcaV2(meta) {
				fmt.Printf("Found DCA order, tx: %q", tx.TransactionDetails.Signatures)
				orders = append(orders, tx)
			}
		}
		return orders
	}
}

func isLogMesssagesExists(meta model.Meta) bool {
	return meta.LogMessages != nil
}

func isOpenDcaV2(meta model.Meta) bool {
	for _, log := range meta.LogMessages {
		if log == openDcaV2Comment {
			return true
		}
	}
	return false
}

