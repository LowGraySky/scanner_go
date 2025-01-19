package model

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

type GetBlockResponseBody struct {
	JsonRpc string `json:"jsonrpc"`
	GetBlockResponseResultBody  `json:"result"`
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
	PostBalances []int64 `json:"PostBalances"`
	PreBalances []int64 `json:"PreBalances"`
}

type TransactionDetails struct {
	Signatures []string `json:"Signatures"`
}