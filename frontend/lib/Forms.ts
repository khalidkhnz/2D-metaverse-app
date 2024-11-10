import { ILogin, IRegister } from "@/types/auth";

export class FormsBody {
  public static register(initialVals?: IRegister): IRegister {
    return {
      fullName: initialVals?.fullName || "",
      email: initialVals?.email || "",
      password: initialVals?.password || "",
      roleId: initialVals?.password || "672c869f830b5de5ae075a6f",
      permissions: initialVals?.permissions || [],
    };
  }

  public static login(initialVals?: ILogin): ILogin {
    return {
      email: initialVals?.email || "",
      password: initialVals?.password || "",
    };
  }
}
