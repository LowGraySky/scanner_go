package service

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"web3.kz/solscan/config"
)

const (
	chatId int64 = 8135757102
    telegramBotToken = "7460083410:AAF08myRfMh53DMJkefZvNhOQpddcJxPO5Q"
)

type RealTelegramCaller struct {}

var bot *gotgbot.Bot

func init() {
	bot, _ = gotgbot.NewBot(telegramBotToken, nil)
}

func (tc *RealTelegramCaller) SendMessage(message string) {
	_, err := bot.SendMessage(chatId, message, nil)
	if err != nil {
		config.Log.Error("Message hasn't delivered!")
	}
}