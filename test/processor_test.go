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
	mockTokenFetcher := new(MockTokenFethcer)
	processor := service.RealProcessor{
		Analyser:       mockAnalyser,
		Serialiser:     &service.RealSerializer{
			TokenFetcher: mockTokenFetcher,
		},
		TokenFetcher: mockTokenFetcher,
		SolanaCaller:   mockSolanaCaller,
		RedisCaller:    mockRedisCaller,
		TelegramCaller: mockTelergamCaller,
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
	mockAnalyser.On("Analyse", getSlotResponse.Result, mock.Anything).Return(getBlockResponseBody.Result.Transactions)
	mockTelergamCaller.On("SendMessage", mock.Anything).Return(&gotgbot.Message{MessageId: 1}, nil)
	mockRedisCaller.On("Set", mock.Anything, "DG5XkaGPGVywyygWxWBCLycmh7QDW6J6pToJeQqCFwPr", int64(1), mock.Anything).Return(nil)
	mockTokenFetcher.On("GetTokenInfo", "DVZrNS9fctrrDmhZUZAu6p63xU6d9cqYxRRhJbtJ4z8G").Return(model.TokenInfo{
		DailyVolume: 1.0,
		Symbol: "ROSS",
	}, nil)
	mockTokenFetcher.On("IsExistsOnMexc", mock.Anything).Return(true)
	mockTokenFetcher.On("IsExistsOnGate", mock.Anything).Return(true)
	mockTokenFetcher.On("IsExistsOnBitget", mock.Anything).Return(true)

	processor.Process()

	expectation := mockSolanaCaller.AssertExpectations(t)

	if !expectation {
		t.Errorf("")
	}
}

func TestProcessCloseOrder(t *testing.T) {
	mockCaller := new(MockSolanaCaller)
	mockAnalyser := new(MockAnalyser)
	mockRedisCaller := new(MockRedisCaller)
	mockTelegramCaller := new(MockTelegramCaller)
	processor := service.RealProcessor{
		Analyser:     mockAnalyser,
		Serialiser:   &service.RealSerializer{},
		SolanaCaller: mockCaller,
		RedisCaller: mockRedisCaller,
		TelegramCaller: mockTelegramCaller,
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
	var messageId int64 = 2
	getBlockResponseBody := ReadBlockResponseFromFile("files/test_data_close_order.txt")
	mockCaller.On("GetSlot").Return(getSlotResponse, nil)
	mockCaller.On("GetBlock", getSlotResponse.Result).Return(getBlockResponseBody, nil)
	mockAnalyser.On("Analyse", getSlotResponse.Result, mock.Anything).Return(getBlockResponseBody.Result.Transactions)
	mockRedisCaller.On("Get", mock.Anything, "EDRW6cwEGdvAuyATqnaYG3meTqHT3FNLM3KS7N69Pd1F").Return(messageId, nil)
	mockTelegramCaller.On("SendReplyMessage", "DCA closed by user", messageId).Return(nil)

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
	args := tc.Called(message)
	return args.Get(0).(*gotgbot.Message), args.Error(1)
}

func (tc *MockTelegramCaller) SendReplyMessage(message string, messageId int64) error {
	args := tc.Called(message, messageId)
	return args.Error(0)
}

type MockRedisCaller struct {
	mock.Mock
}

func (mr *MockRedisCaller) Get(ctx context.Context, key string) (int64, error) {
	args := mr.Called(ctx, key)
	return args.Get(0).(int64), args.Error(1)
}

func (mr *MockRedisCaller) Set(ctx context.Context, key string, value int64, expiration time.Duration) error {
	args := mr.Called(ctx, key, value, expiration)
	return args.Error(0)
}

type MockTokenFethcer struct {
	mock.Mock
}

func (mtf *MockTokenFethcer) GetTokenInfo(address string) (model.TokenInfo, error) {
	args := mtf.Called(address)
	return args.Get(0).(model.TokenInfo), args.Error(1)
}

func (mtf *MockTokenFethcer) IsExistsOnMexc(symbol string) bool {
	args := mtf.Called(symbol)
	return args.Bool(0)
}

func (mtf *MockTokenFethcer) IsExistsOnGate(symbol string) bool {
	args := mtf.Called(symbol)
	return args.Bool(0)
}

func (mtf *MockTokenFethcer) IsExistsOnBitget(symbol string) bool {
	args := mtf.Called(symbol)
	return args.Bool(0)
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
