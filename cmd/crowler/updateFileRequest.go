package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func updateFileRequest(params string, file protocols.File, results chan<- string) {
	con, err := net.Dial("tcp", "localhost:19200")
	updateFileRequest := protocols.CreateOrUpdateFileRequest{}
	updateFileRequest.Command = "updateSite"
	file.Lastseen = time.Now()
	updateFileRequest.File = file
	updateFileRequest_json, err := json.Marshal(updateFileRequest)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}
	_, err = con.Write(updateFileRequest_json)
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
	updateFileRes := protocols.CreateOrUpdateFileResponse{}
	json.Unmarshal([]byte(received[0:nb]), &updateFileRes)
	con.Close()
	results <- string(received[0:nb])
}
