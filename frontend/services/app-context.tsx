"use client";

import { ILogin, IRegister } from "@/types/auth";
import { createContext, useContext, useEffect, useRef, useState } from "react";
import { getMethod, postMethod } from "./ApiInterceptor";
import { ENDPOINTS } from "@/lib/Endpoints";
import { Toast } from "@/lib/toast";
import { capitalize } from "@/lib/utils";
import useRealTimeLocalStorage from "@/hooks/useRealTimeLocalStorage";
import { IUser, IUserApiResponse } from "@/types/user";

interface IAppContext {
  current_user: IUser | null | undefined;
  token: string;
  setToken: (newValue: string, expirationTime?: number) => void;
  handleRegister: (values: IRegister) => void;
  handleLogin: (values: ILogin) => void;
}

const AppContext = createContext<IAppContext | null>(null);

export function AppContextProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [token, setToken] = useRealTimeLocalStorage<string>("token", "");
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
      console.log(response);
      return response as IUserApiResponse;
    } else {
      return null;
    }
  }

  async function handleLogin(values: ILogin) {
    const response = await postMethod<{
      success?: boolean;
      data?: any;
      message?: string;
    }>(ENDPOINTS.AUTH.LOGIN_ACCOUNT, values).catch((err) => console.log(err));

    if (response && response.success) {
      setToken(response?.data?.token, 7 * 24 * 60 * 60 * 1000);
      await handleGetAndSetCurrentUser();
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
      await handleLogin({
        email: values.email,
        password: values.password,
      });
    }
    if (response && !response.success) {
      Toast.error(capitalize(response?.message?.toLowerCase()));
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
