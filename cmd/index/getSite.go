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
	resSite := protocols.Site{Id: 0}
	for _, site := range database.Sites {

		if site.Lastseen.IsZero() {
			println("Its Zeroo")
			resSite = site
			break
		}
	}
	res.Command = "getSite"
	if resSite.Id == 0 {
		res.Status = 404
	} else {
		res.Sites = []protocols.Site{resSite}
		res.Status = 200
	}

	fmt.Printf("res => : %v\n", res)
	listres_json, err := json.Marshal(res)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}

	responseChan <- listres_json

}
