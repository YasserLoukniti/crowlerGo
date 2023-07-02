package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func createSite(buffer []byte, database protocols.Database, responseChan chan<- []byte) {

	createSiteReq := protocols.CreateOrUpdateSiteRequest{}
	json.Unmarshal([]byte(buffer), &createSiteReq)
	res := protocols.CreateOrUpdateSiteResponse{}

	last := len(database.Sites)
	createSiteReq.Site.Id = int64(last)
	database.Sites = append(database.Sites, createSiteReq.Site)
	res.Command = "createSite"
	res.Site = createSiteReq.Site
	res.Status = 200

	fmt.Printf("res => : %v\n", res)
	createSiteRes_json, err := json.Marshal(res)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}

	responseChan <- createSiteRes_json

}
