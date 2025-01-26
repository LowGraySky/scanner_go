package test

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"
	"web3.kz/solscan/model"
)

func ProcessTest(t *testing.T) {
	slot := model.GetSlotResponseBody{
		JsonRpc: "2.0",
		Result:  2,
		Id:      1,
		Error: model.Error{
			Code:    0,
			Message: "",
		},
	}

	block := readBlockResponseFromFile()

	if result != expect {
		t.Errorf("ERROR 1 <-----")
	}
}

func readBlockResponseFromFile() model.GetBlockResponseBody {
	file, err := os.Open("./test/files/procees_test.txt")
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
