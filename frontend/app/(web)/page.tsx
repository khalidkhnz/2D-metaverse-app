"use client";

import Avatar from "@/components/avatar";
import { useFormik } from "formik";
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "@/components/ui/hover-card";
import { socketService } from "@/services/socket";
import React, { useEffect, useState } from "react";

import { cn } from "@/lib/utils";
import gsap from "gsap";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";
import Button from "@/components/Button";
import { Plus } from "lucide-react";
import { Input } from "@/components/ui/input";
import { useAppContext } from "@/services/app-context";
import { Y } from "@/lib/Y";
import { FormsBody } from "@/lib/Forms";
import { Toast } from "@/lib/toast";
import Image from "next/image";

interface IStateImageType {
  result: string | ArrayBuffer | null | undefined;
  file: File | null;
}

const HomePage = () => {
  const [active, setActive] = useState(0);
  const [showAddAccountForm, setShowAddAccountForm] = useState(false);
  const [image, setImage] = useState<IStateImageType>({
    file: null,
    result: null,
  });

  useEffect(() => {
    const handleSocketMessage = (event: MessageEvent<any>) => {
      console.log({ socketHome: JSON.parse(event.data) });
    };
    socketService.addMessageHandler(handleSocketMessage);

    const showScreenTimer = setTimeout(() => {
      gsap.to(".home-main", {
        display: "flex",
        opacity: 1,
        duration: 2,
        ease: "power2.out",
      });
    }, 500);

    return () => {
      clearTimeout(showScreenTimer);
      socketService.removeMessageHandler(handleSocketMessage);
    };
  }, []);

  function handleContinue(idx?: number) {
    handleAnimation();
    if (active === 0 || idx == 0) {
      console.log("ADD ACCOUNT");
      setShowAddAccountForm(true);
      gsap.to(".carousel__btn", {
        opacity: 0,
      });
      gsap.to(".auth__form", {
        translateY: "-200px",
        display: "flex",
        opacity: 1,
        pointerEvents: "auto",
      });
    } else {
      console.log("LOGGIN IN PREVIOUS ACCOUNT");
    }
  }

  function handleAnimation() {
    gsap.to("#welcome-text", {
      translateY: "-200px",
      opacity: 0,
      duration: 0.3,
    });
    gsap.to("#welcome-carousel", {
      translateY: "-200px",
      duration: 0.3,
    });
    gsap.to("#welcome-continue", {
      opacity: 0,
      duration: 0.3,
      onComplete: () => {
        gsap.set("#welcome-continue", {
          pointerEvents: "none",
        });
      },
    });
  }

  return (
    <main
      className={cn(
        "home-main hidden overflow-hidden opacity-0",
        "relative flex h-screen w-full flex-col items-center justify-center bg-transparent",
      )}
    >
      <h1
        id="welcome-text"
        className="text-[30px] font-light text-white md:text-[4vw] 2xl:text-[64px]"
      >
        Welcome to 2D Metaverse
      </h1>
      <span
        id="welcome-text"
        className="mt-2 text-[14px] font-light text-white lg:text-[1.3vw] 2xl:text-[20px]"
      >
        Welcome Back | स्वागत हे | مرحباً | Bienvenido
      </span>
      <div
        id="welcome-carousel"
        className="mt-[100px] w-[100%] justify-center py-4"
      >
        <HomeCarousel
          handleImageChange={{
            image,
            setImage,
          }}
          showAddAccountForm={showAddAccountForm}
          onClick={handleContinue}
          data={[]}
          active={active}
          setActive={setActive}
        />
      </div>
      <AuthForm className="auth__form pointer-events-none hidden opacity-0" />
      <div id="welcome-continue" className="">
        <Button
          onClick={() => handleContinue()}
          className="w-[160px]"
          customVariants="primary"
        >
          {active === 0 ? "Add Account" : "Continue"}
        </Button>
      </div>
    </main>
  );
};

function AuthForm({ className }: { className?: string }) {
  const [isLoginForm, setIsLoginForm] = useState(false);
  const { handleRegister, handleLogin } = useAppContext();

  const registerFk = useFormik({
    initialValues: FormsBody.register(),
    onSubmit: handleRegister,
    validationSchema: Y.register,
  });

  const loginFk = useFormik({
    initialValues: FormsBody.login(),
    onSubmit: handleLogin,
    validationSchema: Y.login,
  });

  useEffect(() => {
    if (isLoginForm) {
      gsap.to("#welcome-carousel", {
        opacity: 0,
        pointerEvents: "none",
      });
    } else {
      gsap.to("#welcome-carousel", {
        opacity: 1,
        pointerEvents: "auto",
      });
    }
  }, [isLoginForm]);

  const LoginFormMap = [
    {
      name: "email",
      placeholder: "Email",
      type: "email",
    },
    {
      name: "password",
      placeholder: "Password",
      type: "password",
    },
  ];

  const RegisterFormMap = [
    {
      name: "fullName",
      placeholder: "Full Name",
      type: "text",
    },
    ...LoginFormMap,
  ];

  return (
    <div className={cn("w-full max-w-[350px] flex-col gap-2", className)}>
      {(isLoginForm ? LoginFormMap : RegisterFormMap).map((field, index) => {
        return (
          <Input
            key={`${field.name}-${index}-${field.type}`}
            name={field.name}
            onChange={(isLoginForm ? loginFk : registerFk).handleChange}
            onBlur={(isLoginForm ? loginFk : registerFk).handleBlur}
            customVariant="primary"
            type={field.type}
            placeholder={field.placeholder}
          />
        );
      })}
      <span
        onClick={() => setIsLoginForm((prev) => !prev)}
        className="cursor-pointer text-sm font-normal text-white/70"
      >
        {isLoginForm ? "Dont have Account? Register." : "Have Account? Login."}
      </span>
      <Button
        type="button"
        disabled={
          Object.keys(isLoginForm ? loginFk.errors : registerFk.errors)
            .length !== 0
        }
        className="mt-2 h-12 font-normal text-white/90"
        customVariants="primary"
        onClick={() => (isLoginForm ? loginFk : registerFk).handleSubmit()}
      >
        {isLoginForm ? "Login" : "Create Account"}
      </Button>
    </div>
  );
}

function HomeCarousel({
  handleImageChange: { image, setImage },
  active = 0,
  setActive,
  data,
  onClick,
  showAddAccountForm,
}: {
  handleImageChange: {
    image: IStateImageType;
    setImage: React.Dispatch<React.SetStateAction<IStateImageType>>;
  };
  showAddAccountForm?: boolean;
  onClick?: (idx: number) => void;
  data?: any[];
  active: number;
  setActive: React.Dispatch<React.SetStateAction<number>>;
}) {
  const LENGHT = data?.length || 0;

  useEffect(() => {
    gsap.to(`.active-user-card`, {
      scale: "1.2",
    });
    gsap.to(`.inactive-user-card`, {
      scale: "1",
    });
    gsap.to(`.active-user-avatar`, {
      boxShadow: "2px 0 3px 1px #ffffff",
    });
    gsap.to(`.inactive-user-avatar`, {
      boxShadow: "0px 0 10px 1px transparent",
    });
  }, [active]);

  return (
    <Carousel className="mx-auto w-[90%] md:w-[80%]">
      <CarouselContent>
        <CarouselItem
          className={cn({
            "basis-1/1": LENGHT === 1,
            "basis-1/2": LENGHT === 2,
            "basis-1/3": LENGHT >= 3 && LENGHT < 5,
            "basis-1/3 lg:basis-1/5": LENGHT >= 5,
          })}
        >
          <div
            className={cn(
              "flex items-center justify-center p-1 py-5 pb-[80px]",
            )}
          >
            <User
              handleImageChange={{
                image: image,
                setImage: setImage,
              }}
              variant={"add-account"}
              name={showAddAccountForm ? "Add Profile Picture" : "Add Account"}
              onClick={() => {
                setActive(0);
                if (onClick) {
                  onClick(0);
                }
              }}
              onMouseEnter={() => setActive(0)}
              className={
                0 === active ? "active-user-card" : "inactive-user-card"
              }
              avatarClassName={
                0 === active ? "active-user-avatar" : "inactive-user-avatar"
              }
            />
          </div>
        </CarouselItem>
        {Array.from({ length: LENGHT }).map((_, index) => (
          <CarouselItem
            key={index + 1}
            className={cn({
              "basis-1/3": LENGHT < 5,
              "basis-1/3 lg:basis-1/5": LENGHT >= 5,
            })}
          >
            <div
              className={cn(
                "flex items-center justify-center p-1 py-5 pb-[80px]",
              )}
            >
              <User
                onClick={() => {
                  setActive(index + 1);
                  if (onClick) {
                    onClick(index + 1);
                  }
                }}
                onMouseEnter={() => setActive(index + 1)}
                name={`khalid.khnz ${index + 1}`}
                className={
                  index + 1 === active
                    ? "active-user-card"
                    : "inactive-user-card"
                }
                avatarClassName={
                  index + 1 === active
                    ? "active-user-avatar"
                    : "inactive-user-avatar"
                }
              />
            </div>
          </CarouselItem>
        ))}
      </CarouselContent>
      {LENGHT !== 0 && (
        <>
          <CarouselPrevious className="carousel__btn" />
          <CarouselNext className="carousel__btn" />
        </>
      )}
    </Carousel>
  );
}

function User({
  handleImageChange,
  name,
  className,
  avatarClassName,
  onClick,
  onMouseEnter,
  variant,
}: {
  handleImageChange?: {
    image: IStateImageType;
    setImage: React.Dispatch<React.SetStateAction<IStateImageType>>;
  };
  name?: string;
  className?: string;
  avatarClassName?: string;
  active?: boolean;
  variant?: "add-account";
  onClick?: () => void;
  onMouseEnter?: () => void;
}) {
  if (variant === "add-account") {
    return (
      <div
        className={cn(
          "flex flex-col items-center justify-center gap-2",
          className,
        )}
        onClick={onClick}
        onMouseEnter={onMouseEnter}
      >
        <div
          className={cn(
            "relative h-[10vw] max-h-[160px] min-h-[110px] w-[10vw] min-w-[110px] max-w-[160px] cursor-pointer",
            "mt-[10px] border-[4px] border-transparent",
            "flex items-center justify-center overflow-hidden rounded-full bg-white/20 backdrop-blur-md",
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
            className="absolute left-0 top-0 h-full w-full opacity-0"
          />
        </div>
        <h2 className="cursor-pointer text-[12px] font-light text-white md:text-[1.5vw] lg:text-[20px]">
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
            className,
          )}
          onClick={onClick}
          onMouseEnter={onMouseEnter}
        >
          <Avatar
            className={cn(
              "h-[10vw] max-h-[160px] min-h-[110px] w-[10vw] min-w-[110px] max-w-[160px] cursor-pointer",
              "mt-[10px] border-[4px] border-transparent",
              avatarClassName,
            )}
            variant={"default"}
          />
          <h2 className="cursor-pointer text-[18px] font-light text-white md:text-[1.8vw] lg:text-[25px]">
            {name}
          </h2>
        </div>
      </HoverCardTrigger>
      <HoverCardContent className="mt-4 flex items-center justify-center border-none bg-white/20 p-2 backdrop-blur-md">
        <span className="overflow-hidden overflow-ellipsis text-nowrap text-sm font-light text-white">
          eternalkhalidkhnz@gmail.com
        </span>
      </HoverCardContent>
    </HoverCard>
  );
}

export default HomePage;
