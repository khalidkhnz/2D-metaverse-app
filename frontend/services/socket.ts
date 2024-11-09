"use client";

import { Toast } from "@/lib/toast";
import { wait } from "@/lib/util";

type ClientEvents =
  | "CLIENT:REMOTE-ADDR"
  | "CLIENT:POSITION"
  | "CLIENT:SEND-MESSAGE";
type ServerEvents =
  | "SERVER:REMOTE-ADDR"
  | "SERVER:LOBBY-POSITIONS"
  | "SERVER:RECV-MESSAGE";

class WS {
  public socket!: WebSocket;
  private retryAttempts = 0;
  private maxRetries = 1000;
  private retryDelay = 5000;
  private externalMessageHandlers: ((event: MessageEvent<any>) => void)[] = [];

  public setupSocketConnection() {
    this.connectSocket();
  }

  private connectSocket() {
    this.socket = new WebSocket("ws://localhost:4000/ws");
    this.socketConnectionEvents({
      handleSocketEvents: this.handleSocketEvents,
      onConnectSend: { type: "CONNECT", payload: { userId: "123123" } },
    });
  }

  private retryConnection() {
    if (this.retryAttempts < this.maxRetries) {
      this.retryAttempts++;
      console.log(`Attempting to reconnect... (Attempt ${this.retryAttempts})`);
      Toast.info(`Attempting to reconnect... (Attempt ${this.retryAttempts})`);
      wait(this.retryDelay).then(() => this.connectSocket());
    } else {
      console.error("Max retry attempts reached. Could not reconnect.");
      Toast.error("Unable to reconnect to the server.");
    }
  }

  private socketConnectionEvents = async ({
    handleSocketEvents,
    onConnectSend,
  }: {
    handleSocketEvents: (event: MessageEvent<any>) => void;
    onConnectSend: any;
  }) => {
    this.socket.onopen = async () => {
      console.log("WebSocket connection established");
      this.retryAttempts = 0;
      await wait(1000);
      this.socket.send(JSON.stringify(onConnectSend));
    };

    this.socket.onmessage = (event) => {
      handleSocketEvents(event);
      this.externalMessageHandlers.forEach((handler) => handler(event));
    };

    this.socket.onerror = () => {
      console.log("WebSocket error occurred");
      this.retryConnection();
    };

    this.socket.onclose = () => {
      console.log("WebSocket connection closed");
      this.retryConnection();
    };
  };

  private handleSocketEvents(e: MessageEvent<any>) {
    const data: { type: ServerEvents; payload: any } = JSON.parse(e.data);
    console.log("JSON from server:", data);

    switch (data?.type) {
      case "SERVER:REMOTE-ADDR": {
        const payload: { remoteAddr: string } = data.payload;
        console.log("REMOTE-ADDR", payload);
        break;
      }
      case "SERVER:RECV-MESSAGE": {
        const payload = data.payload;
        console.log("SERVER:RECV-MESSAGE", payload);
        break;
      }
      case "SERVER:LOBBY-POSITIONS": {
        const payload = data.payload;
        console.log("SERVER:LOBBY-POSITIONS", payload);
        break;
      }
    }
  }

  public sendSocketEvent({
    type,
    payload,
  }: {
    type: ClientEvents;
    payload: any;
  }) {
    this.socket?.send(JSON.stringify({ type, payload }));
  }

  public addMessageHandler(handler: (event: MessageEvent<any>) => void) {
    this.externalMessageHandlers.push(handler);
  }

  public removeMessageHandler(handler: (event: MessageEvent<any>) => void) {
    this.externalMessageHandlers = this.externalMessageHandlers.filter(
      (h) => h !== handler,
    );
  }
}

export const socketService = new WS();
