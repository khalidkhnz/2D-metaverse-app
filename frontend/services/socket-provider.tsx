"use client";

import { useEffect } from "react";
import { socketService } from "./socket";
import { useAppContext } from "./app-context";

export function SocketProvider({ children }: { children: React.ReactNode }) {
  const { current_user, token } = useAppContext();

  useEffect(() => {
    // Setup socket connection on component mount
    if (current_user && token) {
      socketService.setupSocketConnection();
    }
    return () => {
      // Ensure socket cleanup if necessary
      socketService.socket?.close();
    };
  }, [current_user, token]);

  return children;
}
