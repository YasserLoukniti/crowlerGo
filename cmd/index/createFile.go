package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func createFile(buffer []byte, database protocols.Database, responseChan chan<- []byte) {

	createFileReq := protocols.CreateOrUpdateFileRequest{}
	json.Unmarshal([]byte(buffer), &createFileReq)
	res := protocols.CreateOrUpdateFileResponse{}

	last := len(database.Files)
	createFileReq.File.Id = int64(last)
	database.Files = append(database.Files, createFileReq.File)
	res.Command = "createFile"
	res.File = createFileReq.File
	res.Status = 200

	fmt.Printf("res => : %v\n", res)
	createFileRes_json, err := json.Marshal(res)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}

	responseChan <- createFileRes_json

}
