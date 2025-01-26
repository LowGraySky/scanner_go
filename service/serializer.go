package service

import (
	"encoding/hex"
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
	Amount float32
	CycleFrequency string
}

type Instruction struct {
	Data string `json:"data"`
	ProgramId string `json:"programId"`
}

func Serialize(slotNumber uint, orders []model.Transaction) []DcaOrderCoreInformation {
	var dcaOrders []DcaOrderCoreInformation
	for _, tx := range orders {
		var data string
		d := findData(slotNumber, tx.Meta)
		if d == nil {
			config.Log.Errorf("Cant find information abount DCA order data in slot: %d", slotNumber)
			return make([]DcaOrderCoreInformation, 0)
		}

		data = *d
		instructions := serializeInstructionData(data)
		order :=  DcaOrderCoreInformation{
			Amount: getUiTokenAmount(tx),
			CycleFrequency: instructions.CycleFrequency,
		}
		dcaOrders = append(dcaOrders, order)
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

func getUiTokenAmount(tx model.Transaction) float32 {
	return tx.Meta.PreTokenBalances[0].UiTokenAmount.UiAmount
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
