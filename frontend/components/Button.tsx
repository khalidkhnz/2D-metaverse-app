import React from "react";
import { Button as Btn, ButtonProps } from "./ui/button";
import { cn } from "@/lib/utils";

interface ExtendedButtonProps extends ButtonProps {
  customVariants?: "default" | "primary" | "secondary";
}

const Button = React.forwardRef<HTMLButtonElement, ExtendedButtonProps>(
  (
    {
      className,
      variant,
      size,
      asChild = false,
      customVariants = "default",
      children,
      ...props
    },
    ref,
  ) => {
    if (customVariants == "default")
      return (
        <Btn
          className={className}
          variant={variant}
          size={size}
          asChild={asChild}
          ref={ref}
          {...props}
        >
          {children}
        </Btn>
      );

    if (customVariants == "primary")
      return (
        <Btn
          className={cn(
            "bg-white/20 backdrop-blur-md hover:bg-white/30",
            "disabled:bg-white/5",
            className,
          )}
          variant={variant}
          size={size}
          asChild={asChild}
          ref={ref}
          {...props}
        >
          {children}
        </Btn>
      );

    if (customVariants == "secondary")
      return (
        <Btn
          className={className}
          variant={variant}
          size={size}
          asChild={asChild}
          ref={ref}
          {...props}
        >
          {children}
        </Btn>
      );
  },
);

Button.displayName = "Button";

export default Button;
