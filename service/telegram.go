package service

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"web3.kz/solscan/config"
)

const (
	chatId           int64 = 8135757102
	telegramBotToken       = "7460083410:AAF08myRfMh53DMJkefZvNhOQpddcJxPO5Q"
)

type RealTelegramCaller struct{}

func (tc *RealTelegramCaller) StartBot() (*gotgbot.Bot, error) {
	return gotgbot.NewBot(telegramBotToken, nil)
}

func (tc *RealTelegramCaller) SendMessage(bot gotgbot.Bot, message string) {
	messageOptions := gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	}
	_, err := bot.SendMessage(chatId, message, &messageOptions)
	if err != nil {
		config.Log.Errorf("Message %s hasn't delivered!, err: %q", message, err.Error())
	} else {
		config.Log.Infof("Succuess send telegram message: %s", message)
	}
}
