import {
  TokenResponse,
  useGoogleLogin as useGoogle,
} from "@react-oauth/google";
import { useCallback } from "react";

import { useApiGoogleLogin } from "src/api/user/hooks/useApiAuth";

import { BaseUseApiProps } from "src/types/api/api.types";
import { UserLoginResponse } from "src/types/api/user.types";

type UseApiGoogleLoginProps = {
  onSuccessLogin?: BaseUseApiProps<UserLoginResponse>["onSuccess"];
  onErrorLogin?: BaseUseApiProps<UserLoginResponse>["onError"];
  onResponse?: BaseUseApiProps<UserLoginResponse>["onResponse"];
};

type UseApiGoogleLoginResult = {
  googleLogin: () => void;
};

const useGoogleLogin = ({
  onSuccessLogin,
  onErrorLogin,
  onResponse,
}: UseApiGoogleLoginProps = {}): UseApiGoogleLoginResult => {
  const { doRequest } = useApiGoogleLogin({
    onSuccess: onSuccessLogin,
    onError: onErrorLogin,
    onResponse: onResponse,
  });

  const googleLogin = useGoogle({
    flow: "implicit",
    onSuccess: useCallback(
      (
        tokenResponse: Omit<
          TokenResponse,
          "error" | "error_description" | "error_uri"
        >,
      ) => {
        doRequest({ token: tokenResponse.access_token });
      },
      [doRequest],
    ),
  });

  return { googleLogin };
};

export default useGoogleLogin;
