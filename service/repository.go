package service

import (
	"database/sql"
	"web3.kz/solscan/model"
)

type RealTokenRepository struct {
	Db sql.DB
}

func (tr *RealTokenRepository) JupiterTokenByAddress(address string) (bool, string, error) {
	r := tr.Db.QueryRow("select symbol from dca.jupiter_token where address = $1", address)
	var symbol string
	err := r.Scan(&symbol)
	if err != nil {
		return false, "", err
	}
	return true, symbol, err
}

func (tr *RealTokenRepository) SaveJupiterToken(address string, symbol string) error {
	_, err := tr.Db.Exec("insert into dca.jupiter_token values ($1, $2)", address, symbol)
	if err != nil {
		return err
	}
	return nil
}

func (tr *RealTokenRepository) ExchangeTokenInfo(symbol string) (bool, model.Token, error) {
	r := tr.Db.QueryRow("select * from dca.token_info where symbol = $1", symbol)
	var symb string
	var isExistsOnMexc sql.NullBool
	var isExistsOnGate sql.NullBool
	var isExistsOnBitget sql.NullBool

	err := r.Scan(&symb, &isExistsOnMexc, &isExistsOnBitget, &isExistsOnGate)
	if err != nil {
		return false, model.Token{}, err
	}
	return true,model.Token{
		Symbol:         symb,
		IsExistsMexc:   isExistsOnMexc,
		IsExistsGate:   isExistsOnGate,
		IsExistsBitget: isExistsOnBitget,
	}, nil
}

func (tr *RealTokenRepository) UpdateExchangeTokenInfo(token model.Token) error {
	_, err := tr.Db.Exec(
		"update dca.token_info(is_exists_on_mexc, is_exists_on_bitget, is_exists_on_gate) set ($1, $2, $3) where symbol = $4",
		token.IsExistsMexc,
		token.IsExistsBitget,
		token.IsExistsGate,
		token.Symbol)
	if err != nil {
		return err
	}
	return nil
}
