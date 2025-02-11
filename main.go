package main

import (
	"database/sql"
	"errors"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/go-redis/redis/v8"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
	"web3.kz/solscan/config"
	"web3.kz/solscan/service"
)

const (
	telegramBotToken = "7460083410:AAF08myRfMh53DMJkefZvNhOQpddcJxPO5Q"
	database         = "postgres://postgres:1122@localhost:5432/postgres?sslmode=disable"
)

func main() {
	go schedule()

	select {}
}

func schedule() {
	config.Log.Info("Start analyse task")
	ticker := time.NewTicker(2 * time.Second)
	db, err := applyMigrations()
	if err != nil {
		return
	}
	defer db.Close()
	defer ticker.Stop()
	processor, err1 := initProcessor()
	if err1 != nil {
		return
	}
	for range ticker.C {
		processor.Process()
	}
}

func applyMigrations() (*sql.DB, error) {
	conn, err := sql.Open("pgx/v5", database)
	if err != nil {
		config.Log.Errorf("Error when connect to database, error: %q", err.Error())
		return nil, err
	}
	driver, err1 := pgx.WithInstance(conn, &pgx.Config{

	})
	if err1 != nil {
		config.Log.Errorf("Error when creating driver, error: %q", err1.Error())
		return nil, err1
	}
	m, err2 := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err2 != nil {
		config.Log.Errorf("Error create database instance, error: %q", err2.Error())
		return nil, err2
	}
	err3 := m.Up()
	if err3 != nil && !errors.Is(err3, migrate.ErrNoChange) {
		config.Log.Errorf("Error apply database migrations, error: %q", err3.Error())
		return nil, err3
	}
	config.Log.Info("Migrations applied successfully")
	return conn, nil
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
	tokenFetcher := &service.RealTokenFetcher{
		JupiterCaller: &service.RealJupiterCaller{},
		MexcCaller:    &service.RealMexcCaller{},
		GateCaller:    &service.RealGateCaller{},
	}
	return &service.RealProcessor{
		Analyser: &service.RealAnalyser{},
		Serialiser: &service.RealSerializer{
			TokenFetcher: tokenFetcher,
		},
		TokenFetcher: tokenFetcher,
		RedisCaller: &service.RealRedisCaller{
			RedisClient: *redisClient,
		},
		SolanaCaller: &service.RealSolanaCaller{},
		TelegramCaller: &service.RealTelegramCaller{
			Bot: *bot,
		},
	}, nil
}
