package main

import (
	"encoding/json"
	"net"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func updateFileRequest(conn net.Conn) {
	updateFileRequest := protocols.CreateOrUpdateFileRequest{}
	updateFileRequest.Command = "updateFile"
	updateFileRequest_json, err := json.Marshal(updateFileRequest)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write(updateFileRequest_json)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
}
