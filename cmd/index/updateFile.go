package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func updateFile(buffer []byte, database protocols.Database, responseChan chan<- []byte) {

	updateFilereq := protocols.CreateOrUpdateFileRequest{}
	json.Unmarshal([]byte(buffer), &updateFilereq)
	res := protocols.CreateOrUpdateFileResponse{}
	resFile := protocols.File{Id: 0}
	for i, file := range database.Files {
		if updateFilereq.File.Id == file.Id {
			database.Files[i] = updateFilereq.File
			resFile = updateFilereq.File
		}
	}
	res.Command = "updateFile"
	res.File = resFile
	res.Status = 200

	fmt.Printf("res => : %v\n", res)
	updateFileRes_json, err := json.Marshal(res)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}

	responseChan <- updateFileRes_json

}
