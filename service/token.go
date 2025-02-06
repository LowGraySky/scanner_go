package service

import (
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

type RealTokenFetcher struct {
	JupiterCaller JupiterCaller
}

var answers = make(map[string]model.TokenInfo)

func (tf *RealTokenFetcher) GetTokenInfo(address string) (model.TokenInfo, error) {
	if tf.isExists(address) {
        val, _ := answers[address]
        config.Log.Infof("Got token information by token: %s in cache, use it", val.Symbol)
        return val, nil
	} else {
		res, err := tf.JupiterCaller.GetToken(address)
		if err != nil {
			config.Log.Errorf("Error when request to jupiter token info, error: %q", err.Error())
			return model.TokenInfo{}, err
		}
		answers[address] = res
		return res, nil
	}
}

func (tf *RealTokenFetcher) isExists(address string) bool {
    _, exists := answers[address]
    return exists
}