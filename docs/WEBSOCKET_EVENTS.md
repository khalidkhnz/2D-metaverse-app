# WebSocket Events Documentation

This document describes all WebSocket events used in the 2D Metaverse application for real-time communication between clients and server.

## Connection Flow

1. Client connects to WebSocket server at `ws://localhost:4000/ws`
2. Client must provide a valid token in the URL: `ws://localhost:4000/ws?token=<token>`
3. Upon successful connection, server assigns a unique remote address to the client

## Client Events (CLIENT:*)

These events are sent from the client to the server.

### CLIENT:REMOTE-ADDR
- **Purpose**: Request client's remote address
- **Payload**: None
- **Example**:
```json
{
  "type": "CLIENT:REMOTE-ADDR"
}
```

### CLIENT:POSITION
- **Purpose**: Update player's position in the game world
- **Payload**: 
  - `x`: number (x-coordinate)
  - `y`: number (y-coordinate)
- **Example**:
```json
{
  "type": "CLIENT:POSITION",
  "payload": {
    "x": 400,
    "y": 300
  }
}
```

### CLIENT:SEND-MESSAGE
- **Purpose**: Send a message in the game (chat functionality)
- **Payload**: Message content
- **Example**:
```json
{
  "type": "CLIENT:SEND-MESSAGE",
  "payload": {
    "message": "Hello, world!"
  }
}
```

## Server Events (SERVER:*)

These events are sent from the server to the client.

### SERVER:REMOTE-ADDR
- **Purpose**: Provide client with their unique remote address
- **Payload**: 
  - `remoteAddr`: string (unique identifier)
- **Example**:
```json
{
  "type": "SERVER:REMOTE-ADDR",
  "payload": {
    "remoteAddr": "user123"
  }
}
```

### SERVER:LOBBY-POSITIONS
- **Purpose**: Broadcast all players' positions to everyone in the game
- **Payload**: Object mapping player IDs to their positions
- **Example**:
```json
{
  "type": "SERVER:LOBBY-POSITIONS",
  "payload": {
    "user123": { "x": 400, "y": 300 },
    "user456": { "x": 500, "y": 200 }
  }
}
```

### SERVER:RECV-MESSAGE
- **Purpose**: Broadcast a received message to all clients
- **Payload**: Message content and metadata
- **Example**:
```json
{
  "type": "SERVER:RECV-MESSAGE",
  "payload": {
    "sender": "user123",
    "message": "Hello, world!",
    "timestamp": "2024-03-14T12:00:00Z"
  }
}
```

## Game Flow

1. **Connection**:
   - Client connects to WebSocket server with authentication token
   - Server sends `SERVER:REMOTE-ADDR` with client's unique ID
   - Client creates local player with received ID

2. **Position Updates**:
   - Client sends `CLIENT:POSITION` when player moves
   - Server broadcasts `SERVER:LOBBY-POSITIONS` to all clients
   - Clients update other players' positions based on received data

3. **Disconnection**:
   - When a client disconnects, server removes them from the game state
   - Next `SERVER:LOBBY-POSITIONS` broadcast won't include the disconnected player
   - Other clients remove the disconnected player's character

## Implementation Details

### Client-Side
- Uses Phaser.js for game rendering and physics
- Maintains local state of all players
- Green circle represents local player
- Red circles represent other players
- Arrow keys for movement
- Automatic reconnection attempts on disconnection

### Server-Side
- Maintains global game state
- Handles player connections/disconnections
- Broadcasts position updates to all connected clients
- Authenticates clients using tokens
- Provides unique identifiers for each client
