import { ENV } from "./Env";
export const ENDPOINTS = {
  AUTH: {
    CREATE_ACCOUNT: (mode?: "CREATOR") =>
      `${ENV.API_ENDPOINT}/auth/signup${mode ? `?mode=${mode}` : ""}`,
    LOGIN_ACCOUNT: `${ENV.API_ENDPOINT}/auth/login`,
    CURRENT_USER: `${ENV.API_ENDPOINT}/auth/current-user`,
    GENERATE_WS_TOKEN: `${ENV.API_ENDPOINT}/auth/ws-token`,
    GET_USER_BY_ID: (_id: string) => `${ENV.API_ENDPOINT}/auth/get/${_id}`,
  },
};
