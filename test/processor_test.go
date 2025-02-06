package test

import (
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"web3.kz/solscan/model"
	"web3.kz/solscan/service"
)

func TestProcessOpenOrder(t *testing.T) {
	mockSolanaCaller := new(MockSolanaCaller)
	mockRedisCaller := new(MockRedisCaller)
	mockAnalyser := new(MockAnalyser)
	mockTelergamCaller := new(MockTelegramCaller)
	processor := service.RealProcessor{
		Analyser:     mockAnalyser,
		Serialiser:   &service.RealSerializer{},
		SolanaCaller: mockSolanaCaller,
		RedisCaller: mockRedisCaller,
		TelegramCaller : mockTelergamCaller,
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
	getBlockResponseBody := ReadBlockResponseFromFile("files/test_data_open_order.txt")
	mockSolanaCaller.On("GetSlot").Return(getSlotResponse, nil)
	mockSolanaCaller.On("GetBlock", getSlotResponse.Result).Return(getBlockResponseBody, nil)
	mockAnalyser.On("Analyse", getSlotResponse.Result, mock.Anything).Return(make([]model.Transaction, 0))

	processor.Process()

	expectation := mockSolanaCaller.AssertExpectations(t)

	if !expectation {
		t.Errorf("")
	}
}

func TestProcessCloseOrder(t *testing.T)  {
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
	getBlockResponseBody := ReadBlockResponseFromFile("files/test_data_close_order.txt")
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

type MockTelegramCaller struct {
	mock.Mock
}

func (tc *MockTelegramCaller) SendMessage(message string) (*gotgbot.Message, error) {
	args := tc.Called()
}

func (tc *MockTelegramCaller) SendReplyMessage(message string, messageId int64) error {
	args := tc.Called()
}

type MockRedisCaller struct {
	mock.Mock
}

func (mr *MockRedisCaller) Get(ctx context.Context, key string) (int64, error) {
	args := mr.Called()
}

func (mr *MockRedisCaller) Set(ctx context.Context, key string, value int64, expiration time.Duration) error {
	args := mr.Called()
	
}


type MockSolanaCaller struct {
	mock.Mock
}

func (ms *MockSolanaCaller) GetSlot() (model.GetSlotResponseBody, error) {
	args := ms.Called()
	return args.Get(0).(model.GetSlotResponseBody), args.Error(1)
}
func (ms *MockSolanaCaller) GetBlock(slotNumber uint) (model.GetBlockResponseBody, error) {
	args := ms.Called(slotNumber)
	return args.Get(0).(model.GetBlockResponseBody), args.Error(1)
}
