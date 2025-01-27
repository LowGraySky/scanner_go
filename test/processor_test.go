package test

import (
	"github.com/stretchr/testify/mock"
	"testing"
	"web3.kz/solscan/model"
	"web3.kz/solscan/service"
)

func TestProcess(t *testing.T) {
	mockCaller := new(MockSolanaCaller)
	mockAnalyser := new(MockAnalyser)
	processor := service.RealProcessor{
		Analyser:     mockAnalyser,
		Serialiser:   &service.RealSerializer{},
		SolanaCaller: mockCaller,
	}
	getSlotResponse := model.GetSlotResponseBody{
		JsonRpc: "2.0",
		Result:  2,
		Id:      1,
		Error: model.Error{
			Code:    0,
			Message: "",
		},
	}
	getBlockResponseBody := ReadBlockResponseFromFile()
	mockCaller.On("GetSlot").Return(getSlotResponse, nil)
	mockCaller.On("GetBlock", getSlotResponse.Result).Return(getBlockResponseBody, nil)
	mockAnalyser.On("Analyse", getSlotResponse.Result, mock.Anything).Return(make([]model.Transaction, 0))

	processor.Process()

	expectation := mockCaller.AssertExpectations(t)

	if !expectation {
		t.Errorf("")
	}
}

type MockAnalyser struct {
	mock.Mock
}

func (ma *MockAnalyser) Analyse(slotNumber uint, transactions []model.Transaction) []model.Transaction {
	args := ma.Called(slotNumber, transactions)
	return args.Get(0).([]model.Transaction)
}

type MockSolanaCaller struct {
	mock.Mock
}

func (m *MockSolanaCaller) GetSlot() (model.GetSlotResponseBody, error) {
	args := m.Called()
	return args.Get(0).(model.GetSlotResponseBody), args.Error(1)
}
func (m *MockSolanaCaller) GetBlock(slotNumber uint) (model.GetBlockResponseBody, error) {
	args := m.Called(slotNumber)
	return args.Get(0).(model.GetBlockResponseBody), args.Error(1)
}

