package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BaseUrl = "https://api.mainnet-beta.solana.com"

const (
	ApplicationJsonContentType = "application/json"
	JsonRpcValue = "2.0"
	IdValue = 1
	GetSlotMethodName = "getSlot"
	GetBlockMethodName = "getBlock"
	DefaultGetBlockParamsBody =  GetBlockParamsBody {
		Enconding: "json"
		TransactionVersion: 0
		TransactionDetails: "full"
		Rewards: ""
	}
)

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
	Result  `json:"result"`
	Id int `json:"id"`
}

type GetBlockParamsBody struct {
	Enconding string `json:"encoding"`
	TransactionVersion int `json:"maxSupportedTransactionVersion"`
	TransactionDetails string `json:"transactionDetails"`
	Rewards bool `json:"rewards"`
}

"encoding": "json",
"maxSupportedTransactionVersion":0,
  "transactionDetails":"full",
  "rewards":false

type GetBlockResponseResultBody struct {

}

type Transaction

func GetSlot() GetSlotResponseBody {
	rpcCall := RpcCallWithoutParameters {
		JsonRpc: JsonRpcValue,
		Id: IdValue,
		Method: GetSlotMethodName,
	}
	res, err := http.Post(BaseUrl, ApplicationJsonContentType, toJsonIoReader(rpcCall))
	if err != nil {
		fmt.Printf("Error when request to solana RPC, method: '%q', error: %q\n", GetSlotMethodName, err.Error())
	}
	fmt.Printf("Got reponse from solana RPC, method: '%q', code: %d\n", GetSlotMethodName, res.StatusCode)
	var slotResponse GetSlotResponseBody
	readResponseBody(res.Body, slotResponse)
	return slotResponse
}

func GetBlock(slotNumber int) Get {
	rpcCall := RpcCallWithParameters {
		JsonRpc: JsonRpcValue,
		Id: IdValue,
		Method: GetBlockMethodName,
		Params: []interface{}{ slotNumber,  }
	}
	res, err := http.Post(BaseUrl, ApplicationJsonContentType, toJsonIoReader(rpcCall))
	if err != nil {
		fmt.Printf("Error when request to solana RPC, method: '%q', error: %q\n", GetBlockMethodName, err.Error())
	}
	fmt.Printf("Got reponse from solana RPC, method: '%q', code: %d\n", GetBlockMethodName, res.StatusCode)
	var slotResponse GetSlotResponseBody
	readResponseBody(res.Body, slotResponse)
	return slotResponse
}

func readResponseBody[T any](closer io.ReadCloser, t T) {
	body, err := io.ReadAll(closer)
	if err != nil {
		fmt.Printf("Error reading response body: %q\n", err.Error())
	}
	err1 := json.Unmarshal(body, t)
	if err1 != nil {
		fmt.Printf("Error when convert body to json, error : %q\n", err1.Error())
	}
}

func toJsonIoReader(v any) io.Reader {
	res, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("Error when convert value: %q to json, error: %q\n", v, err.Error())
	}
	return bytes.NewBuffer(res)
}