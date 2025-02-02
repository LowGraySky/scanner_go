package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web3.kz/solscan/service"
)

func TestSerializer(t *testing.T) {
	serializer := service.RealSerializer{
		JupiterCaller: &service.RealJupiterCaller{},
	}
	transactions := ReadBlockResponseFromFile().Result.Transactions

	actual := serializer.Serialize(1, transactions)

	c := assert.Equal(t, len(actual), 1)
	res := actual[0]

	ame := assert.Equal(t, res.InstructionData.InAmount, "8272121.570535")
	ampe := assert.Equal(t, res.InstructionData.InAmountPerCycle, "33088.486282")
	cfe := assert.Equal(t, res.InstructionData.CycleFrequency, "60")
	tkcae := assert.Equal(t, res.Token, "DVZrNS9fctrrDmhZUZAu6p63xU6d9cqYxRRhJbtJ4z8G")
	tksymbe := assert.Equal(t, res.TokenSymbol, "Ross")
	opee := assert.Equal(t, res.Operation.String(), "SELL")
	usae := assert.Equal(t, res.User, "7DiaCzvNmMcA7z8J3McC3VaUDJJTdKPQCd9YTAThSTaY")

	r := c && ame && ampe && cfe && tkcae && tksymbe && opee && usae
	if !r {
		t.Error("")
	}
}