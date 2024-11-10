import { AppContextProvider } from "@/services/app-context";
import { SocketProvider } from "@/services/socket-provider";
import { Toaster } from "sonner";
import TanstackProvider from "./TanstackProvider";

const Provider = ({ children }: { children: React.ReactNode }) => {
  return (
    <TanstackProvider>
      <AppContextProvider>
        <SocketProvider>{children}</SocketProvider>
        <Toaster richColors theme="dark" position="bottom-right" />
      </AppContextProvider>
    </TanstackProvider>
  );
};

export default Provider;
