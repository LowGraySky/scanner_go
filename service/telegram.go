package service

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

const (
	chatId           int64 = 8135757102
	telegramBotToken       = "7460083410:AAF08myRfMh53DMJkefZvNhOQpddcJxPO5Q"
)

type RealTelegramCaller struct{}

func (tc *RealTelegramCaller) StartBot() (*gotgbot.Bot, error) {
	return gotgbot.NewBot(telegramBotToken, nil)
}

func (tc *RealTelegramCaller) SendMessage(bot gotgbot.Bot, message string) (*gotgbot.Message, error) {
	messageOptions := gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	}
	return bot.SendMessage(chatId, message, &messageOptions)
}

func (tc *RealTelegramCaller) SendReplyMessage(bot gotgbot.Bot, message string, messageId int64) error {
	replyMessageOptions := gotgbot.SendMessageOpts{
		ReplyParameters: &gotgbot.ReplyParameters{
			MessageId: messageId,
		},
	}
	_, err := bot.SendMessage(chatId, message, &replyMessageOptions)
	return err
}
