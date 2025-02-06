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

type RealJupiterCaller struct{}

func (jc *RealJupiterCaller) GetToken(address string) (model.TokenInfo, error) {
    res, err := http.Get(JupiterBaseUrl + TokenInfoPath + address)
    if err != nil {
        return model.TokenInfo{}, err
    }
    config.Log.Infof("Got response from jupiter token info, code: %d", res.StatusCode)
    var response model.TokenInfo
    readResponseBody(res.Body, &response)
    return response, nil
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
