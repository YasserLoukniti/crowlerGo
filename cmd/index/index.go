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
	nb, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("json server => : %s\n", buffer[0:nb])
	cleanBuffer := buffer[0:nb]
	req := protocols.GenericRequest{}
	json.Unmarshal([]byte(cleanBuffer), &req)
	if req.Command == "getSite" {
		getSites(cleanBuffer, req, conn, database)
	}

	// write data to response
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[:]), time)
	conn.Write([]byte(responseStr))

	// close conn
	conn.Close()
}
