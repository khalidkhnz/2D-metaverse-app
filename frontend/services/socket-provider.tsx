"use client";

import { useEffect } from "react";
import { socketService } from "./socket";

export function SocketProvider({ children }: { children: React.ReactNode }) {
  useEffect(() => {
    // Setup socket connection on component mount
    socketService.setupSocketConnection();

    return () => {
      // Ensure socket cleanup if necessary
      socketService.socket?.close();
    };
  }, []);

  return children;
}
