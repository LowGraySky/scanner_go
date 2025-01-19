package main

import (
	"time"
	"web3.kz/solscan/config"
	"web3.kz/solscan/service"
)

func main() {
	go schedule()

	select {}
}

func schedule() {
	config.Log.Println("Start analyse task")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		service.Process()
	}
}