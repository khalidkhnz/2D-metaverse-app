"use client";

import dynamic from "next/dynamic";
import React from "react";

const Game = dynamic(() => import("./Game"), {
  ssr: false,
  loading: () => (
    <div className="flex h-screen w-full items-center justify-center bg-[#242424]">
      <h1 className="text-xl font-bold text-white">LOADING...</h1>
    </div>
  ),
});

const GamePage: React.FC = () => {
  return (
    <main className="w-full">
      <Game />
    </main>
  );
};

export default GamePage;
