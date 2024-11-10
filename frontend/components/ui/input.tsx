import * as React from "react";

import { cn } from "@/lib/utils";

interface ExtendedInputProps {
  customVariant?: "default" | "primary";
}

const Input = React.forwardRef<
  HTMLInputElement,
  ExtendedInputProps & React.ComponentProps<"input">
>(({ className, type, customVariant = "default", ...props }, ref) => {
  if (customVariant === "default")
    return (
      <input
        type={type}
        className={cn(
          "flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
          className,
        )}
        ref={ref}
        {...props}
      />
    );

  if (customVariant === "primary")
    return (
      <input
        type={type}
        className={cn(
          "flex h-12 w-full rounded-md border-none bg-white/10 px-[20px] py-1 text-base text-white shadow-sm outline-none backdrop-blur-md transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-white/50 focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
          "focus-visible:bg-white/20",
          className,
        )}
        ref={ref}
        {...props}
      />
    );
});
Input.displayName = "Input";

export { Input };
