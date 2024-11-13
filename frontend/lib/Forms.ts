import { ILogin, IRegister } from "@/types/auth";

export const BODY = {
  AUTH: {
    LOGIN: (body?: ILogin): ILogin => {
      return {
        email: body?.email || "",
        password: body?.password || "",
      };
    },
    REGISTER: (body?: IRegister): IRegister => {
      return {
        fullName: body?.fullName || "",
        email: body?.email || "",
        password: body?.password || "",
        permissions: body?.permissions || [],
        roleId: body?.roleId || "672c869f830b5de5ae075a6f",
      };
    },
  },
  SPACE: {
    CREATE_SPACE: (name?: string) => {
      return {
        name: "",
      };
    },
  },
};
