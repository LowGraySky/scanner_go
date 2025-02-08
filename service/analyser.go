package service

import (
	"fmt"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

type RealAnalyser struct{}

func (a *RealAnalyser) Analyse(
	slotNumber uint,
	transactions []model.Transaction) []model.Transaction{
	if len(transactions) == 0 {
		config.Log.Infof("Skip analyse block with slot number %q, transaction count is 0!", slotNumber)
		return make([]model.Transaction, 0)
	} else {
		var orders []model.Transaction
		config.Log.Infof("Slot with number %d got transaction, start search", slotNumber)
		for _, tx := range transactions {
			meta := tx.Meta
			if meta.IsLogMesssagesExists() && (meta.IsOpenDca() || meta.IsCloseDca()) {
				config.Log.Info("Found DCA order, tx: %q", tx.TransactionDetails.Signatures)
				orders = append(orders, tx)
			}
		}
		if orders == nil {
			return make([]model.Transaction, 0)
		} else {
			return orders
		}
	}
}
