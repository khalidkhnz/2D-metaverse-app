"use client";

import Avatar from "@/components/avatar";
import { Button } from "@/components/ui/button";
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "@/components/ui/hover-card";
import { socketService } from "@/services/socket";
import React, { useEffect } from "react";

const HomePage = () => {
  useEffect(() => {
    const handleSocketMessage = (event: MessageEvent<any>) => {
      console.log({ socketHome: JSON.parse(event.data) });
    };
    socketService.addMessageHandler(handleSocketMessage);
    return () => {
      socketService.removeMessageHandler(handleSocketMessage);
    };
  }, []);

  return (
    <main className="relative flex h-screen w-full items-center justify-center bg-[#02050d]">
      <TimeHoverCard />

      <h1 className="text-[4vw] font-light text-white 2xl:text-[64px]">
        Welcome to 2D Metaverse
      </h1>
      <div className="flex">
        <Avatar variant={"circle"} />
        <Button
          onClick={() => {
            if (socketService.socket) {
              socketService.socket.send(
                JSON.stringify({
                  type: "TEST",
                  message: "TEST MSG",
                }),
              );
            }
          }}
        >
          send
        </Button>
      </div>
    </main>
  );
};

function TimeHoverCard() {
  return (
    <HoverCard>
      <HoverCardTrigger className="absolute right-6 top-4 flex min-w-[200px] cursor-pointer items-center justify-center">
        <span className="text-[1.9vw] font-light uppercase text-white">
          {new Date().toLocaleTimeString([], {
            hour: "2-digit",
            minute: "2-digit",
            hour12: true,
          })}
        </span>
      </HoverCardTrigger>
      <HoverCardContent className="mx-2 flex items-center justify-center border-[#050b1c] bg-[#050b1c]">
        <span className="text-[16px] font-light text-white">
          {`Today is ${new Date().toLocaleDateString()} ${new Date().toLocaleDateString([], { weekday: "long" })}`}
        </span>
      </HoverCardContent>
    </HoverCard>
  );
}

export default HomePage;
