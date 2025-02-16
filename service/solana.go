package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const (
	SolanaNodeBaseUrl          = "https://solana-mainnet.g.alchemy.com/v2/AT8Nj54Fcwv-ezshyh3DgkJv7CgFCE_a"
	ApplicationJsonContentType = "application/json"
	JsonRpcValue = "2.0"
	IdValue = 1
	GetSlotMethodName = "getSlot"
	GetBlockMethodName = "getBlock"
)

var defaultGetBlockParamsBody = model.GetBlockParamsBody {
	Enconding: "jsonParsed",
	TransactionVersion: 0,
	Rewards: false,
}

type RealSolanaCaller struct {}

func (sc *RealSolanaCaller) GetSlot() (model.GetSlotResponseBody, error) {
	rpcCall := model.RpcCallWithoutParameters {
		JsonRpc: JsonRpcValue,
		Id: IdValue,
		Method: GetSlotMethodName,
	}
	reader := toJsonIoReader(rpcCall)
	res, err := http.Post(SolanaNodeBaseUrl, ApplicationJsonContentType, reader)
	defer reader.Close()
	if err != nil {
		config.Log.Errorf("Error when request to solana RPC, method: '%q', error: %q", GetSlotMethodName, err.Error())
		return model.GetSlotResponseBody{}, err
	}
	config.Log.Infof("Got reponse from solana RPC, method: '%q', code: %d", GetSlotMethodName, res.StatusCode)
	var slotResponse model.GetSlotResponseBody
	readResponseBody(res.Body, &slotResponse)
	defer res.Body.Close()
	config.Log.Infof("Response body: %q", slotResponse)
	return slotResponse, nil
}

func (sc *RealSolanaCaller) GetBlock(slotNumber uint) (model.GetBlockResponseBody, error) {
	rpcCall := model.RpcCallWithParameters {
		JsonRpc: JsonRpcValue,
		Id: IdValue,
		Method: GetBlockMethodName,
		Params: []interface{}{ slotNumber, defaultGetBlockParamsBody },
	}
	reader := toJsonIoReader(rpcCall)
	res, err := http.Post(SolanaNodeBaseUrl, ApplicationJsonContentType, reader)
	defer reader.Close()
	if err != nil {
		config.Log.Errorf("Error when request to solana RPC, method: '%q', error: %q\n", GetBlockMethodName, err.Error())
		return model.GetBlockResponseBody{}, err
	}
	config.Log.Infof("Got reponse from solana RPC, method: '%q', code: %d", GetBlockMethodName, res.StatusCode)
	var blockResponse model.GetBlockResponseBody
	readResponseBody(res.Body, &blockResponse)
	defer res.Body.Close()
	return blockResponse, nil
}

func readResponseBody[T any](closer io.ReadCloser, t T) {
	body, err := io.ReadAll(closer)
	if err != nil {
		config.Log.Errorf("Error reading response body: %q\n", err.Error())
	}
	err1 := json.Unmarshal(body, &t)
	if err1 != nil {
		config.Log.Errorf("Error when convert body to json, error : %q\n", err1.Error())
	}
}

func toJsonIoReader(v any) io.ReadCloser {
	res, err := json.Marshal(v)
	if err != nil {
		config.Log.Errorf("Error when convert value: %q to json, error: %q\n", v, err.Error())
	}
	 reader := bytes.NewReader(res)
	 return io.NopCloser(reader)
}
