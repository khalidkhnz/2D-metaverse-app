"use client";

import { useEffect, useRef } from "react";
import { socketService } from "./socket";

export function SocketProvider({ children }: { children: React.ReactNode }) {
  const hasSetupSocket = useRef(false);

  useEffect(() => {
    // Setup socket connection on component mount
    if (!hasSetupSocket.current) {
      socketService.setupSocketConnection();
      hasSetupSocket.current = true;
    }
    return () => {
      // Ensure socket cleanup if necessary
      socketService.socket?.close();
    };
  }, []);

  return children;
}
