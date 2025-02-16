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
	"sync"
	_ "sync"
	"time"
	"web3.kz/solscan/config"
	"web3.kz/solscan/service"
)

const (
	telegramBotToken = "7460083410:AAF08myRfMh53DMJkefZvNhOQpddcJxPO5Q"
	database         = "postgres://postgres:1122@localhost:5432/postgres?sslmode=disable"
)

func main() {
	db, err := applyMigrations()
	if err != nil {
		return
	}
	defer db.Close()

	processor, err1 := initProcessor(db)
	if err1 != nil {
		return
	}

	taskQueue := make(chan service.Task)

	workers := make([]service.Worker, 5)

	for i := 1; i <= 5; i++ {
		workers[i-1] = service.Worker{
			Id:       uint(i),
			JobQueue: taskQueue,
		}
	}

	var wg sync.WaitGroup

	for _, w := range workers {
		go func() {
			for task := range w.JobQueue {
				config.Log.Debugf("Worker-%d start execute task", w.Id)
				task()
				wg.Done()
			}
		}()
	}

	config.Log.Info("Start analyse task")
	ticker := time.NewTicker(500 * time.Millisecond)

	defer ticker.Stop()

	for range ticker.C {
		wg.Add(1)
		config.Log.Debugf("<- Append task to queue")
		taskQueue <- func() {
			processor.Process()
		}
	}

	wg.Wait()

	select {}
}

//func schedule(db *sql.DB, taskQueue chan service.Task, processor service.Processor) {
//	config.Log.Info("Start analyse task")
//	ticker := time.NewTicker(500 * time.Millisecond)
//
//	defer ticker.Stop()
//
//	for range ticker.C {
//		config.Log.Infof("<- Append task to queue")
//		taskQueue <- func() {
//			processor.Process()
//		}
//	}
//}

func applyMigrations() (*sql.DB, error) {
	conn, err := sql.Open("pgx/v5", database)
	if err != nil {
		config.Log.Errorf("Error when connect to database, error: %q", err.Error())
		return nil, err
	}
	driver, err1 := pgx.WithInstance(conn, &pgx.Config{})
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

func initProcessor(db *sql.DB) (*service.RealProcessor, error) {
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
		BitgetCaller:  &service.RealBitgetCaller{},
		TokenRepository: &service.RealTokenRepository{
			Db: *db,
		},
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
