package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func updateSite(buffer []byte, database protocols.Database, responseChan chan<- []byte) {

	updateSitereq := protocols.CreateOrUpdateSiteRequest{}
	json.Unmarshal([]byte(buffer), &updateSitereq)
	res := protocols.CreateOrUpdateSiteResponse{}
	resSite := protocols.Site{Id: 0}
	for i, site := range database.Sites {
		if updateSitereq.Site.Id == site.Id {
			database.Sites[i] = updateSitereq.Site
			resSite = updateSitereq.Site
		}
	}
	res.Command = "updateSite"
	res.Site = resSite
	res.Status = 200

	fmt.Printf("res => : %v\n", res)
	updateSiteRes_json, err := json.Marshal(res)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}

	responseChan <- updateSiteRes_json

}
