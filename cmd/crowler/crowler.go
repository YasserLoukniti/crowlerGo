package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

const (
	HOST = "localhost"
	PORT = "19200"
	TYPE = "tcp"
)

func main() {
	for {
		tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
		results := make(chan string)

		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP(TYPE, nil, tcpServer)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
		}

		getSiteRequest := protocols.GetSiteRequest{}
		getSiteRequest.Command = "getSite"
		getSiteRequest_json, err := json.Marshal(getSiteRequest)
		if err != nil {
			println("marshal failed:", err.Error())
			os.Exit(1)
		}
		_, err = conn.Write(getSiteRequest_json)
		if err != nil {
			println("Write data failed:", err.Error())
			os.Exit(1)
		}

		// buffer to get data
		received := make([]byte, 1024)
		nb, err := conn.Read(received)
		if err != nil {
			log.Fatal(err)
		}
		cleanReceived := received[0:nb]
		println("Received message:", string(cleanReceived))
		listRes := protocols.GetSiteResponse{}
		json.Unmarshal([]byte(cleanReceived), &listRes)
		if listRes.Command == "getSite" && listRes.Status == 200 {
			println("Entered")
			// Visiter le site et sauvgarde fichiers
			go visitSite(listRes.Sites[0], results)
			// get files
			go getFileRequest(results, listRes.Sites[0].Id)
			// Update Site
			go updateSiteRequest("any", listRes.Sites[0], results)
		} else {
			time.Sleep(5 * time.Second)
		}

		for res := range results {
			fmt.Println("Result:", res)
		}

		conn.Close()
	}

}
