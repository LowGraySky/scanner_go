package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"
	"web3.kz/solscan/config"
	"web3.kz/solscan/model"
)

type RealTokenFetcher struct {
	JupiterCaller   JupiterCaller
	MexcCaller      MexcCaller
	GateCaller      GateCaller
	TokenRepository TokenRepository
}

func (tf *RealTokenFetcher) GetTokenInfo(address string) (model.TokenInfo, error) {
	exists, symbol, _ := tf.TokenRepository.JupiterTokenByAddress(address)
	if !exists {
		res, err1 := tf.JupiterCaller.GetToken(address)
		if err1 != nil {
			config.Log.Errorf("Error when request to jupiter token by address: %s info, error: %q", address, err1.Error())
			return model.TokenInfo{}, err1
		}
		err2 := tf.TokenRepository.SaveJupiterToken(address, res.Symbol)
		if err2 != nil {
			config.Log.Errorf("Error when save info by symbol %s, error: %q", symbol, err2.Error())
			return model.TokenInfo{}, err2
		}
		return res, nil
	}
	config.Log.Infof("Token by address: %s EXISTS in repository, symbol: %s", address, symbol)
	return model.TokenInfo{Symbol: symbol}, nil
}

func (tf *RealTokenFetcher) ExchangeTokenInfo(symbol string) model.Token {
	exists, token, _ := tf.TokenRepository.ExchangeTokenInfo(symbol)
	if !exists {
		res, err := tf.fetchInfoAbountToken(symbol)
		if err != nil {
			config.Log.Errorf("Error when getting information about token, error: %q", err.Error())
			return model.Token{}
		}
		return res
	}
	return token
}

func (tf *RealTokenFetcher) fetchInfoAbountToken(symbol string) (model.Token, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var mexc sql.NullBool
	var gate sql.NullBool
	go func() {
		defer wg.Done()
		ex, err := tf.IsExistsOnMexc(symbol)
		if err != nil {
			mexc = sql.NullBool{
				Valid: false,
			}
		} else {
			mexc = sql.NullBool{
				Bool: ex,
				Valid: true,
			}
		}
	}()
	go func() {
		defer wg.Done()
		ex, err := tf.IsExistsOnGate(symbol)
		if err != nil {
			gate = sql.NullBool{
				Valid: false,
			}
		} else {
			gate = sql.NullBool{
				Bool: ex,
				Valid: true,
			}
		}
	}()
	wg.Wait()
	token := model.Token{
		Symbol: symbol,
		IsExistsMexc: mexc,
		IsExistsBitget: sql.NullBool{
			Valid: false,
		},
		IsExistsGate: gate,
	}
	err := tf.TokenRepository.UpdateExchangeTokenInfo(token)
	if err != nil {
		return model.Token{}, err
	}
	return token, nil

}

func (tf *RealTokenFetcher) IsExistsOnMexc(symbol string) (bool, error) {
	mexcSymbol := symbol + "_USDT"
	resp, err := tf.MexcCaller.GetToken(mexcSymbol)
	if err != nil {
		config.Log.Errorf("Error when find token %s on MEXC, error: %q", symbol, err.Error())
		return false, err
	}
	if resp.IsNotExistst() {
		config.Log.Warnf("Not found token: %s on MEXC", symbol)
		return false, nil
	}
	if resp.IsSuccess() {
		config.Log.Infof("Found token: %s on MEXC", symbol)
		return true, nil
	} else {
		config.Log.Warnf("Not found token: %s on MEXC, data: %s", symbol, fmt.Sprintf("Success: %s, Code: %d", resp.Success, resp.Code))
		return false, errors.New("error to call")
	}
}

func (tf *RealTokenFetcher) IsExistsOnBitget(symbol string) (bool, error) {
	// TODO <--
	return false, nil
}

func (tf *RealTokenFetcher) IsExistsOnGate(symbol string) (bool, error) {
	symbolGate := strings.ToUpper(symbol) + "_USDT"
	_, err := tf.GateCaller.GetToken(symbolGate)
	if err != nil {
		config.Log.Errorf("Not found token: %s on GATE", symbol)
		return false, err
	}
	config.Log.Infof("Found token: %s on GATE", symbol)
	return true, nil
}
