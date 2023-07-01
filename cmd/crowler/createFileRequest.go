package main

import (
	"encoding/json"
	"net"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func createFileRequest(conn net.Conn) {
	createFileRequest := protocols.CreateOrUpdateFileRequest{}
	createFileRequest.Command = "createFile"
	createFileRequest_json, err := json.Marshal(createFileRequest)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write(createFileRequest_json)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
}
