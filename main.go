package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/go-redis/redis/v8"
	"time"
	"web3.kz/solscan/config"
	"web3.kz/solscan/service"
)

const telegramBotToken = "7460083410:AAF08myRfMh53DMJkefZvNhOQpddcJxPO5Q"

func main() {
	go schedule()

	select {}
}

func schedule() {
	config.Log.Info("Start analyse task")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	processor, err := initProcessor()
	if err != nil {
		return
	}
	for range ticker.C {
		processor.Process()
	}
}

func initProcessor() (*service.RealProcessor, error) {
	bot, err := gotgbot.NewBot(telegramBotToken, nil)
	if err != nil {
		config.Log.Errorf("Error when statring telegram bot, error: %q", err.Error())
		return &service.RealProcessor{}, err
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &service.RealProcessor{
		Analyser: &service.RealAnalyser{},
		Serialiser: &service.RealSerializer{
			TokenFetcher: &service.RealTokenFetcher{
				JupiterCaller: &service.RealJupiterCaller{},
			},
		},
		RedisCaller: &service.RealRedisCaller{
			RedisClient: *redisClient,
		},
		SolanaCaller: &service.RealSolanaCaller{},
		TelegramCaller: &service.RealTelegramCaller{
			Bot: *bot,
		},
	}, nil
}
