"use client";

import Avatar from "@/components/avatar";
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "@/components/ui/hover-card";
import { socketService } from "@/services/socket";
import React, { useEffect, useState } from "react";

import { cn } from "@/lib/utils";
import gsap from "gsap";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";
import Button from "@/components/Button";
import { Plus } from "lucide-react";

const HomePage = () => {
  const [active, setActive] = useState(0);

  useEffect(() => {
    const handleSocketMessage = (event: MessageEvent<any>) => {
      console.log({ socketHome: JSON.parse(event.data) });
    };
    socketService.addMessageHandler(handleSocketMessage);
    return () => {
      socketService.removeMessageHandler(handleSocketMessage);
    };
  }, []);

  function handleContinue(mode?: "add-account" | "login") {
    if (active === 0 || mode == "add-account") {
      console.log("ADD ACCOUNT");
    } else {
      console.log("LOGGIN IN PREVIOUS ACCOUNT");
    }
  }

  return (
    <main
      className={cn(
        "relative flex h-screen w-full flex-col items-center justify-center bg-transparent",
      )}
    >
      <TimeHoverCard />

      <h1 className="text-[30px] font-light text-white md:text-[4vw] 2xl:text-[64px]">
        Welcome to 2D Metaverse
      </h1>
      <span className="mt-2 text-[14px] font-light text-white lg:text-[1.3vw] 2xl:text-[20px]">
        Welcome Back | स्वागत हे | مرحباً | Bienvenido
      </span>
      <div className="mt-[100px] w-[100%] justify-center py-4">
        <HomeCarousel
          onClick={handleContinue}
          data={[]}
          active={active}
          setActive={setActive}
        />
      </div>
      <div className="">
        <Button
          onClick={() => handleContinue()}
          className="w-[160px]"
          customVariants="primary"
        >
          {active === 0 ? "Add Account" : "Continue"}
        </Button>
      </div>
    </main>
  );
};

function HomeCarousel({
  active = 0,
  setActive,
  data,
  onClick,
}: {
  onClick?: (val: "add-account" | "login") => void;
  data?: any[];
  active: number;
  setActive: React.Dispatch<React.SetStateAction<number>>;
}) {
  const LENGHT = data?.length || 0;

  useEffect(() => {
    gsap.to(`.active-user-card`, {
      scale: "1.2",
    });
    gsap.to(`.inactive-user-card`, {
      scale: "1",
    });
    gsap.to(`.active-user-avatar`, {
      boxShadow: "2px 0 3px 1px #ffffff",
    });
    gsap.to(`.inactive-user-avatar`, {
      boxShadow: "0px 0 10px 1px transparent",
    });
  }, [active]);

  return (
    <Carousel className="mx-auto w-[90%] md:w-[80%]">
      <CarouselContent>
        <CarouselItem
          className={cn({
            "basis-1/1": LENGHT === 1,
            "basis-1/2": LENGHT === 2,
            "basis-1/3": LENGHT >= 3 && LENGHT < 5,
            "basis-1/3 lg:basis-1/5": LENGHT >= 5,
          })}
        >
          <div
            onClick={() => onClick && onClick("add-account")}
            className={cn(
              "flex items-center justify-center p-1 py-5 pb-[80px]",
            )}
          >
            <User
              variant={"add-account"}
              onClick={() => setActive(0)}
              onMouseEnter={() => setActive(0)}
              className={
                0 === active ? "active-user-card" : "inactive-user-card"
              }
              avatarClassName={
                0 === active ? "active-user-avatar" : "inactive-user-avatar"
              }
            />
          </div>
        </CarouselItem>
        {Array.from({ length: LENGHT }).map((_, index) => (
          <CarouselItem
            key={index + 1}
            className={cn({
              "basis-1/3": LENGHT < 5,
              "basis-1/3 lg:basis-1/5": LENGHT >= 5,
            })}
          >
            <div
              className={cn(
                "flex items-center justify-center p-1 py-5 pb-[80px]",
              )}
            >
              <User
                onClick={() => setActive(index + 1)}
                onMouseEnter={() => setActive(index + 1)}
                name={`khalid.khnz ${index + 1}`}
                className={
                  index + 1 === active
                    ? "active-user-card"
                    : "inactive-user-card"
                }
                avatarClassName={
                  index + 1 === active
                    ? "active-user-avatar"
                    : "inactive-user-avatar"
                }
              />
            </div>
          </CarouselItem>
        ))}
      </CarouselContent>
      <CarouselPrevious />
      <CarouselNext />
    </Carousel>
  );
}

function User({
  name,
  className,
  avatarClassName,
  onClick,
  onMouseEnter,
  variant,
}: {
  name?: string;
  className?: string;
  avatarClassName?: string;
  active?: boolean;
  variant?: "add-account";
  onClick?: () => void;
  onMouseEnter?: () => void;
}) {
  if (variant === "add-account") {
    return (
      <div
        className={cn(
          "flex flex-col items-center justify-center gap-2",
          className,
        )}
        onClick={onClick}
        onMouseEnter={onMouseEnter}
      >
        <div
          className={cn(
            "h-[10vw] max-h-[160px] min-h-[110px] w-[10vw] min-w-[110px] max-w-[160px]",
            "mt-[10px] border-[4px] border-transparent",
            "flex items-center justify-center rounded-full bg-white/20 backdrop-blur-md",
            avatarClassName,
          )}
        >
          <Plus
            className={cn(
              "h-[3vw] max-h-[50px] min-h-[25px] w-[3vw] min-w-[25px] max-w-[50px]",
              "aspect-square text-white",
            )}
          />
        </div>
        <h2 className="text-[12px] font-light text-white md:text-[1.5vw] lg:text-[20px]">
          Add Account
        </h2>
      </div>
    );
  }

  return (
    <HoverCard>
      <HoverCardTrigger>
        <div
          className={cn(
            "flex flex-col items-center justify-center gap-2",
            className,
          )}
          onClick={onClick}
          onMouseEnter={onMouseEnter}
        >
          <Avatar
            className={cn(
              "h-[10vw] max-h-[160px] min-h-[110px] w-[10vw] min-w-[110px] max-w-[160px]",
              "mt-[10px] border-[4px] border-transparent",
              avatarClassName,
            )}
            variant={"default"}
          />
          <h2 className="text-[18px] font-light text-white md:text-[1.8vw] lg:text-[25px]">
            {name}
          </h2>
        </div>
      </HoverCardTrigger>
      <HoverCardContent className="mt-4 flex items-center justify-center border-none bg-white/20 p-2 backdrop-blur-md">
        <span className="overflow-hidden overflow-ellipsis text-nowrap text-sm font-light text-white">
          eternalkhalidkhnz@gmail.com
        </span>
      </HoverCardContent>
    </HoverCard>
  );
}

function TimeHoverCard() {
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
      <HoverCardTrigger className="absolute right-6 top-4 flex min-w-[200px] cursor-pointer items-center justify-center">
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

export default HomePage;
