import type { Metadata } from "next";
import "@/styles/globals.css";
import Provider from "@/providers/Provider";
import { Open_Sans } from "next/font/google";

export const metadata: Metadata = {
  title: "2D Metaverse",
  description: "",
};

const OpenSans = Open_Sans({
  subsets: ["latin"],
  display: "swap",
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${OpenSans.className} antialiased`}>
        <Provider>{children}</Provider>
      </body>
    </html>
  );
}
