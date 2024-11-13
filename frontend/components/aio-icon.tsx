import { cn } from "@/lib/utils";
import Avatar from "./avatar";
import { HoverCard, HoverCardContent, HoverCardTrigger } from "./ui/hover-card";
import { Plus } from "lucide-react";
import Image from "next/image";
import { Toast } from "@/lib/toast";
import gsap from "gsap";

interface IStateImageType {
  result: string | ArrayBuffer | null | undefined;
  file: File | null;
}

export default function AioIcon({
  handleImageChange,
  name,
  animationId,
  className,
  avatarClassName,
  onClick,
  onMouseEnter,
  variant,
  email,
  enableAnimation = false,
  disableToolTip = false,
  textClassName,
}: {
  enableAnimation?: boolean;
  disableToolTip?: boolean;
  handleImageChange?: {
    image: IStateImageType;
    setImage: React.Dispatch<React.SetStateAction<IStateImageType>>;
  };
  animationId: string;
  name?: string;
  email?: string;
  className?: string;
  textClassName?: string;
  avatarClassName?: string;
  active?: boolean;
  variant?: "add-account";
  onClick?: () => void;
  onMouseEnter?: () => void;
}) {
  function onMouseEnterAnimation() {
    if (!enableAnimation) return;
    gsap.to(`.aio-card-${animationId}`, {
      scale: "1.2",
    });
    gsap.to(`.aio-avatar-${animationId}`, {
      boxShadow: "2px 0 3px 1px #ffffff",
    });
  }

  function onMouseLeaveAnimation() {
    if (!enableAnimation) return;
    gsap.to(`.aio-card-${animationId}`, {
      scale: "1",
    });

    gsap.to(`.aio-avatar-${animationId}`, {
      boxShadow: "0px 0 10px 1px transparent",
    });
  }

  if (variant === "add-account") {
    return (
      <div
        className={cn(
          "flex flex-col items-center justify-center gap-2",
          `aio-card-${animationId}`,
          className,
        )}
        onClick={onClick}
        onMouseEnter={() => {
          onMouseEnter && onMouseEnter();
          onMouseEnterAnimation();
        }}
        onMouseLeave={onMouseLeaveAnimation}
      >
        <div
          className={cn(
            "relative h-[10vw] max-h-[160px] min-h-[110px] w-[10vw] min-w-[110px] max-w-[160px] cursor-pointer",
            "mt-[10px] border-[4px] border-transparent",
            "flex items-center justify-center overflow-hidden rounded-full bg-white/20 backdrop-blur-md",
            `aio-avatar-${animationId}`,
            avatarClassName,
          )}
        >
          <Plus
            className={cn(
              "h-[3vw] max-h-[50px] min-h-[25px] w-[3vw] min-w-[25px] max-w-[50px]",
              "aspect-square text-white",
            )}
          />
          {handleImageChange?.image.result && (
            <Image
              className="object-cover"
              src={handleImageChange?.image.result as string}
              alt="profile"
              fill
            />
          )}
          <input
            type="file"
            onChange={(e) => {
              if (handleImageChange && !handleImageChange?.setImage) return;
              const file = e?.target?.files?.[0];
              if (file && file.type.startsWith("image/")) {
                const reader = new FileReader();
                reader.onload = (ev) =>
                  handleImageChange?.setImage({
                    result: ev?.target?.result,
                    file: file,
                  });
                reader.readAsDataURL(file);
              } else Toast.warning("Invalid image please select another");
            }}
            className="home_add_image_picker pointer-events-none absolute left-0 top-0 h-full w-full opacity-0"
          />
        </div>
        <h2
          className={cn(
            "cursor-pointer text-[18px] font-light text-white md:text-[1.8vw] lg:text-[25px]",
            textClassName,
          )}
        >
          {name}
        </h2>
      </div>
    );
  }

  if (disableToolTip) {
    return (
      <div
        className={cn(
          "flex flex-col items-center justify-center gap-2",
          `aio-card-${animationId}`,
          className,
        )}
        onClick={onClick}
        onMouseEnter={() => {
          onMouseEnter && onMouseEnter();
          onMouseEnterAnimation();
        }}
        onMouseLeave={onMouseLeaveAnimation}
      >
        <Avatar
          className={cn(
            "h-[10vw] max-h-[160px] min-h-[110px] w-[10vw] min-w-[110px] max-w-[160px] cursor-pointer",
            "mt-[10px] border-[4px] border-transparent",
            `aio-avatar-${animationId}`,
            avatarClassName,
          )}
          variant={"default"}
        />
        <h2
          className={cn(
            "cursor-pointer text-[18px] font-light text-white md:text-[1.8vw] lg:text-[25px]",
            textClassName,
          )}
        >
          {name}
        </h2>
      </div>
    );
  }

  return (
    <HoverCard>
      <HoverCardTrigger>
        <div
          className={cn(
            "flex flex-col items-center justify-center gap-2",
            `aio-card-${animationId}`,
            className,
          )}
          onClick={onClick}
          onMouseEnter={() => {
            onMouseEnter && onMouseEnter();
            onMouseEnterAnimation();
          }}
          onMouseLeave={onMouseLeaveAnimation}
        >
          <Avatar
            className={cn(
              "h-[10vw] max-h-[160px] min-h-[110px] w-[10vw] min-w-[110px] max-w-[160px] cursor-pointer",
              "mt-[10px] border-[4px] border-transparent",
              `aio-avatar-${animationId}`,
              avatarClassName,
            )}
            variant={"default"}
          />
          <h2
            className={cn(
              "cursor-pointer text-[18px] font-light text-white md:text-[1.8vw] lg:text-[25px]",
              textClassName,
            )}
          >
            {name}
          </h2>
        </div>
      </HoverCardTrigger>
      <HoverCardContent className="mt-4 flex items-center justify-center border-none bg-white/20 p-2 backdrop-blur-md">
        <span className="overflow-hidden overflow-ellipsis text-nowrap text-sm font-light text-white">
          {email}
        </span>
      </HoverCardContent>
    </HoverCard>
  );
}
