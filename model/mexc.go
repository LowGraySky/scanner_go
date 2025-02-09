package model

const (
	notExistsResponseCode = 1001
)

type MexcTokenInfoResponse struct {
	Success bool   `json:"success"`
	Code    uint   `json:"code"`
	Data    string `json:"data"`
}

func (mtir *MexcTokenInfoResponse) IsSuccess() bool {
	return mtir.Success == true && mtir.Code == 0
}

func (mtir *MexcTokenInfoResponse) IsNotExistst() bool {
	return !mtir.IsSuccess() && mtir.Code == notExistsResponseCode
}
