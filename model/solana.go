package model

import "fmt"

type RpcCallWithoutParameters struct {
	JsonRpc string `json:"jsonrpc"`
	Id uint `json:"id"`
	Method string `json:"method"`
}

type RpcCallWithParameters struct {
	JsonRpc string `json:"jsonrpc"`
	Id uint `json:"id"`
	Method string `json:"method"`
	Params []interface{} `json:"params"`
}

type GetBlockParamsBody struct {
	Enconding string `json:"encoding"`
	TransactionVersion uint `json:"maxSupportedTransactionVersion"`
	Rewards bool `json:"rewards"`
}

type GetSlotResponseBody struct {
	JsonRpc string `json:"jsonrpc"`
	Result uint `json:"result"`
	Id int `json:"id"`
	Error Error `json:"error"`
}

func (r GetSlotResponseBody) String() string {
	return fmt.Sprintf("Result: %d, Error: %q", r.Result, r.Error)
}

func (e Error) String() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

type GetBlockResponseBody struct {
	JsonRpc string `json:"jsonrpc"`
	Result GetBlockResponseResultBody `json:"result"`
	Id uint `json:"id"`
	Error Error `json:"error"`
}

type Error struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type GetBlockResponseResultBody struct {
	BlockHash string `json:"blockhash"`
	PreviousBlockhash string `json:"previousBlockhash"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Meta Meta `json:"meta"`
	TransactionDetails TransactionDetails `json:"transaction"`
}

type Meta struct {
	LogMessages []string `json:"logMessages"`
}

type TransactionDetails struct {
	Signatures []string `json:"signatures"`
}