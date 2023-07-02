package main

import (
	"encoding/json"
	"log"
	"net"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func createFileRequest(params string, results chan<- string) {
	con, err := net.Dial("tcp", "localhost:19200")
	createFileRequest := protocols.CreateOrUpdateFileRequest{}
	createFileRequest.Command = "createFile"
	createFileRequest_json, err := json.Marshal(createFileRequest)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}
	_, err = con.Write(createFileRequest_json)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	// buffer to get data
	received := make([]byte, 1024)
	nb, err := con.Read(received)
	if err != nil {
		log.Fatal("errrr ", err.Error())
	}
	listFileRes := protocols.CreateOrUpdateFileRequest{}
	json.Unmarshal([]byte(received[0:nb]), &listFileRes)
	con.Close()
	results <- string(received[0:nb])
	close(results)
}
