"use client";
import { useEffect, useState } from "react";

type KeyType = string;
type InitialValueType<T> = T | (() => T);
type ReturnType<T> = [T, (newValue: T, expirationTime?: number) => void];

interface StoredValue<T> {
  value: T;
  expiry: number | null; // Expiry is null if no expiration time is set
}

function useRealTimeLocalStorage<T>(
  key: KeyType,
  initialValue: InitialValueType<T>,
): ReturnType<T> {
  const localStorageAvailable = typeof localStorage !== "undefined";
  const [value, setValue] = useState<T>(() => {
    if (localStorageAvailable) {
      const storedData = localStorage.getItem(key);
      if (storedData) {
        try {
          const parsedData: StoredValue<T> = JSON.parse(storedData);
          if (parsedData.expiry === null || parsedData.expiry > Date.now()) {
            return parsedData.value;
          } else {
            localStorage.removeItem(key); // Remove expired item
          }
        } catch (error) {
          console.error("Error parsing localStorage value:", error);
        }
      }
    }
    return typeof initialValue === "function"
      ? (initialValue as () => T)()
      : initialValue;
  });

  useEffect(() => {
    const handleStorageChange = (event: StorageEvent) => {
      if (event.key === key && event.newValue !== null) {
        try {
          const parsedData: StoredValue<T> = JSON.parse(event.newValue);
          if (parsedData.expiry === null || parsedData.expiry > Date.now()) {
            setValue(parsedData.value);
          } else {
            localStorage.removeItem(key); // Remove expired item
          }
        } catch (error) {
          console.error("Error parsing localStorage value:", error);
        }
      }
    };

    if (localStorageAvailable) {
      window.addEventListener("storage", handleStorageChange);
    }

    return () => {
      if (localStorageAvailable) {
        window.removeEventListener("storage", handleStorageChange);
      }
    };
  }, [key, localStorageAvailable]);

  const updateValue = (newValue: T, expirationTime?: number) => {
    setValue(newValue);
    if (localStorageAvailable) {
      const expiry = expirationTime ? Date.now() + expirationTime : null;
      const data: StoredValue<T> = { value: newValue, expiry };
      localStorage.setItem(key, JSON.stringify(data));
    }
  };

  return [value, updateValue];
}

export default useRealTimeLocalStorage;
