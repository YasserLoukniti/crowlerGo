package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

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

	// gorilla/mux => HTTP
	router := mux.NewRouter()
	router.HandleFunc("/search", createSiteHandler).Methods("POST")

	// Start The Server
	fmt.Printf("Serveur HTTP démarré sur http://%s:%s\n", HTTP_ADRESS, HTTP_PORT)
	log.Fatal(http.ListenAndServe(HTTP_ADRESS+":"+HTTP_PORT, router))
}

func createSiteHandler(resWriter http.ResponseWriter, req *http.Request) {
	// Lire le corps (body) de la requête
	var requestData struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(req.Body).Decode(&requestData)
	if err != nil {
		http.Error(resWriter, "Erreur lors de la lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	request := protocols.CreateOrUpdateSiteRequest{}
	request.Command = "createSite"
	request.Site = protocols.Site{Domain: requestData.URL}
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
