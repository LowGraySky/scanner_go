package model

const (
	notExistsResponseCode = 1001
)

type MexcTokenInfoResponse[T any] struct {
	Success bool `json:"success"`
	Code    uint `json:"code"`
	Data    T    `json:"data"`
}

func (mtir *MexcTokenInfoResponse[T]) IsSuccess() bool {
	return mtir.Success == true && mtir.Code == 0
}

func (mtir *MexcTokenInfoResponse[T]) IsNotExistst() bool {
	return !mtir.IsSuccess() && mtir.Code == notExistsResponseCode
}
