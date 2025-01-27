package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web3.kz/solscan/model"
	"web3.kz/solscan/service"
)

func TestAnalyseExists(t *testing.T) {
	analyser := new(service.RealAnalyser)
	blockResponse := ReadBlockResponseFromFile()
	transactions := blockResponse.Result.Transactions

	actual := analyser.Analyse(1, transactions)

	n := assert.NotNil(t, actual)
	if !n {
		t.Error("")
	}
	c := assert.Equal(t, len(actual), 1)
	if !c {
		t.Error("")
	}
}

func TestAnalyseNotExists(t *testing.T) {
	analyser := new(service.RealAnalyser)
	transactions := []model.Transaction{
		model.Transaction{
			Meta: model.Meta{
				LogMessages: []string{
					"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3158 of 216383 compute units",
					"Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success"},
			},
			TransactionDetails: model.TransactionDetails{
				Signatures: []string{"1"},
			},
		},
	}

	actual := analyser.Analyse(1, transactions)

	n := assert.NotNil(t, actual)
	if !n {
		t.Error("")
	}
	c := assert.Equal(t, len(actual), 0)
	if !c {
		t.Error("")
	}
}
