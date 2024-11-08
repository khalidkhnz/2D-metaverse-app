import { useState, useEffect } from "react";
import useDebounce from "./useDebounce"; // Assuming you have a useDebounce hook

function getWindowDimensions() {
  if (typeof window === "undefined") {
    return {
      width: 1280,
      height: 720,
    };
  }

  const { innerWidth: width, innerHeight: height } = window;
  return {
    width,
    height,
  };
}

export default function useWindowDimensions() {
  const [windowDimensions, setWindowDimensions] = useState(
    getWindowDimensions()
  );
  const debouncedWindowDimensions = useDebounce(windowDimensions, 200);

  useEffect(() => {
    function handleResize() {
      setWindowDimensions(getWindowDimensions());
    }

    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  return debouncedWindowDimensions;
}
