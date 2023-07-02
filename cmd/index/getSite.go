package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func getSites(buffer []byte, database protocols.Database, responseChan chan<- []byte) {

	listReq := protocols.GetSiteRequest{}
	json.Unmarshal([]byte(buffer), &listReq)
	res := protocols.GetSiteResponse{}
	res.Command = "getSite"
	res.Sites = database.Sites
	res.Status = 200
	fmt.Printf("res => : %v\n", res)
	listres_json, err := json.Marshal(res)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}
	responseChan <- listres_json

}
