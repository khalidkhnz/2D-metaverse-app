interface Interceptor<T> {
  use(handler: (input: T) => T): void;
  handler: (input: T) => T;
}

interface RequestInterceptor {
  input: RequestInfo;
  init?: RequestInit;
}

interface ResponseInterceptor {
  ok: boolean;
  status: number;
  statusText: string;
  headers: Headers;
  text: () => Promise<string>;
  json: () => Promise<any>;
}

type ApiResponse<T = unknown> = T;

interface FetchInterceptor {
  request: Interceptor<RequestInterceptor>;
  response: Interceptor<ResponseInterceptor>;
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

async function enhancedFetch(
  input: RequestInfo,
  init?: RequestInit,
): Promise<any> {
  const interceptor = await createInterceptor(); // Wait for interceptor creation
  const modifiedRequest = interceptor.request.handler({ input, init });
  const response = await fetch(modifiedRequest.input, modifiedRequest.init);
  return interceptor.response.handler(response);
}

//////////////////////////
/* CUSTOM FETCH METHODS */
//////////////////////////

export async function getMethod<T>(
  url: string,
  options?: RequestInit,
): Promise<ApiResponse<T>> {
  return enhancedFetch(url, { ...options, method: "GET" }).then((r) =>
    r.json(),
  );
}

export async function postMethod<T>(
  url: string,
  data?: any,
  options?: RequestInit,
): Promise<ApiResponse<T>> {
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
): Promise<ApiResponse<T>> {
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
): Promise<ApiResponse<T>> {
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
): Promise<ApiResponse<T>> {
  return enhancedFetch(url, { ...options, method: "DELETE" }).then((r) =>
    r.json(),
  );
}

export async function getSwrMethod<T>(url: string): Promise<ApiResponse<T>> {
  return await fetch(url).then((r) => r.json());
}
