"use client";

import { useEffect, useState } from "react";
import { HoverCard, HoverCardContent, HoverCardTrigger } from "./ui/hover-card";
import { cn } from "@/lib/utils";

export default function TimeHoverCard({ className }: { className?: string }) {
  const [currentTime, setCurrentTime] = useState(
    new Date().toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
      hour12: true,
    }),
  );

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentTime(
        new Date().toLocaleTimeString([], {
          hour: "2-digit",
          minute: "2-digit",
          // second: "2-digit",
          hour12: true,
        }),
      );
    }, 60000);
    return () => clearInterval(interval);
  }, []);

  return (
    <HoverCard>
      <HoverCardTrigger
        className={cn(
          "flex min-w-[200px] cursor-pointer items-center justify-center",
          className,
        )}
      >
        <span className="text-[1.9vw] font-light uppercase text-white">
          {currentTime}
        </span>
      </HoverCardTrigger>
      <HoverCardContent
        className={cn(
          "mx-2 flex items-center justify-center border-[#050b1c] bg-[#050b1c]",
          "border-none bg-white/20 backdrop-blur-md",
        )}
      >
        <span className="text-[16px] font-light text-white">
          {`Today is ${new Date().toLocaleDateString()} ${new Date().toLocaleDateString([], { weekday: "long" })}`}
        </span>
      </HoverCardContent>
    </HoverCard>
  );
}