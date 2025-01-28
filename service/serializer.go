package service

import (
	"encoding/hex"
	"github.com/mr-tron/base58"
	"math/big"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const dcaOpenV2ProgramId = "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M"

type RealSerializer struct {}

func (s *RealSerializer) Serialize(slotNumber uint, orders []model.Transaction) []model.InstructionData {
	var dcaOrders []model.InstructionData
	for _, tx := range orders {
		var data string
		d := findData(slotNumber, tx.Meta)
		if d == nil {
			config.Log.Errorf("Cant find information abount DCA order data in slot: %d", slotNumber)
			return make([]model.InstructionData, 0)
		}
		data = *d
		instructions := serializeInstructionData(data)
		dcaOrders = append(dcaOrders, instructions)
	}
	return dcaOrders
}

func findData(slotNumber uint, meta model.Meta) *string {
	for _, inst := range meta.InnerInstructions {
		if inst.ProgramId == dcaOpenV2ProgramId {
			config.Log.Infof("Find data: %s in slot: %d", *inst.Data, slotNumber)
			return inst.Data
		}
	}
	return nil
}

func serializeInstructionData(data string) model.InstructionData {
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
	return model.InstructionData{
		CycleFrequency:   cycleFrequency.String(),
		InAmount:         inAmount.String(),
		InAmountPerCycle: inAmountPerCycle.String(),
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
