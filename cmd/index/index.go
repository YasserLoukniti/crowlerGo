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

var database protocols.Database

func main() {
	database.Sites = append(database.Sites, protocols.Site{
		Id:     1,
		HostIp: "212.27.63.171",
		Domain: "http://maxicool5.free.fr/Potager/",
	})
	database.Files = append(database.Files, protocols.File{
		Id:     1,
		Name:   "Calendrier semis 1.jpg",
		Url:    "http://maxicool5.free.fr/Potager/Calendrier%20semis%201.jpg",
		SiteId: 1,
	}, protocols.File{
		Id:     2,
		Name:   "Calendrier semis 2.jpg",
		Url:    "http://maxicool5.free.fr/Potager/Calendrier%20semis%202.jpg",
		SiteId: 1,
	})
	listen, err := net.Listen(TYPE, HOST+":"+PORT)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	cleanBuffer := cleaningBuffer(conn, buffer)
	req := protocols.GenericRequest{}
	json.Unmarshal([]byte(cleanBuffer), &req)

	responseChan := make(chan []byte)
	switch req.Command {
	case "getSite":
		go getSites(cleanBuffer, database, responseChan)
	case "getFile":
		go getFile(cleanBuffer, database, responseChan)
	case "updateSite":
		go updateSite(cleanBuffer, database, responseChan)
	default:
		fmt.Println("Command not found")
	}

	response := <-responseChan

	// write response to connection
	conn.Write([]byte(response))

	// close connection
	conn.Close()
}

func cleaningBuffer(conn net.Conn, buffer []byte) []byte {
	nb, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	cleanBuffer := buffer[0:nb]
	return cleanBuffer
}
