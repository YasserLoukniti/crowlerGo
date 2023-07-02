package main

import (
	"encoding/json"
	"log"
	"net"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func getFileRequest(results chan<- string, siteId int64) {
	con, err := net.Dial("tcp", "localhost:19200")
	getFileRequest := protocols.GetFileRequest{}
	getFileRequest.Command = "getFile"
	getFileRequest.SiteId = siteId
	getFileRequest_json, err := json.Marshal(getFileRequest)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}
	_, err = con.Write(getFileRequest_json)
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
	listFileRes := protocols.GetFileResponse{}
	json.Unmarshal([]byte(received[0:nb]), &listFileRes)
	con.Close()
	results <- string(received[0:nb])
}
