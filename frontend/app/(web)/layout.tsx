import { AuroraBackground } from "@/components/aurora-background";
import React from "react";

type Props = {
  children: React.ReactNode;
};

const Layout = ({ children }: Props) => {
  return <AuroraBackground>{children}</AuroraBackground>;
};

export default Layout;
