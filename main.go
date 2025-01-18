package main

import (
	"fmt"
	"time"
	"web3.kz/solscan/service"
)

func main() {
	go schedule()

	select {}
}

func schedule() {
	fmt.Print("Start analyse task\n")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		service.Analyse()
	}
}