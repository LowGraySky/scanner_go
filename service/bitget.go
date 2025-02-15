package service

import (
	"errors"
	"fmt"
	"net/http"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const bitgetBaseUrl = "https://api.bitget.com/api/v2/mix/market/ticker"

type RealBitgetCaller struct {}

func (bgc *RealBitgetCaller) GetToken(symbol string) (model.GateResponse, error)  {
	res, err := http.Get(bitgetBaseUrl + "?productType=USDT-FUTURES&symbol=" + symbol)
	if err != nil {
		config.Log.Errorf("Error when request to BITGET token info by symbol: %s, error: %q", symbol, err.Error())
		return model.GateResponse{}, err
	}
	config.Log.Infof("Got reponse from BITGET token info by symbol: '%s', code: %d", symbol, res.StatusCode)
	if res.StatusCode != 200 {
		config.Log.Errorf("Unsuccess response code: %d from BITGET by symbol: %s", res.StatusCode, symbol)
		return model.GateResponse{}, errors.New("token with symbol is not exitst on BITGET")
	}
	var response model.GateResponse
	readResponseBody(res.Body, &response)
	config.Log.Infof("Response body: " + fmt.Sprintf("Code: %s, Message: %s", response.Code, response.Message))
	return response, nil
}