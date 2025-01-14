package main

import "time"
import "web3.kz/solscan/service"

func main() {
	Scdedule()
}

func Scdedule() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case  <-ticker.C:
			go service.Analyse()
		}
	}
}