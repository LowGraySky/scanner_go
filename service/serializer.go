package service

import (
	"encoding/hex"
	"errors"
	"github.com/mr-tron/base58"
	"math/big"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const dcaOpenV2ProgramId = "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M"

type InstructionData struct {
	CycleFrequency string
	InAmount string
	InAmountPerCycle string
}

type DcaOrderCoreInformation struct {
	
}

func Serialize(tx model.Transaction) DcaOrderCoreInformation {
	data, err := findData(tx.Meta.InnerInstructions)
	if err != nil {
		config.Log.Errorf("")
	}
	instrucitonData := serializeInstructionData(data)

}

func findData(instructions model.InnerInstructions) (string, error) {
	for _, inst := range instructions.Instructions {
		if inst.ProgramId == dcaOpenV2ProgramId {
			return inst.Data, nil
		}
	}
	return "", errors.New("")
}

func serializeInstructionData(data string) InstructionData {
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
		cycleFrequency.String(),
		inAmount.String(),
		inAmountPerCycle.String(),
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
