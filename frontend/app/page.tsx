import Avatar from "@/components/avatar";
import React from "react";

const HomePage = () => {
  return (
    <main className="relative flex h-screen w-full items-center justify-center bg-[#02050d]">
      <span className="absolute right-6 top-4 text-[1.9vw] font-light text-white">
        10:53 PM
      </span>
      <h1 className="text-[4vw] font-light text-white 2xl:text-[64px]">
        Welcome to 2D Metaverse
      </h1>
      <div className="flex">
        <Avatar variant={"circle"} />
      </div>
    </main>
  );
};

export default HomePage;
