import { URLS } from "src/constants/urls";

import { api } from "src/services/api/axios";

import { UserLoginResponse } from "src/types/api/user.types";

export const makeApiGoogleLogin =
  () =>
  ({ token }: { token: string }) =>
    api.get<UserLoginResponse>(URLS.user.loginGoogle, {
      headers: {
        Authorization: token,
      },
    });

export const makeApiLogout = () => () => api.get<void>(URLS.user.logout);
