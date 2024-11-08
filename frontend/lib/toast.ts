import { ReactNode } from "react";
import { ExternalToast, toast } from "sonner";

type titleT = (() => React.ReactNode) | React.ReactNode;
type PromiseT<Data = any> = Promise<Data> | (() => Promise<Data>);

export class Toast {
  public static default(
    message: titleT,
    data?: ExternalToast
  ): string | number {
    return toast(message, data);
  }

  public static success(
    message: titleT | ReactNode,
    data?: ExternalToast
  ): string | number {
    return toast.success(message, data);
  }

  public static error(
    message: titleT | ReactNode,
    data?: ExternalToast
  ): string | number {
    return toast.error(message, data);
  }

  public static warning(
    message: titleT | ReactNode,
    data?: ExternalToast
  ): string | number {
    return toast.warning(message, data);
  }

  public static info(
    message: titleT | ReactNode,
    data?: ExternalToast
  ): string | number {
    return toast.info(message, data);
  }

  public static promise<ToastData>(
    promise: PromiseT<ToastData>,
    data?: any
  ):
    | (string & {
        unwrap: () => Promise<ToastData>;
      })
    | (number & {
        unwrap: () => Promise<ToastData>;
      })
    | {
        unwrap: () => Promise<ToastData>;
      } {
    return toast.promise(promise, data);
  }
}
