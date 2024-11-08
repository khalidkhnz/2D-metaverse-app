import { SocketProvider } from "@/services/socket";
import { Toaster } from "sonner";

const Provider = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <SocketProvider>{children}</SocketProvider>
      <Toaster richColors theme="dark" position="bottom-right" />
    </>
  );
};

export default Provider;
