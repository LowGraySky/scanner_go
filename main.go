package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/go-redis/redis/v8"
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
	var bot gotgbot.Bot
	tgbot, err := telegramCaller.StartBot()
	if err != nil {
		config.Log.Errorf("Error when statring telegram bot, error: %q", err.Error())
		return
	}
	bot = *tgbot
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	processor := &service.RealProcessor{
		Bot:      bot,
		Analyser: &service.RealAnalyser{},
		Serialiser: &service.RealSerializer{
			TokenFetcher: &service.RealTokenFetcher{
				JupiterCaller: &service.RealJupiterCaller{},
			},
		},
		RedisCaller:    &service.RealRedisCaller{
			RedisClient: *redisClient,
		},
		SolanaCaller:   &service.RealSolanaCaller{},
		TelegramCaller: telegramCaller,
	}

	for range ticker.C {
		processor.Process()
	}
}
