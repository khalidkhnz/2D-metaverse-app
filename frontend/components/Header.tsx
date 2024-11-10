"use client";

import React, { useState } from "react";
import Avatar from "./avatar";
import { Assets } from "@/lib/Assets";
import Image from "next/image";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuPortal,
  DropdownMenuSub,
  DropdownMenuSubContent,
  DropdownMenuSubTrigger,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { cn } from "@/lib/utils";
import { useAppContext } from "@/services/app-context";
import { usePathname } from "next/navigation";
import dynamic from "next/dynamic";

const TimeHoverCard = dynamic(() => import("@/components/time-hover-card"), {
  ssr: false,
  loading: () => <div className="w-[80px]" />,
});

const Header = () => {
  const pathname = usePathname();

  return (
    <div className="absolute left-0 top-0 z-50 flex h-fit w-full items-center justify-end gap-8 bg-transparent py-2">
      <h1 className="ml-[30px] mr-auto cursor-pointer text-[28px] font-light text-white">
        {pathname
          .split("/")
          .map((i) => i.trim())
          .join(" ")
          .toUpperCase()}
      </h1>
      <Image
        className="aspect-square h-9 w-9 cursor-pointer"
        src={Assets.SETTINGS}
        alt="settings"
      />
      <AvatarDropDownMenu />
      <TimeHoverCard />
    </div>
  );
};

function AvatarDropDownMenu() {
  const [activeStatus, setActiveStatus] = useState({
    name: "ONLINE",
    icon: <div className="h-3 w-3 rounded-full bg-green-600" />,
  });
  const { current_user } = useAppContext();

  const statuses = {
    ONLINE: {
      name: "ONLINE",
      icon: <div className="h-3 w-3 rounded-full bg-green-600" />,
    },
    AWAY: {
      name: "AWAY",
      icon: <div className="h-3 w-3 rounded-full bg-yellow-600" />,
    },
    BUSY: {
      name: "BUSY",
      icon: <div className="h-3 w-3 rounded-full bg-orange-800" />,
    },
    "IN A MEETING": {
      name: "IN A MEETING",
      icon: <div className="h-3 w-3 rounded-full bg-red-600" />,
    },
    "OUT OF OFFICE": {
      name: "OUT OF OFFICE",
      icon: <div className="h-3 w-3 rounded-full bg-pink-600" />,
    },
  };

  const options = [
    ...(current_user
      ? [{ name: current_user.fullName, className: "font-normal" }]
      : []),
    {
      name: activeStatus.name,
      subMenu: Object.keys(statuses).map(
        (key) => statuses[key as keyof typeof statuses],
      ),
      icon: activeStatus.icon,
    },
    { name: "Profile" },
    { name: "Space" },
    { name: "Settings" },
    {
      name: "Logout",
      className:
        "focus:bg-red-500 font-normal focus:text-white backdrop-blur-md",
    },
  ];

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="rounded-full border-none outline-none">
        <Avatar />
      </DropdownMenuTrigger>
      <DropdownMenuContent className="mt-2 min-w-[230px] border-none bg-gradient-to-bl from-[#161a1e] to-[#262b2c] backdrop-blur-md">
        {options?.map((option, idx) => {
          if (option.subMenu?.length) {
            return (
              <DropdownMenuSub key={`subMenu-${idx}`}>
                <DropdownMenuSubTrigger
                  className={cn(
                    "text-md flex items-center justify-start rounded-[2px] p-2 px-4 font-light text-white focus-within:text-black focus:text-black data-[state=open]:text-black",
                    {
                      "border-t-[1px] border-[#1b1e22]": idx !== 0,
                    },
                    option.className,
                  )}
                >
                  {option?.icon && option.icon}
                  <span className="capitalize">
                    {option.name?.toLowerCase()}
                  </span>
                </DropdownMenuSubTrigger>
                <DropdownMenuPortal>
                  <DropdownMenuSubContent className="border-none bg-gradient-to-bl from-[#161a1e] to-[#262b2c] backdrop-blur-md">
                    {option.subMenu.map((subOption, index) => {
                      return (
                        <DropdownMenuItem
                          className={cn(
                            "text-md flex w-[160px] items-center justify-between rounded-[2px] p-2 px-4 font-light text-white",
                            {
                              "border-t-[1px] border-[#1b1e22]": idx !== 0,
                            },
                          )}
                          key={`subOption-${index}`}
                          onClick={() => setActiveStatus(subOption)}
                        >
                          <span className="text-sm capitalize">
                            {subOption.name.toLowerCase()}
                          </span>
                          {subOption?.icon && subOption.icon}
                        </DropdownMenuItem>
                      );
                    })}
                  </DropdownMenuSubContent>
                </DropdownMenuPortal>
              </DropdownMenuSub>
            );
          }

          return (
            <DropdownMenuItem
              className={cn(
                "text-md flex items-center justify-between rounded-[2px] p-2 px-4 font-light text-white",
                {
                  "border-t-[1px] border-[#1b1e22]": idx !== 0,
                },
                option.className,
              )}
              key={`option-${idx}`}
            >
              <span>{option.name}</span>
              {option?.icon && option.icon}
            </DropdownMenuItem>
          );
        })}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

export default Header;
