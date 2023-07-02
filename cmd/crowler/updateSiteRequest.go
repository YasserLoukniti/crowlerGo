package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func updateSiteRequest(params string, site protocols.Site, results chan<- string) {
	con, err := net.Dial("tcp", "localhost:19200")
	updateSiteRequest := protocols.CreateOrUpdateSiteRequest{}
	updateSiteRequest.Command = "updateSite"
	site.Lastseen = time.Now()
	updateSiteRequest.Site = site
	updateSiteRequest_json, err := json.Marshal(updateSiteRequest)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}
	_, err = con.Write(updateSiteRequest_json)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	// buffer to get data
	received := make([]byte, 1024)
	nb, err := con.Read(received)
	if err != nil {
		log.Fatal("errrr ", err.Error())
	}
	updateSiteRes := protocols.CreateOrUpdateSiteResponse{}
	json.Unmarshal([]byte(received[0:nb]), &updateSiteRes)
	con.Close()
	results <- string(received[0:nb])
}
