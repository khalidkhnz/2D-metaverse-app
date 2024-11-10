import { cn } from "@/lib/utils";
import { Avatar as AvatarShad, AvatarFallback, AvatarImage } from "./ui/avatar";
import { CSSProperties } from "react";

interface props {
  variant?: "circle" | "default";
  fallback?: any;
  className?: string;
  imgClassName?: string;
  fallbackClassName?: string;
  imgUrl?: string;
  style?: CSSProperties;
}

export default function Avatar({
  variant = "default",
  fallback,
  className,
  fallbackClassName,
  imgClassName,
  style,
  ...props
}: props) {
  if (variant === "default")
    return (
      <AvatarShad style={{ ...style }} className={cn(className)} {...props}>
        <AvatarImage
          className={imgClassName}
          src={"https://github.com/shadcn.png"}
        />
        <AvatarFallback className={fallbackClassName}>
          {fallback}
        </AvatarFallback>
      </AvatarShad>
    );
}
