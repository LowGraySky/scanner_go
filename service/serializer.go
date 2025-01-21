package service

import (
	"encoding/hex"
	"github.com/mr-tron/base58"
	"math/big"
)

type InstructionData struct {
	CycleFrequency big.Int
	InAmount big.Int
	InAmountPerCycle
}

func SerializeInstructionData(data string) InstructionData {
	decodedData, _ := base58.Decode(data)
	hexString := hex.EncodeToString(decodedData)

	inAmountBytes := hexString[16*2 : 24*2]
	cycleFrequencyBytes := hexString[32*2 : (32+8)*2]
	inAmountPerCycleBytes := hexString[24*2 : 32*2]

	reversedCycleFrequencyBytes := reverseHexBytes(cycleFrequencyBytes)
	reversedInAmountBytes := reverseHexBytes(inAmountBytes)
	reversedInAmountPerCycleBytes := reverseHexBytes(inAmountPerCycleBytes)

	cycleFrequency := new(big.Int)
	cycleFrequency.SetString(reversedCycleFrequencyBytes, 16)
	inAmount := new(big.Int)
	inAmount.SetString(reversedInAmountBytes, 16)
	inAmountPerCycle := new(big.Int)
	inAmountPerCycle.SetString(reversedInAmountPerCycleBytes, 16)
	return InstructionData{
		*cycleFrequency,
		*inAmount,
		*inAmountPerCycle
	}
}


func reverseHexBytes(hexStr string) string {
	n := len(hexStr)
	reversed := make([]byte, n)
	for i := 0; i < n; i += 2 {
		reversed[n-2-i] = hexStr[i]
		reversed[n-1-i] = hexStr[i+1]
	}
	return string(reversed)
}
