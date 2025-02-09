package model

import "fmt"

const (
	openDcaV2Comment  = "Program log: Instruction: OpenDcaV2"
	closeDcaV2Comment = "Program log: Instruction: CloseDca"
	dcaOpenV2ProgramId = "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M"
)

type RpcCallWithoutParameters struct {
	JsonRpc string `json:"jsonrpc"`
	Id      uint   `json:"id"`
	Method  string `json:"method"`
}

type RpcCallWithParameters struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      uint          `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type GetBlockParamsBody struct {
	Enconding          string `json:"encoding"`
	TransactionVersion uint   `json:"maxSupportedTransactionVersion"`
	Rewards            bool   `json:"rewards"`
}

type GetSlotResponseBody struct {
	JsonRpc string `json:"jsonrpc"`
	Result  uint   `json:"result"`
	Id      int    `json:"id"`
	Error   Error  `json:"error"`
}

func (r GetSlotResponseBody) String() string {
	return fmt.Sprintf("Result: %d, Error: %q", r.Result, r.Error)
}

type GetBlockResponseBody struct {
	JsonRpc string                     `json:"jsonrpc"`
	Result  GetBlockResponseBodyResult `json:"result"`
	Id      uint                       `json:"id"`
	Error   Error                      `json:"error"`
}

type GetBlockResponseBodyResult struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Meta               Meta               `json:"meta"`
	TransactionDetails TransactionDetails `json:"transaction"`
}

type Meta struct {
	PostTokenBalances []struct {
		Mint string `json:"mint"`
	} `json:"postTokenBalances"`
	LogMessages []string `json:"logMessages"`
}

func (m *Meta) IsLogMesssagesExists() bool {
	return m.LogMessages != nil
}

func (m *Meta) IsCloseDca() bool {
	for _, log := range m.LogMessages {
		if log == closeDcaV2Comment {
			return true
		}
	}
	return false
}

func (m *Meta) IsOpenDca() bool {
	for _, log := range m.LogMessages {
		if log == openDcaV2Comment {
			return true
		}
	}
	return false
}

func (m *Meta) GetTokenAddress() []string  {
	tokens := make([]string, 2)
	tokens[0] = m.PostTokenBalances[0].Mint
	for _, t := range m.PostTokenBalances {
		if t.Mint != tokens[0] {
			tokens[1] = t.Mint
			return tokens
		}
	}
	return tokens
}

type TransactionDetails struct {
	Message struct {
		AccountKeys []struct {
			Pubkey string `json:"pubkey"`
			Signer bool   `json:"signer"`
		} `json:"accountKeys"`
		Instructions []struct {
			Accounts  []string `json:"accounts"`
			ProgramId string   `json:"programId"`
			Data      *string  `json:"data,omitempty"`
		} `json:"instructions"`
	} `json:"message"`
	Signatures []string `json:"signatures"`
}

func (t *TransactionDetails) GetUserCA() string {
	for _, acc := range t.Message.AccountKeys {
		if acc.Signer == true {
			return acc.Pubkey
		}
	}
	return ""
}

func (t TransactionDetails) String() string {
	return fmt.Sprintf("Signatures: %s", t.Signatures)
}

func (t *TransactionDetails) GetDcaKeyOpen() string  {
	for _, inst := range t.Message.Instructions {
		if inst.ProgramId == dcaOpenV2ProgramId {
			return inst.Accounts[0]
		}
	}
	return ""
}

func (t *TransactionDetails) GetDcaKeyClose() string  {
	for _, inst := range t.Message.Instructions {
		if inst.ProgramId == dcaOpenV2ProgramId {
			return inst.Accounts[1]
		}
	}
	return ""
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) String() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}
