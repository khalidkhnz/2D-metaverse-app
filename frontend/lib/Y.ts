import * as yup from "yup";

export class Y {
  public static register = yup.object().shape({
    fullName: yup.string().required("Full name is required"),
    email: yup
      .string()
      .email("Invalid email format")
      .required("Email is required"),
    password: yup
      .string()
      .min(8, "Password must be at least 8 characters")
      .required("Password is required"),
    permissions: yup.array(),
  });

  public static login = yup.object().shape({
    email: yup
      .string()
      .email("Invalid email format")
      .required("Email is required"),
    password: yup.string().required("Password is required"),
  });
}
