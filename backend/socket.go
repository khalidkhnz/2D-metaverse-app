package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
	authService "github.com/khalidkhnz/2D-metaverse-app/backend/services/auth"
	"github.com/khalidkhnz/2D-metaverse-app/backend/types"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:4000" || r.Header.Get("Origin") == "http://localhost:3000" || r.Header.Get("Origin") == "http://127.0.0.1:4000" || r.Header.Get("Origin") == "http://127.0.0.1:3000"
	},
}

var SOCKET_CONNECTIONS = make(map[string]*websocket.Conn)


func WSHandler(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("token")
	if token=="" {
		log.Println("WebSocket connection failed token not found")
		return
	}

	userProfile,err := authService.GetUserFromToken(token,true)
	if err!=nil {
		log.Println("WebSocket connection failed: ",err.Error())
		return
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket connection to 'ws://localhost:4000/ws' failed: ", err)
		return
	}

	log.Printf("New WebSocket connection from %s\n with Auth_Id: %s\n", conn.RemoteAddr(),userProfile.User.ID.Hex())

	// Store the new connection in SOCKET_CONNECTIONS
	SOCKET_CONNECTIONS[userProfile.User.ID.Hex()] = conn

	// Emit "REMOTE-ADDR" event to the new connection
	err = conn.WriteJSON(map[string]any{
		"type":"SERVER:REMOTE-ADDR",
		"payload": map[string]any{"remoteAddr": conn.RemoteAddr().String()},
	})
	if err != nil {
		log.Println(err)
		return
	}

	// GET MESSAGE AND RESPONSE
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("Unexpected WebSocket closure:", err)
			} else {
				log.Println("WebSocket closed:", err)
			}
			break // Exit the loop on close
		}
	
		// Try to unmarshal the message as JSON
		var jsonData map[string]any
		if err := json.Unmarshal(message, &jsonData); err == nil {
			// If message is JSON, handle it
			err = handleSocketEvents(jsonData, conn, userProfile)
			if err != nil {
				log.Printf("Error handling event from %s : %s \n", userProfile.User.FullName, err)
				// Optionally, notify the client of this specific error
				conn.WriteJSON(map[string]any{
					"type":    "ERROR",
					"payload": map[string]string{"message": err.Error()},
				})
			}
		} else {
			// If message is not JSON, send error response to client and continue
			log.Printf("Received non-JSON message from %s: %s\n", conn.RemoteAddr(), string(message))
			err = conn.WriteJSON(map[string]any{
				"type":    "ERROR",
				"payload": map[string]string{"message": "Only JSON messages are allowed"},
			})
			if err != nil {
				log.Println("Error sending error response:", err)
				break // Exit if unable to communicate with the client
			}
		}
	}
	
}

func handleSocketEvents(data map[string]any, conn *websocket.Conn, userProfile *types.FullProfile) (error) {
	log.Printf("Recv JSON from %s: \n %+v\n", userProfile.User.FullName, data)

	if data["type"] == "BROADCAST" {
		err := conn.WriteJSON(map[string]any{
			"type": "BROADCAST",
			"payload": data,
		})
		return err
	}

	if data["type"] == "DISCONNECT" {
		RemoveSocketConnection(userProfile.User.ID.Hex())
	}
	
	return fmt.Errorf("invalid event")
}

func RemoveSocketConnection(id string) {
    delete(SOCKET_CONNECTIONS, id)
}

func EmitEventTo(socketIds []string, eventType string, payload map[string]any) {
	for socketId, conn := range SOCKET_CONNECTIONS {
		if lib.Contains(socketIds, socketId) {
			err := conn.WriteJSON(map[string]any{
				"type":    eventType,
				"payload": payload,
			})
			log.Println(err.Error())
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
