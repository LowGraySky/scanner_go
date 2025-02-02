package service

import (
    "net/http"
    "web3.kz/solscan/config"
    "web3.kz/solscan/model"
)

const (
    JupiterBaseUrl = "https://api.jup.ag"
    TokenInfoPath  = "/tokens/v1/token/"
)

var answers = make(map[string]model.TokenInfo)

type RealJupiterCaller struct{}

func (jc *RealJupiterCaller) GetToken(address string) (model.TokenInfo, error) {
    if isExists(address) {
        val, _ := answers[address]
        return val, nil
    } else {
        response, err := fetchData(address)
        if err != nil {
            return model.TokenInfo{}, err
        }
        answers[address] = response
        return response, nil
    }
}

func fetchData(address string) (model.TokenInfo, error) {
    res, err := http.Get(JupiterBaseUrl + TokenInfoPath + address)
    if err != nil {
        config.Log.Errorf("Error when request to jupiter token info, error: %q", err.Error())
        return model.TokenInfo{}, err
    }
    config.Log.Infof("Got response from jupiter token info, code: %d", res.StatusCode)
    var tokenInfo model.TokenInfo
    readResponseBody(res.Body, &tokenInfo)
    return tokenInfo, nil
}

func isExists(address string) bool {
    _, exists := answers[address]
    return exists
}
