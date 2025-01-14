package service

import (
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


func GetSlot() {
	rpcCall := RpcCallWithoutParameters {
		JsonRpc: JsonRpcValue,
		Id: IdValue,
		Method: GetSlotMethodName,
	}
	res, err := http.Post(BaseUrl, ApplicationJsonContentType, )
	if err != nil {
		fmt.Printf("Error when request to solana RPC, method: '%q', error: %q\n", GetSlotMethodName, err.Error())
		return
	}
	fmt.Printf("Got reponse from solana RPC, method: '%q', code: %d\n", GetBlockMethodName, res.StatusCode)
	bytes := readResponseBody(res.Body)
	var slotResponse GetSlotResponseBody
	resBody, err := fromJson(bytes, &slotResponse)

}


func GetBlock() {

}


func readResponseBody(closer io.ReadCloser, interface{}) any {
	body, err := io.ReadAll(closer)
	if err != nil {
		fmt.Printf("Error reading response body: %q\n", err.Error())
	}
	err1 := json.Unmarshal(body, )
	if err1 != nil {
		fmt.Printf("Error when convert body to json, error : %q\n", err.Error())
	}
	return res
}

func toJson(v any) []byte {
	res, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("Error when convert value: %q to json, error: %q\n", v, err.Error())
	}
	return res
}