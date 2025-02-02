package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web3.kz/solscan/service"
)

func TestSerializer(t *testing.T) {
	serializer := new(service.RealSerializer)
	transactions := ReadBlockResponseFromFile().Result.Transactions

	actual := serializer.Serialize(1, transactions)

	c := assert.Equal(t, len(actual), 1)
	ame := assert.Equal(t, actual[0].InstructionData.InAmount, "8272121570535")
	ampe := assert.Equal(t, actual[0].InstructionData.InAmountPerCycle, "33088486282")
	cfe := assert.Equal(t, actual[0].InstructionData.CycleFrequency, "60")
	res := c && ame && ampe && cfe
	if !res {
		t.Error("")
	}
}