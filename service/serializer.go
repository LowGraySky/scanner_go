package service

import (
	"encoding/hex"
	"errors"
	"github.com/mr-tron/base58"
	"math/big"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

var stables = map[string]struct{} {
	"EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v": {},
	"So11111111111111111111111111111111111111112": {},
	"Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB": {},
}

const (
	dcaOpenV2ProgramId = "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M"
)

type RealSerializer struct {
	TokenFetcher TokenFetcher
}

func (rs *RealSerializer) Serialize(slotNumber uint, order model.Transaction) (model.TransactionData, error) {
	var data string
	d := findData(slotNumber, order.TransactionDetails)
	if d == nil {
		config.Log.Errorf("Cant find information abount DCA order data in slot: %d", slotNumber)
		return model.TransactionData{}, errors.New("transaction data is empty")
	}
	data = *d
	instruction := serializeInstructionData(data)
	tranasctionData, err := rs.createTransactionAditionalData(order, instruction)
	if err != nil {
		config.Log.Errorf("Error when serilizing transaction in slot: %d, error: %q", slotNumber, err.Error())
		return model.TransactionData{}, err
	}
	return tranasctionData, nil
}

func (rs *RealSerializer) createTransactionAditionalData(tx model.Transaction, inst model.InstructionData) (model.TransactionData, error) {
	token, operation := defineTokenAndOrderOperation(tx)
	tokenInfo, err := rs.TokenFetcher.GetTokenInfo(token)
	if err != nil {
		return model.TransactionData{}, err
	}
	return model.TransactionData{
		DcaKey:			 getDcakey(tx),
		Token:           token,
		TokenSymbol:     tokenInfo.Symbol,
		User:            tx.TransactionDetails.GetUserCA(),
		Operation:       operation,
		InstructionData: inst,
		Signature:       tx.TransactionDetails.Signatures[0],
	}, nil
}

func getDcakey(tx model.Transaction) string  {
	if tx.Meta.IsOpenDca() {
		return tx.TransactionDetails.GetDcaKeyOpen()
	} else {
		return tx.TransactionDetails.GetDcaKeyClose()
	}
}

func defineTokenAndOrderOperation(tx model.Transaction) (string, model.OrderOperation) {
	tokens := tx.Meta.GetTokenAddress()
	_, in := stables[tokens[0]]
	_, out := stables[tokens[1]]
	if in {
		return tokens[1], model.BUY
	}
	if out {
		return tokens[0], model.SELL
	}
	return tokens[1], model.BUY
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
