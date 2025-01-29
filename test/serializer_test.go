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
	if !c {
		t.Error("")
	}
	ame := assert.Equal(t, actual[0].InAmount, "8272121570535")
	if !ame {
		t.Error("")
	}
	ampe := assert.Equal(t, actual[0].InAmountPerCycle, "33088486282")
	if !ampe {
		t.Error("")
	}
	cfe := assert.Equal(t, actual[0].CycleFrequency, "60")
	if !cfe {
		t.Error("")
	}
}