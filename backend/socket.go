package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:4000" || r.Header.Get("Origin") == "http://localhost:3000" || r.Header.Get("Origin") == "http://127.0.0.1:4000" || r.Header.Get("Origin") == "http://127.0.0.1:3000"
	},
}

var SOCKET_CONNECTIONS = make(map[*websocket.Conn]string)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket connection to 'ws://localhost:4000/ws' failed: ", err)
		return
	}
	log.Printf("New WebSocket connection from %s\n", conn.RemoteAddr())

	// Store the new connection in SOCKET_CONNECTIONS
	SOCKET_CONNECTIONS[conn] = conn.RemoteAddr().String()

	// Emit "Hello" event to the new connection
	err = conn.WriteJSON(map[string]any{
		"success":"true",
		"data": map[string]any{"socketId": conn.RemoteAddr().String()},
	})
	if err != nil {
		log.Println(err)
		return
	}

	// GET MESSAGE AND RESPONSE
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Try to unmarshal the message as JSON
		var jsonData map[string]any
		if err := json.Unmarshal(message, &jsonData); err == nil {
			// If message is JSON
			log.Printf("Received JSON message from %s: %+v\n", conn.RemoteAddr(), jsonData)
			err = conn.WriteJSON(map[string]any{
				"success": true,
				"data":    jsonData,
			})
		} else {
			// If message is not JSON, treat it as a string
			log.Printf("Received string message from %s: %s\n", conn.RemoteAddr(), string(message))
			err = conn.WriteJSON(map[string]any{
				"success": true,
				"data":    string(message),
			})
		}


		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}


// Below is the JavaScript code to connect to the WebSocket server
///////////////////////
// const socket = new WebSocket('ws://localhost:4000/ws');

// socket.onopen = function(event) {
//     console.log('WebSocket connection established');

//     // Create a JSON object to send
//     const jsonData = {
//         message: "Hello Server!",
//         timestamp: new Date().toISOString()
//     };

//     // Send the JSON object as a string
//     socket.send(JSON.stringify(jsonData));
// };

// socket.onmessage = function(event) {
//     try {
//         // Attempt to parse the server response as JSON
//         const data = JSON.parse(event.data);
//         console.log('JSON Message from server:', data);
//     } catch (e) {
//         // If parsing fails, log the message as a string
//         console.log('Message from server:', event.data);
//     }
// };

// socket.onerror = function(event) {
//     console.error('WebSocket error:', event);
// };

// socket.onclose = function(event) {
//     console.log('WebSocket connection closed');
// };
///////////////////////

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
