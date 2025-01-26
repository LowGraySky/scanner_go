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
	config.Log.Info("Start analyse task")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	processor := createProcessor()
	for range ticker.C {
		processor.Process()
	}
}

func createProcessor() service.Processor {
	return &service.RealProcessor {
		Analyser: &service.RealAnalyser{},
		Serialiser: &service.RealSerializer{},
		SolanaCaller: &service.RealSolanaCaller{},
	}
}