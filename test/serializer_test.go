package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"web3.kz/solscan/model"
	"web3.kz/solscan/service"
)

func TestSerializer(t *testing.T) {
	mockTokenRepository := new(MockTokenRepository)
	serializer := service.RealSerializer{
		TokenFetcher: &service.RealTokenFetcher{
			TokenRepository: mockTokenRepository,
			JupiterCaller: &service.RealJupiterCaller{},
		},
	}
	mockTokenRepository.On("JupiterTokenByAddress", mock.Anything).Return(true, "Ross", nil)
	transactions := ReadBlockResponseFromFile("files/test_data_open_order.txt").Result.Transactions

	actual, _ := serializer.Serialize(1, transactions[0])

	res := actual

	dke := assert.Equal(t, res.DcaKey, "DG5XkaGPGVywyygWxWBCLycmh7QDW6J6pToJeQqCFwPr")
	ame := assert.Equal(t, res.InstructionData.InAmount, "8272121570535")
	ampe := assert.Equal(t, res.InstructionData.InAmountPerCycle, "33088486282")
	cfe := assert.Equal(t, res.InstructionData.CycleFrequency, "60")
	tkcae := assert.Equal(t, res.Token, "DVZrNS9fctrrDmhZUZAu6p63xU6d9cqYxRRhJbtJ4z8G")
	tksymbe := assert.Equal(t, res.TokenSymbol, "Ross")
	opee := assert.Equal(t, res.Operation.String(), "SELL")
	usae := assert.Equal(t, res.User, "7DiaCzvNmMcA7z8J3McC3VaUDJJTdKPQCd9YTAThSTaY")
	se := assert.Equal(t, res.Signature, "4LFqsgwRWWsQpcy3P9ZxmxQo8dX5fob8oU9Zs71VbSWbt8rnc7ovXdnCx9U2N3khxZogLCpyPbsKiZ5Nsr1GYv7k")

	r := ame && ampe && cfe && tkcae && tksymbe && opee && usae && se && dke
	if !r {
		t.Error("")
	}
}

func TestSerializerCorrectSymbolAndOperation1(t *testing.T)  {
	mockTokenRepository := new(MockTokenRepository)
	serializer := service.RealSerializer{
		TokenFetcher: &service.RealTokenFetcher{
			JupiterCaller: &service.RealJupiterCaller{},
			TokenRepository: mockTokenRepository,
		},
	}
	mockTokenRepository.On("JupiterTokenByAddress", mock.Anything).Return(true, "TRUMP", nil)
	transactions := ReadBlockResponseFromFile("files/test_data_s_and_op1.txt ").Result.Transactions

	actual, _ := serializer.Serialize(1, transactions[0])

	res := actual

	dke := assert.Equal(t, res.DcaKey, "4FY3uWuRN16Tnq4GQNrTnmCmwYvBFsKZc2SKv2Qum6Q4")
	ame := assert.Equal(t, res.InstructionData.InAmount, "23048316276")
	ampe := assert.Equal(t, res.InstructionData.InAmountPerCycle, "1152415814")
	cfe := assert.Equal(t, res.InstructionData.CycleFrequency, "60")
	tkcae := assert.Equal(t, res.Token, "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN")
	tksymbe := assert.Equal(t, res.TokenSymbol, "TRUMP")
	opee := assert.Equal(t, res.Operation.String(), "SELL")
	usae := assert.Equal(t, res.User, "Af1gwGFchrcKojHrmL9yhcqC89kURuahCmYvCBnq9oKU")
	se := assert.Equal(t, res.Signature, "3tiXLgrpgTNUkTrspZ8HzfKhEy5Ngw8ZwewKrGNyS23zN68XTh9sa5cL6rN9fvccW67i1UoQQjaN6djpnGXfMhNX")

	r := ame && ampe && cfe && tkcae && tksymbe && opee && usae && se && dke
	if !r {
		t.Error("")
	}
}

func TestSerializerCorrectSymbolAndOperation2(t *testing.T)  {
	mockTokenRepository := new(MockTokenRepository)
	serializer := service.RealSerializer{
		TokenFetcher: &service.RealTokenFetcher{
			JupiterCaller: &service.RealJupiterCaller{},
			TokenRepository: mockTokenRepository,
		},
	}
	mockTokenRepository.On("JupiterTokenByAddress", mock.Anything).Return(true, "jailstool", nil)
	transactions := ReadBlockResponseFromFile("files/test_data_s_and_op2.txt").Result.Transactions

	actual, _ := serializer.Serialize(1, transactions[0])

	res := actual

	dke := assert.Equal(t, res.DcaKey, "AASzrsLMPCDDtw1zzJUiYrrc198CaUtxaQQjc7L48NMs")
	ame := assert.Equal(t, res.InstructionData.InAmount, "500000000000")
	ampe := assert.Equal(t, res.InstructionData.InAmountPerCycle, "1858736059")
	cfe := assert.Equal(t, res.InstructionData.CycleFrequency, "60")
	tkcae := assert.Equal(t, res.Token, "AxriehR6Xw3adzHopnvMn7GcpRFcD41ddpiTWMg6pump")
	tksymbe := assert.Equal(t, res.TokenSymbol, "jailstool")
	opee := assert.Equal(t, res.Operation.String(), "BUY")
	usae := assert.Equal(t, res.User, "BNJWKsPtEgpRz7XDyKXbKmmBQeZ8UubGR81W9fM3Y5uC")
	se := assert.Equal(t, res.Signature, "3QYNmrAWAJq6ka93msKGgGgLfUJgenLe3n3BpksUqa57HzsByhU8zsZ3YkH9H6Q85kiijtcW7LPV18zUC1Ya1smM")

	r := ame && ampe && cfe && tkcae && tksymbe && opee && usae && se && dke
	if !r {
		t.Error("")
	}
}

func TestSerializerCorrectSymbolAndOperation3(t *testing.T)  {
	mockTokenRepository := new(MockTokenRepository)
	serializer := service.RealSerializer{
		TokenFetcher: &service.RealTokenFetcher{
			JupiterCaller: &service.RealJupiterCaller{},
			TokenRepository: mockTokenRepository,
		},
	}
	mockTokenRepository.On("JupiterTokenByAddress", mock.Anything).Return(true, "HIBER", nil)
	transactions := ReadBlockResponseFromFile("files/test_data_s_and_op3.txt").Result.Transactions

	actual, _ := serializer.Serialize(1, transactions[0])

	res := actual

	dke := assert.Equal(t, res.DcaKey, "EZSRPqvbWjH3YX3zgkjW9H2WugNMyyau2U652WusaNwv")
	ame := assert.Equal(t, res.InstructionData.InAmount, "25000000000")
	ampe := assert.Equal(t, res.InstructionData.InAmountPerCycle, "500000000")
	cfe := assert.Equal(t, res.InstructionData.CycleFrequency, "60")
	tkcae := assert.Equal(t, res.Token, "FHKiJEg2zmhv9DEeaXMSZa7R4P8BPFX5VTRrzrtJpump")
	tksymbe := assert.Equal(t, res.TokenSymbol, "HIBER")
	opee := assert.Equal(t, res.Operation.String(), "BUY")
	usae := assert.Equal(t, res.User, "HTpoBRnq1fv7AzDgWi6VULf2TCdGkseNdMqoCJJZcS9y")
	se := assert.Equal(t, res.Signature, "5z4TN1K3UCAAipRSW4i815R6ZSVj98TbDDpoxeWxER6K2dYyq171JDcJZZWZnR4SNm3xC1x3QW8oUwz4ksR1cN3b")

	r := ame && ampe && cfe && tkcae && tksymbe && opee && usae && se && dke
	if !r {
		t.Error("")
	}
}

type MockTokenRepository struct {
	mock.Mock
}

func (r *MockTokenRepository) UpdateExchangeTokenInfo(token model.Token) error {
	args := r.Called(token)
	return args.Error(0)
}

func (r *MockTokenRepository) SaveJupiterToken(address string, symbol string) error {
	args := r.Called(address, symbol)
	return args.Error(0)
}

func (r *MockTokenRepository) ExchangeTokenInfo(symbol string) (bool, model.Token, error) {
	args := r.Called(symbol)
	return args.Bool(0), args.Get(0).(model.Token), args.Error(2)
}

func (r *MockTokenRepository) JupiterTokenByAddress(address string) (bool, string, error) {
	args := r.Called(address)
	return args.Bool(0), args.String(1), args.Error(2)
}
