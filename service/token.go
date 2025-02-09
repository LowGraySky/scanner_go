package service

import (
	"fmt"
	"strings"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

type RealTokenFetcher struct {
	JupiterCaller JupiterCaller
	MexcCaller    MexcCaller
	GateCaller    GateCaller
}

var jupiterAnswers = make(map[string]model.TokenInfo)

func (tf *RealTokenFetcher) GetTokenInfo(address string) (model.TokenInfo, error) {
	if tf.isExistsInAnswerMap(address) {
		val, _ := jupiterAnswers[address]
		config.Log.Infof("Got token information by token: %s in cache, use it", val.Symbol)
		return val, nil
	} else {
		res, err := tf.JupiterCaller.GetToken(address)
		if err != nil {
			config.Log.Errorf("Error when request to jupiter token info, error: %q", err.Error())
			return model.TokenInfo{}, err
		}
		jupiterAnswers[address] = res
		return res, nil
	}
}

func (tf *RealTokenFetcher) isExistsInAnswerMap(address string) bool {
	_, exists := jupiterAnswers[address]
	return exists
}

func (tf *RealTokenFetcher) IsExistsOnMexc(symbol string) bool {
	mexcSymbol := symbol + "_USDT"
	resp, err := tf.MexcCaller.GetToken(mexcSymbol)
	if err != nil {
		config.Log.Errorf("Error when find token %s on MEXC, error: %q", symbol, err.Error())
		return false
	}
	if resp.IsNotExistst() {
		config.Log.Warnf("NOT FOUND token: %s on MEXC", symbol)
		return false
	}
	if resp.IsSuccess() {
		config.Log.Infof("FOUND token: %s on MEXC", symbol)
		return true
	} else {
		config.Log.Warnf("NOT FOUND token: %s on MEXC, data: %s", symbol, fmt.Sprintf("Success: %s, Code: %d", resp.Success, resp.Code))
		return false
	}
}

func (tf *RealTokenFetcher) IsExistsOnBitget(symbol string) bool {
	// TODO <--
	return false
}

func (tf *RealTokenFetcher) IsExistsOnGate(symbol string) bool {
	symbolGate := strings.ToUpper(symbol) + "_USDT"
	_, err := tf.GateCaller.GetToken(symbolGate)
	if err != nil {
		config.Log.Errorf("NOT FOUND token: %s on GATE", symbol)
		return false
	}
	config.Log.Infof("FOUND token: %s on GATE", symbol)
	return true
}
