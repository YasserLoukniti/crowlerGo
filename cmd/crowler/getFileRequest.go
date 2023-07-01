package main

import (
	"encoding/json"
	"net"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func getFileRequest(conn net.Conn) {
	getFileRequest := protocols.GetFileRequest{}
	getFileRequest.Command = "getFile"
	getFileRequest_json, err := json.Marshal(getFileRequest)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write(getFileRequest_json)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
}
