package service

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

const (
	chatId int64 = 8135757102
)

type RealTelegramCaller struct {
	Bot gotgbot.Bot
}

func (tc *RealTelegramCaller) SendMessage(message string) (*gotgbot.Message, error) {
	messageOptions := gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	}
	return tc.Bot.SendMessage(chatId, message, &messageOptions)
}

func (tc *RealTelegramCaller) SendReplyMessage(message string, messageId int64) error {
	replyMessageOptions := gotgbot.SendMessageOpts{
		ReplyParameters: &gotgbot.ReplyParameters{
			MessageId: messageId,
		},
	}
	_, err := tc.Bot.SendMessage(chatId, message, &replyMessageOptions)
	return err
}
