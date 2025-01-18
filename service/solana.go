package service

import (
	"bytes"
	"encoding/json"
	"web3.kz/solscan/config"
	"io"
	"net/http"
)

const (
	BaseUrl = "https://api.mainnet-beta.solana.com"
	ApplicationJsonContentType = "application/json"
	JsonRpcValue = "2.0"
	IdValue = 1
	GetSlotMethodName = "getSlot"
	GetBlockMethodName = "getBlock"
)

func GetSlot() GetSlotResponseBody {
	rpcCall := RpcCallWithoutParameters {
		JsonRpc: JsonRpcValue,
		Id: IdValue,
		Method: GetSlotMethodName,
	}
	res, err := http.Post(BaseUrl, ApplicationJsonContentType, toJsonIoReader(rpcCall))
	if err != nil {
		config.Log.Printf("Error when request to solana RPC, method: '%q', error: %q\n", GetSlotMethodName, err.Error())
	}
	config.Log.Printf("Got reponse from solana RPC, method: '%q', code: %d\n", GetSlotMethodName, res.StatusCode)
	var slotResponse GetSlotResponseBody
	readResponseBody(res.Body, slotResponse)
	return slotResponse
}

func GetBlock(slotNumber int64) GetBlockResponseBody {
	rpcCall := RpcCallWithParameters {
		JsonRpc: JsonRpcValue,
		Id: IdValue,
		Method: GetBlockMethodName,
		Params: []interface{}{ slotNumber, defaultGetBlockParamsBody },
	}
	res, err := http.Post(BaseUrl, ApplicationJsonContentType, toJsonIoReader(rpcCall))
	if err != nil {
		config.Log.Printf("Error when request to solana RPC, method: '%q', error: %q\n", GetBlockMethodName, err.Error())
	}
	config.Log.Printf("Got reponse from solana RPC, method: '%q', code: %d\n", GetBlockMethodName, res.StatusCode)
	var blockResponse GetBlockResponseBody
	readResponseBody(res.Body, blockResponse)
	return blockResponse
}

func readResponseBody[T any](closer io.ReadCloser, t T) {
	body, err := io.ReadAll(closer)
	if err != nil {
		config.Log.Printf("Error reading response body: %q\n", err.Error())
	}
	err1 := json.Unmarshal(body, &t)
	if err1 != nil {
		config.Log.Printf("Error when convert body to json, error : %q\n", err1.Error())
	}
}

func toJsonIoReader(v any) io.Reader {
	res, err := json.Marshal(v)
	if err != nil {
		config.Log.Printf("Error when convert value: %q to json, error: %q\n", v, err.Error())
	}
	return bytes.NewBuffer(res)
}

type RpcCallWithoutParameters struct {
	JsonRpc string `json:"jsonRpc"`
	Id int `json:"id"`
	Method string `json:"method"`
}

type RpcCallWithParameters struct {
	JsonRpc string `json:"jsonRpc"`
	Id int `json:"id"`
	Method string `json:"method"`
	Params []interface{} `json:"params"`
}

type GetSlotResponseBody struct {
	JsonRpc string `json:"jsonRpc"`
	Result int64 `json:"result"`
	Id int `json:"id"`
}

type GetBlockResponseBody struct {
	JsonRpc string `json:"jsonRpc"`
	GetBlockResponseResultBody  `json:"result"`
	Id int `json:"id"`
}

type GetBlockParamsBody struct {
	Enconding string `json:"encoding"`
	TransactionVersion int `json:"maxSupportedTransactionVersion"`
	TransactionDetails string `json:"transactionDetails"`
	Rewards bool `json:"rewards"`
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

var defaultGetBlockParamsBody = GetBlockParamsBody {
	Enconding: "jsonParsed",
	TransactionVersion: 0,
	TransactionDetails: "full",
	Rewards: true,
}
