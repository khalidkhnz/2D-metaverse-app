"use client";

import { useEffect, useRef } from "react";
import { socketService } from "./socket";
import { useAppContext } from "./app-context";

export function SocketProvider({ children }: { children: React.ReactNode }) {
  const { current_user, token } = useAppContext();
  const isConnected = useRef(false);

  useEffect(() => {
    if (current_user && token && !isConnected.current) {
      socketService.setupSocketConnection();
      isConnected.current = true;
    }

    return () => {
      if (isConnected.current) {
        socketService.socket?.close();
      }
    };
  }, []);

  return children;
}
