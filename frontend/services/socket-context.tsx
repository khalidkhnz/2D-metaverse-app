"use client";

import { Toast } from "@/lib/toast";
import { wait } from "@/lib/util";
import { createContext, useContext, useEffect, useState } from "react";

type ClientEvents =
  | "CLIENT:REMOTE-ADDR"
  | "CLIENT:POSITION"
  | "CLIENT:SEND-MESSAGE";

type ServerEvents =
  | "SERVER:REMOTE-ADDR"
  | "SERVER:LOBBY-POSITIONS"
  | "SERVER:RECV-MESSAGE";

interface ISocketContext {
  socket: WebSocket | null;
  sendSocketEvent({
    type,
    payload,
  }: {
    type: ClientEvents;
    payload: any;
  }): void;
}

const SocketContext = createContext<ISocketContext | null>(null);

export function SocketProvider({ children }: { children: React.ReactNode }) {
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    let socket: WebSocket | null = null;
    let retryTimeout: NodeJS.Timeout;

    const connect = () => {
      socket = new WebSocket("ws://localhost:4000/ws");

      socket.onopen = async function (event) {
        console.log("WebSocket connection established");
        await wait(1000);
        setWs(socket);
        onConnectEmit();
      };

      socket.onmessage = function (event) {
        try {
          handleSocketEvents(event);
        } catch (e) {
          Toast.default(`Message From Server: ${event.data}`);
        }
      };

      socket.onerror = function (event) {
        console.log("WebSocket error:", event);
        setWs(null);
        retryConnection();
      };

      socket.onclose = function (event) {
        console.log("WebSocket connection closed");
        setWs(null);
        retryConnection();
      };
    };

    const retryConnection = () => {
      if (socket) {
        socket.close();
      }
      retryTimeout = setTimeout(() => {
        console.log("Attempting to reconnect...");
        connect();
      }, 5000); // Retry every 5 seconds
    };

    connect();

    return () => {
      if (socket) {
        socket.close();
      }
      clearTimeout(retryTimeout);
    };
  }, []);

  function handleSocketEvents(e: MessageEvent<any>) {
    const data: { type: ServerEvents; payload: any } = JSON.parse(e.data);
    console.log("JSON Message from server:", data);

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

  function onConnectEmit() {
    ws?.send(
      JSON.stringify({ type: "CONNECT", payload: { userId: "123123" } })
    );
  }

  function sendSocketEvent({
    type,
    payload,
  }: {
    type: ClientEvents;
    payload: any;
  }) {
    ws?.send(JSON.stringify({ type, payload }));
  }

  return (
    <SocketContext.Provider value={{ socket: ws, sendSocketEvent }}>
      {children}
    </SocketContext.Provider>
  );
}

export function useSocket() {
  const ctx = useContext(SocketContext);
  if (!ctx) {
    throw new Error("Socket can't be used outside provider boundaries");
  }
  return ctx;
}
