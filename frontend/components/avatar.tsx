import { cn } from "@/lib/utils";
import { Avatar as AvatarShad, AvatarFallback, AvatarImage } from "./ui/avatar";

interface props {
  variant?: "circle" | "default";
  fallback?: any;
  className?: string;
  imgClassName?: string;
  fallbackClassName?: string;
}

export default function Avatar({
  variant = "default",
  fallback,
  className,
  fallbackClassName,
  imgClassName,
}: props) {
  return (
    <AvatarShad className={cn(className)}>
      <AvatarImage className="" src="https://github.com/shadcn.png" />
      <AvatarFallback className="">CN</AvatarFallback>
    </AvatarShad>
  );
}
