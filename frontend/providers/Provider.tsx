import { AppContextProvider } from "@/services/app-context";
import { SocketProvider } from "@/services/socket-provider";
import { Toaster } from "sonner";

const Provider = ({ children }: { children: React.ReactNode }) => {
  return (
    <AppContextProvider>
      <SocketProvider>{children}</SocketProvider>
      <Toaster richColors theme="dark" position="bottom-right" />
    </AppContextProvider>
  );
};

export default Provider;
