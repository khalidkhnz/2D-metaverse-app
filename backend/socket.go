package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:4000"
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket connection to 'ws://localhost:4000/ws' failed: ", err)
		return
	}
	log.Printf("New WebSocket connection from %s\n", conn.RemoteAddr().String())

	// Emit "Hello" event to the new connection
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello"))
	if err != nil {
		log.Println(err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("recv: %s from %s\n", message, conn.RemoteAddr().String())
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}


// Below is the JavaScript code to connect to the WebSocket server

/*
const socket = new WebSocket('ws://localhost:4000/ws');

socket.onopen = function(event) {
    console.log('WebSocket connection established');
    socket.send('Hello Server!');
};

socket.onmessage = function(event) {
    console.log('Message from server: ', event.data);
};

socket.onerror = function(event) {
    console.error('WebSocket error: ', event);
};

socket.onclose = function(event) {
    console.log('WebSocket connection closed');
};
*/
