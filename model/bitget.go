package model

const (
  successCode = "0000"
  successMessage = "success"
  tokenNotExistCode = "40034"
)

type GateResponse struct {
  Code string `json:"code"`
  Message string `json:"msg"`
}

func (gr *GateResponse) IsSuccess() bool {
  return gr.Code == successCode && gr.Message == successMessage
}

func (gr *GateResponse) IsTokenNotExists() bool {
  return gr.Code == tokenNotExistCode
}
