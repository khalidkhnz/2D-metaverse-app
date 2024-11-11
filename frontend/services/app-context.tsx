"use client";

import { ILogin, IRegister } from "@/types/auth";
import { createContext, useContext, useEffect, useRef, useState } from "react";
import { getMethod, postMethod } from "./ApiInterceptor";
import { ENDPOINTS } from "@/lib/Endpoints";
import { Toast } from "@/lib/toast";
import { capitalize, wait } from "@/lib/utils";
import useRealTimeLocalStorage from "@/hooks/useRealTimeLocalStorage";
import { IUser, IUserApiResponse } from "@/types/user";
import { socketService } from "./socket";
import { useRouter } from "next/navigation";
import gsap from "gsap";

interface IAppContext {
  current_user: IUser | null | undefined;
  token: string | null;
  setToken: (newValue: string, expirationTime?: number) => void;
  handleRegister: (values: IRegister) => void;
  handleLogin: (values: ILogin) => void;
  handleLogout: () => void;
}

const AppContext = createContext<IAppContext | null>(null);

export function AppContextProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const [token, setToken] = useRealTimeLocalStorage<string | null>("token", "");
  const [currentUser, setCurrentUser] = useRealTimeLocalStorage<IUser | null>(
    "user",
    null,
  );

  const runOnce = useRef(false);
  useEffect(() => {
    if (runOnce.current && currentUser) return;
    handleGetAndSetCurrentUser();
    runOnce.current = true;
  }, [token]);

  async function handleGetAndSetCurrentUser(): Promise<IUserApiResponse | null> {
    if (!token) return null;
    const response = await getMethod<IUserApiResponse>(
      ENDPOINTS.AUTH.CURRENT_USER,
    ).catch((err) => console.log(err));

    if (response && response?.success) {
      setCurrentUser(response?.data?.user);
      return response as IUserApiResponse;
    } else {
      return null;
    }
  }

  async function handleLogin(values: ILogin) {
    const response = await postMethod<{
      success?: boolean;
      data?: { token: string; user: IUser };
      message?: string;
    }>(ENDPOINTS.AUTH.LOGIN_ACCOUNT, values).catch((err) => console.log(err));

    if (response && response.success && response.data) {
      setToken(response?.data?.token, 7 * 24 * 60 * 60 * 1000);
      setCurrentUser(response?.data?.user);
      socketService.setupSocketConnection();
      gsap.to(".home-main", {
        pointerEvents: "none",
        opacity: 0,
      });
      Toast.success("Logged in successfully");
    }
    if (response && !response.success) {
      Toast.error(capitalize(response?.message?.toLowerCase()));
    }
  }

  async function handleRegister(values: IRegister) {
    const response = await postMethod<{
      success?: boolean;
      data?: any;
      message?: string;
    }>(ENDPOINTS.AUTH.CREATE_ACCOUNT(), values).catch((err) =>
      console.log(err),
    );

    if (response && response.success) {
      Toast.success("Registered in successfully");
      await handleLogin({
        email: values.email,
        password: values.password,
      });
    }
    if (response && !response.success) {
      Toast.error(capitalize(response?.message?.toLowerCase()));
    }
  }

  async function handleLogout() {
    try {
      // Clear the token and current user
      setToken(null);
      setCurrentUser(null);
      await wait(1000);
      // Close the socket connection if it exists
      socketService.socket?.close();

      Toast.success("Logged out successfully");
      router.push("/");
    } catch (error) {
      console.log(error);
      Toast.error("An error occurred during logout");
    }
  }

  return (
    <AppContext.Provider
      value={{
        token,
        current_user: currentUser,
        setToken,
        handleRegister,
        handleLogin,
        handleLogout,
      }}
    >
      {children}
    </AppContext.Provider>
  );
}

export function useAppContext() {
  const ctx = useContext(AppContext);
  if (!ctx) {
    throw new Error("App context can't be used outside provider boundaries");
  }
  return ctx;
}
