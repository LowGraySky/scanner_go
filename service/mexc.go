package service

import (
	"fmt"
	"net/http"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const mexcBaseUrl = "https://contract.mexc.com/api/v1/contract/detail"

type RealMexcCaller struct{}

func (mc *RealMexcCaller) GetToken(symbol string) (model.MexcTokenInfoResponse[string], error) {
	res, err := http.Get(mexcBaseUrl + "/?symbol=" + symbol)
	if err != nil {
		config.Log.Errorf("Error when request to MEXC token info by symbol: %s, error: %q", symbol, err.Error())
		return model.MexcTokenInfoResponse[string]{}, err
	}
	config.Log.Infof("Got reponse from MEXC token info by symbol: '%q', code: %d", symbol, res.StatusCode)
	var response model.MexcTokenInfoResponse[string]
	readResponseBody(res.Body, &response)
	config.Log.Infof("Response body: %s", fmt.Sprintf("Success: %s, Code: %d", response.Success, response.Code))
	return response, nil
}
