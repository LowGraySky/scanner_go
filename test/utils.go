package test

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"web3.kz/solscan/model"
)

func ReadBlockResponseFromFile(path string) model.GetBlockResponseBody {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	reader := bufio.NewReader(file)
	body, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	var block model.GetBlockResponseBody
	err1 := json.Unmarshal(body, &block)
	if err1 != nil {
		log.Fatal(err)
	}
	return block
}