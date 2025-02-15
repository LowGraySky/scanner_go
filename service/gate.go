package service

import (
	"errors"
	"fmt"
	"net/http"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

const gateBaseUrl = "https://api.gateio.ws/api/v4/futures/usdt/contracts/"

type RealGateCaller struct {}

func (gc *RealGateCaller) GetToken(symbol string) (model.GateTokenInfoResponse, error)  {
	res, err := http.Get(gateBaseUrl + symbol)
	if err != nil {
		config.Log.Errorf("Error when request to GATE token info by symbol: %s, error: %q", symbol, err.Error())
		return model.GateTokenInfoResponse{}, err
	}
	config.Log.Infof("Got reponse from GATE token info by symbol: '%s', code: %d", symbol, res.StatusCode)
	if res.StatusCode != 200 {
		config.Log.Errorf("Unsuccess repsonse code: %d from GATE by symbol: %s", res.StatusCode, symbol)
		return model.GateTokenInfoResponse{}, errors.New("token with symbol is not exitst on Gate")
	}
	var response model.GateTokenInfoResponse
	readResponseBody(res.Body, &response)
	config.Log.Infof("Response body: " + fmt.Sprintf("Name: %s", response.Name))
	return response, nil
}