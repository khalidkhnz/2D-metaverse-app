type RequestHandler = (request: { input: RequestInfo; init?: RequestInit }) => {
  input: RequestInfo;
  init?: RequestInit;
};
type ResponseHandler = (response: Response) => Response;

interface InterceptorHandler<T> {
  use: (handler: T) => void;
  handler: T;
}

export interface FetchInterceptor {
  request: InterceptorHandler<RequestHandler>;
  response: InterceptorHandler<ResponseHandler>;
}

async function createInterceptor(): Promise<FetchInterceptor> {
  const interceptors: FetchInterceptor = {
    request: {
      use: (handler) => {
        interceptors.request.handler = handler;
      },
      handler: (request) => {
        const headers = new Headers(request.init?.headers);

        const token = localStorage.getItem("token")
          ? JSON.parse(localStorage.getItem("token") || "")
          : null;

        if (token?.value) {
          headers.append("Authorization", `Bearer ${token?.value}`);
        }

        return { ...request, init: { ...request.init, headers } };
      },
    },
    response: {
      use: (handler) => {
        interceptors.response.handler = handler;
      },
      handler: (response) => response,
    },
  };

  return interceptors;
}

// Enhanced fetch function, uses the interceptor
async function enhancedFetch(
  input: RequestInfo,
  init?: RequestInit,
): Promise<Response> {
  const interceptor = await createInterceptor(); // Wait for interceptor creation
  const modifiedRequest = interceptor.request.handler({ input, init });
  const response = await fetch(modifiedRequest.input, modifiedRequest.init);
  return interceptor.response.handler(response);
}

// Define custom fetch methods for API calls
export async function getMethod<T>(
  url: string,
  options?: RequestInit,
): Promise<T> {
  return enhancedFetch(url, { ...options, method: "GET" }).then((r) =>
    r.json(),
  );
}

export async function postMethod<T>(
  url: string,
  data?: any,
  options?: RequestInit,
): Promise<T> {
  return enhancedFetch(url, {
    ...options,
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json",
      ...(options?.headers || {}),
    },
  }).then((r) => r.json());
}

export async function putMethod<T>(
  url: string,
  data?: any,
  options?: RequestInit,
): Promise<T> {
  return enhancedFetch(url, {
    ...options,
    method: "PUT",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json",
      ...(options?.headers || {}),
    },
  }).then((r) => r.json());
}

export async function patchMethod<T>(
  url: string,
  data?: any,
  options?: RequestInit,
): Promise<T> {
  return enhancedFetch(url, {
    ...options,
    method: "PATCH",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json",
      ...(options?.headers || {}),
    },
  }).then((r) => r.json());
}

export async function deleteMethod<T>(
  url: string,
  options?: RequestInit,
): Promise<T> {
  return enhancedFetch(url, { ...options, method: "DELETE" }).then((r) =>
    r.json(),
  );
}

export async function getSwrMethod<T>(url: string): Promise<T> {
  return await fetch(url).then((r) => r.json());
}
