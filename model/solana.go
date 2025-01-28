package model

import "fmt"

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

func (e Error) String() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func (td TransactionDetails) String() string {
	return fmt.Sprintf("Signatures: %s", td.Signatures)
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
	LogMessages       []string          `json:"logMessages"`
	InnerInstructions []Instruction     `json:"innerInstructions"`
}

type Instruction struct {
	ProgramId string  `json:"programId"`
	Data      *string `json:"data,omitempty"`
}

type PreTokenBalance struct {
	UiTokenAmount UiTokenAmount `json:"uiTokenAmount"`
}

type UiTokenAmount struct {
	UiAmount float32 `json:"uiAmount"`
	Decimals uint    `json:"decimals"`
}

type TransactionDetails struct {
	Message    Message  `json:"message"`
	Signatures []string `json:"signatures"`
}

type AccountKey struct {
	Pubkey string `json:"pubkey"`
	Signer bool   `json:"signer"`
}

type Message struct {
	AccountKeys []AccountKey `json:"accountKeys"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
