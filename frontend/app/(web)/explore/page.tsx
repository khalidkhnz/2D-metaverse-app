"use client";

import Loader from "@/components/loader";
import { Input } from "@/components/ui/input";
import { SpaceService } from "@/services/api/space.service";
import { useQuery } from "@tanstack/react-query";
import gsap from "gsap";
import dynamic from "next/dynamic";
import { useRouter } from "next/navigation";
import React, { Suspense, useEffect } from "react";

const AioIcon = dynamic(() => import("@/components/aio-icon"), {
  ssr: false,
});

export default function App() {
  return (
    <Suspense>
      <Page />
    </Suspense>
  );
}

const Page = () => {
  const router = useRouter();
  const { data: MySpacesResponse, isLoading: MySpacesLoading } = useQuery({
    queryKey: ["my-spaces"],
    queryFn: SpaceService.handleGetAllMySpace,
  });

  useEffect(() => {
    if (!MySpacesLoading) {
      gsap.to(".my-spaces-parent", {
        opacity: 1,
        y: 0,
        duration: 1,
      });
    }
  }, [MySpacesLoading]);

  if (MySpacesLoading)
    return (
      <Loader className="relative flex h-screen w-full items-center justify-center" />
    );

  return (
    <div className="explore-page relative flex h-screen w-full justify-center overflow-hidden text-white">
      <div className="my-spaces-parent hide_scrollbar flex h-full w-full translate-y-[200px] gap-10 overflow-x-auto pr-[100px] opacity-0">
        <AioIcon
          variant="add-account"
          enableAnimation
          onClick={() => router.push("/create-space")}
          animationId={"add-space"}
          disableToolTip
          className="my-auto h-fit w-full min-w-[200px] max-w-[200px] pl-[100px] text-center"
          textClassName="text-nowrap text-sm lg:text-sm ellipse "
          name={"Create Space"}
        />
        {(Array.isArray(MySpacesResponse?.data)
          ? MySpacesResponse.data
          : []
        ).map((space, idx) => {
          return (
            <AioIcon
              enableAnimation
              onClick={() => router.push(`/space/${space?.name}`)}
              animationId={space._id}
              disableToolTip
              className="my-auto h-fit w-full min-w-[200px] max-w-[200px] text-center"
              textClassName="text-nowrap text-sm lg:text-sm ellipse "
              name={space.name}
              key={`space-${idx}-${space._id}`}
            />
          );
        })}
      </div>
    </div>
  );
};
