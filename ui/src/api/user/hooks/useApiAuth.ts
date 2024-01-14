import { useCallback, useMemo } from "react";

import { clearCookie } from "src/helpers/cookie";

import { useApi } from "src/services/api/hooks/useApi";

import { BaseUseApiProps } from "src/types/api/api.types";
import { UserLoginResponse } from "src/types/api/user.types";
import { CookieKey } from "src/types/cookie.types";

import { makeApiGoogleLogin, makeApiLogout } from "../functions/makeApiAuth";

export const useApiGoogleLogin = ({
  onSuccess,
  onResponse,
  onError,
}: BaseUseApiProps<UserLoginResponse> = {}) => {
  const doRequest = useMemo(() => makeApiGoogleLogin(), []);
  return useApi({
    doRequest,
    onSuccess,
    onResponse,
    onError,
  });
};

export const useApiLogout = ({
  onSuccess,
  onResponse,
  onError,
}: BaseUseApiProps<void> = {}) => {
  const doRequest = useMemo(() => makeApiLogout(), []);
  const handleOnResponse = useCallback(() => {
    clearCookie(CookieKey.TOKEN);
    onResponse && onResponse();
  }, [onResponse]);

  return useApi({
    doRequest,
    onSuccess,
    onResponse: handleOnResponse,
    onError,
  });
};
