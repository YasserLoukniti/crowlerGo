package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

const (
	HOST = "localhost"
	PORT = "19200"
	TYPE = "tcp"
)

func main() {
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

	//GetSiteRequest
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
	// TODO: Visiter le site
	println("Received message:", string(cleanReceived))
	listRes := protocols.GetSiteResponse{}
	json.Unmarshal([]byte(cleanReceived), &listRes)
	if listRes.Command == "getSite" && listRes.Status == 200 {
		println("Entered")
		go getFileRequest("any", results)
	}

	for res := range results {
		println("Waiting For Results")
		fmt.Println("Result:", res)
	}

	conn.Close()
}
