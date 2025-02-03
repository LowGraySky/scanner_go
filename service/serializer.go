package service

import (
	"encoding/hex"
	"github.com/mr-tron/base58"
	"math/big"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

var stables = map[string]int {
	"EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v": 0,
	"So11111111111111111111111111111111111111112": 0,
	"Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB": 0,
}

const (
	dcaOpenV2ProgramId      = "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M"
)

type RealSerializer struct {
	JupiterCaller JupiterCaller
}

func (s *RealSerializer) Serialize(slotNumber uint, orders []model.Transaction) []model.TransactionData {
	var dcaOrders []model.TransactionData
	for _, tx := range orders {
		var data string
		d := findData(slotNumber, tx.TransactionDetails)
		if d == nil {
			config.Log.Errorf("Cant find information abount DCA order data in slot: %d", slotNumber)
			return make([]model.TransactionData, 0)
		}
		data = *d
		instructions := serializeInstructionData(data)
		tranasctionData, err := s.createTransactionAditionalData(tx, instructions)
		if err != nil {
			config.Log.Errorf("Error when serilizing transaction in slot: %d, error: %q", slotNumber, err.Error())
			continue
		}
		dcaOrders = append(dcaOrders, tranasctionData)
	}
	return dcaOrders
}

func (s *RealSerializer) createTransactionAditionalData(tx model.Transaction, inst model.InstructionData) (model.TransactionData, error) {
	token, operation := defineTokenAndOrderOperation(tx)
	tokenInfo, err1 := s.JupiterCaller.GetToken(token)
	if err1 != nil {
		return model.TransactionData{}, err1
	}
	return model.TransactionData{
		Token:           token,
		TokenSymbol:     tokenInfo.Symbol,
		User:            findUserCA(tx),
		Operation:       operation,
		InstructionData: inst,
		Signature: 		 tx.TransactionDetails.Signatures[0],
	}, nil
}

func defineTokenAndOrderOperation(tx model.Transaction) (string, model.OrderOperation) {
	tokens := collectTokenAddress(tx)
	_, ex1 := stables[tokens[0]]
	_, ex2 := stables[tokens[1]]
	if ex1 {
		return tokens[1], model.BUY
	}
	if ex2 {
		return tokens[0], model.SELL
	}
	return tokens[1], model.BUY
}

func collectTokenAddress(tx model.Transaction) []string {
	tokens := make([]string, 2)
	for _, balance := range tx.Meta.PostTokenBalances {
		if tokens[0] == "" {
			tokens[0] = balance.Mint
		} else {
			tokens[1] = balance.Mint
		}
	}
	return tokens
}

func findData(slotNumber uint, txDetails model.TransactionDetails) *string {
	for _, inst := range txDetails.Message.Instructions {
		if inst.ProgramId == dcaOpenV2ProgramId {
			config.Log.Infof("Find data: %s in slot: %d", *inst.Data, slotNumber)
			return inst.Data
		}
	}
	return nil
}

func findUserCA(tx model.Transaction) string {
	for _, acc := range tx.TransactionDetails.Message.AccountKeys {
		if acc.Signer == true {
			return acc.Pubkey
		}
	}
	return ""
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
