"use client";

import { useFormik } from "formik";
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
import { Input } from "@/components/ui/input";
import { useAppContext } from "@/services/app-context";
import { Y } from "@/lib/Y";
import { BODY } from "@/lib/Forms";
import { IUser } from "@/types/user";
import { useRouter } from "next/navigation";
import AioIcon from "@/components/aio-icon";
import { IStateImageType } from "@/types/other";

export default function HomePage() {
  const { current_user } = useAppContext();
  const router = useRouter();
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
      gsap.to(".home_add_image_picker", {
        pointerEvents: "auto",
      });
      gsap.to(".carousel__btn", {
        opacity: 0,
      });
      gsap.to(".auth__form", {
        translateY: "-8vh",
        display: "flex",
        opacity: 1,
        pointerEvents: "auto",
      });
    } else {
      handleLoginAnimation();
      setTimeout(() => {
        router.push("/explore");
      }, 2000);
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
      translateY: "-8vh",
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
        className="mt-[4vh] w-[100%] justify-center py-4"
      >
        <HomeCarousel
          handleImageChange={{
            image,
            setImage,
          }}
          showAddAccountForm={showAddAccountForm}
          onClick={handleContinue}
          data={showAddAccountForm ? [] : current_user ? [current_user] : []}
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
}

function AuthForm({ className }: { className?: string }) {
  const [isLoginForm, setIsLoginForm] = useState(false);
  const { handleRegister, handleLogin } = useAppContext();

  const registerFk = useFormik({
    initialValues: BODY.AUTH.REGISTER(),
    onSubmit: handleRegister,
    validationSchema: Y.register,
  });

  const loginFk = useFormik({
    initialValues: BODY.AUTH.LOGIN(),
    onSubmit: handleLogin,
    validationSchema: Y.login,
  });

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
        onClick={() => {
          setIsLoginForm((prev) => {
            if (!prev) {
              gsap.to("#welcome-carousel", {
                opacity: 0,
                pointerEvents: "none",
              });
              gsap.to(".home_add_image_picker", {
                pointerEvents: "none",
              });
            } else {
              gsap.to("#welcome-carousel", {
                opacity: 1,
                pointerEvents: "auto",
              });
              gsap.to(".home_add_image_picker", {
                pointerEvents: "auto",
              });
            }
            return !prev;
          });
        }}
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
  data = [],
  onClick,
  showAddAccountForm,
}: {
  handleImageChange: {
    image: IStateImageType;
    setImage: React.Dispatch<React.SetStateAction<IStateImageType>>;
  };
  showAddAccountForm?: boolean;
  onClick?: (idx: number) => void;
  data?: IUser[];
  active: number;
  setActive: React.Dispatch<React.SetStateAction<number>>;
}) {
  const router = useRouter();
  const LENGHT = data?.length + 1;

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
    <Carousel className="mx-auto h-[240px] w-[90%] overflow-visible md:w-[80%] 2xl:h-auto">
      <CarouselContent>
        <CarouselItem
          className={cn({
            "basis-1/2": LENGHT === 2,
            "basis-1/3": LENGHT >= 3 && LENGHT < 5,
            "basis-1/3 lg:basis-1/5": LENGHT >= 5,
          })}
        >
          <div
            className={cn("flex items-center justify-center p-1 py-5 pb-[3vh]")}
          >
            <AioIcon
              animationId={"random"}
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
        {data?.map((user, index) => (
          <CarouselItem
            key={index + 1}
            className={cn({
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
              <AioIcon
                animationId={"random"}
                onClick={() => {
                  setActive(index + 1);
                  if (onClick) {
                    onClick(index + 1);
                    handleLoginAnimation();
                    setTimeout(() => {
                      router.push("/explore");
                    }, 2000);
                  }
                }}
                onMouseEnter={() => setActive(index + 1)}
                name={`${user?.fullName}`}
                email={`${user?.email}`}
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

function handleLoginAnimation() {
  gsap.to(".home-main", {
    pointerEvents: "none",
    opacity: 0,
  });
}
