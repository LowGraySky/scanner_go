package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"time"
	"web3.kz/solscan/config"
	"web3.kz/solscan/service"
)

func main() {
	go schedule()

	select {}
}

func schedule() {
	config.Log.Info("Start analyse task")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	telegramCaller := &service.RealTelegramCaller{}
	processor := &service.RealProcessor{
		Analyser: &service.RealAnalyser{},
		Serialiser: &service.RealSerializer{
			JupiterCaller: &service.RealJupiterCaller{},
		},
		SolanaCaller:   &service.RealSolanaCaller{},
		TelegramCaller: telegramCaller,
	}
	var bot gotgbot.Bot
	tgbot, err := telegramCaller.StartBot()
	if err != nil {
		config.Log.Errorf("Error when statring telegram bot, error: %q", err.Error())
		return
	}
	bot = *tgbot
	for range ticker.C {
		processor.Process(bot)
	}
}
