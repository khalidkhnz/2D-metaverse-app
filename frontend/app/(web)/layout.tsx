import { AuroraBackground } from "@/components/aurora-background";
import Header from "@/components/Header";
import React from "react";

type Props = {
  children: React.ReactNode;
};

const Layout = ({ children }: Props) => {
  return (
    <AuroraBackground>
      {children}
      <Header />
    </AuroraBackground>
  );
};

export default Layout;
