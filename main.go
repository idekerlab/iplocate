package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// Response holds the senders IP address
type Response struct {
	SenderAddress string `json:"sender_address"`
	SenderPort    string `json:"sender_port"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uri := strings.Split(r.RemoteAddr, ":")
		log.Println("Sender address:", uri[0], "port:", uri[1])
		enc := json.NewEncoder(w)
		enc.Encode(Response{
			SenderAddress: uri[0],
			SenderPort:    uri[1],
		})
	})
	log.Println("Starting iplocate service")
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
