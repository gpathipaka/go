package main

import (
	"io"
	"log"
	"net/http"
)

func main1() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/ws", ws)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request....")
	io.WriteString(w, "Hellow Mr. Gangadhar")
}

func ws(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request....")
	//upgrader := websocket.Upgrader{}
	//upgrade := Websocket.Upgrade{}
	//conn, err := upgrader.Upgrade(w, r, nil)

}
