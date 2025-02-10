package model

import "database/sql"

type JupiterToken struct {
	Contract string
	Symbol   string
}

type Token struct {
	Symbol         string
	IsExistsGate   sql.NullBool
	IsExistsBitget sql.NullBool
	IsExistsMexc   sql.NullBool
}
