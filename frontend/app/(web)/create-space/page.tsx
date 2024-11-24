"use client";

import AioIcon from "@/components/aio-icon";
import Button from "@/components/Button";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { SpaceService } from "@/services/api/space.service";
import { IStateImageType } from "@/types/other";
import { useState } from "react";

const Page = () => {
  const [image, setImage] = useState<IStateImageType>({
    file: null,
    result: null,
  });

  const [name, setName] = useState("");

  return (
    <div className="relative flex h-full w-full justify-center">
      <div className="mt-[15vh] flex w-full max-w-[450px] flex-col gap-3">
        <AioIcon
          enableAnimation
          disableToolTip
          handleImageChange={{ image, setImage }}
          variant="add-account"
          name="Add Space Icon"
          className="mb-10"
          textClassName="text-nowrap text-sm lg:text-sm ellipse "
          animationId="add_space_icon"
        />
        <Input
          name="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="Space Name"
          customVariant="primary"
        />
        <textarea
          placeholder="Description"
          className={cn(
            "flex h-12 w-full rounded-md border-none bg-white/10 px-[20px] py-1 text-base text-white shadow-sm outline-none backdrop-blur-md transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-white/50 focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
            "min-h-[120px] focus-visible:bg-white/20",
          )}
        />
        <Button
          customVariants="primary"
          onClick={() => SpaceService.handleCreateSpace(name)}
        >
          Create
        </Button>
      </div>
    </div>
  );
};

export default Page;
