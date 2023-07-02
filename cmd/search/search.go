package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
	"github.com/gorilla/mux"
)

const (
	HOST        = "localhost"
	PORT        = "19200"
	TYPE        = "tcp"
	HTTP_ADRESS = "127.0.0.1"
	HTTP_PORT   = "8000"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte("This is a message"))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println("Received message:", string(received))

	conn.Close()

	// gorilla/mux => HTTP
	router := mux.NewRouter()
	router.HandleFunc("/search", createUrlHandler).Methods("POST")

	// Start The Server
	fmt.Printf("Serveur HTTP démarré sur http://%s:%s\n", HTTP_ADRESS, HTTP_PORT)
	log.Fatal(http.ListenAndServe(HTTP_ADRESS+":"+HTTP_PORT, router))
}

func createUrlHandler(resWriter http.ResponseWriter, req *http.Request) {
	// Lire le corps (body) de la requête
	var requestData struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(req.Body).Decode(&requestData)
	if err != nil {
		http.Error(resWriter, "Erreur lors de la lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	request := protocols.GenericRequest{
		Command: "createFile",
	}

	requestBuffer, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	conn.Write(requestBuffer)
	fmt.Println("Sent :", request)

	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		fmt.Println("Read data failed:", err.Error())
	}

	fmt.Println(string(received))

	// Spécifier le type de contenu JSON dans l'en-tête de la réponse
	resWriter.Header().Set("Content-Type", "application/json")

	// Écrire la réponse JSON dans le corps de la réponse
	resWriter.Write([]byte("ok"))

}

func getIndex() []string {
	// Retourner ici le tableau d'URLs de l'index
	index := []string{
		"http://example.com/page1",
		"http://example.com/page2",
		"http://example.com/page3",
	}
	return index
}
