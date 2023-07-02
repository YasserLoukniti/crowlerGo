package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func getFile(buffer []byte, database protocols.Database, responseChan chan<- []byte) {

	listReq := protocols.GetFileRequest{}
	json.Unmarshal([]byte(buffer), &listReq)
	res := protocols.GetFileResponse{}
	res.Command = "getFile"
	res.Files = database.Files
	res.Status = 200
	fmt.Printf("res => : %v\n", res)
	listres_json, err := json.Marshal(res)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}

	responseChan <- listres_json

}
