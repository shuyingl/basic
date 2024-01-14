import { googleLogout } from "@react-oauth/google";
import { AxiosError, AxiosPromise, AxiosResponse } from "axios";
import { useCallback, useState } from "react";

import { clearCookie } from "src/helpers/cookie";

import { ApiError, ApiErrorResponse } from "src/types/api/api.types";
import { CookieKey } from "src/types/cookie.types";

import { getApiMessage } from "../functions";

export type ApiHook<RequestParams extends unknown[], ResponseData> = {
  isLoading: boolean;
  doRequest: (...params: RequestParams) => void;
  responseData: ResponseData | undefined;
  success: boolean | null;
  error: null | {
    message?: string;
    code?: number;
  };
};

export type ApiHookProps<RequestParams extends unknown[], ResponseData> = {
  doRequest: (...params: RequestParams) => AxiosPromise<ResponseData>;
  onResponse?: (responseData: ResponseData) => void;
  onRequestStart?: () => void;
  onError?: (error: ApiErrorResponse) => void;
  onSuccess?: () => void;
};

export const useApi = <Params extends unknown[], ResponseData>({
  doRequest,
  onRequestStart,
  onResponse,
  onSuccess,
  onError,
}: ApiHookProps<Params, ResponseData>): ApiHook<Params, ResponseData> => {
  const [responseData, setResponseData] =
    useState<ApiHook<Params, ResponseData>["responseData"]>();
  const [isLoading, setIsLoading] =
    useState<ApiHook<Params, ResponseData>["isLoading"]>(false);
  const [success, setSuccess] =
    useState<ApiHook<Params, ResponseData>["success"]>(null);
  const [error, setError] =
    useState<ApiHook<Params, ResponseData>["error"]>(null);

  const doRequestEnhanced = useCallback(
    (...params: Params) => {
      setIsLoading(true);
      onRequestStart && onRequestStart();
      doRequest(...params)
        .then((response: AxiosResponse<ResponseData>) => {
          try {
            const data = response.data;
            setResponseData(data);
            setError(null);
            setSuccess(true);
            setIsLoading(false);
            onResponse && onResponse(data);
            onSuccess && onSuccess();
          } catch (error) {
            console.warn(error);
          }
        })
        .catch((error: AxiosError<ApiError>) => {
          setSuccess(false);
          setResponseData(undefined);
          setError({
            message: getApiMessage(error.response),
            code: error.response?.status,
          });
          setIsLoading(false);
          if (
            error?.response?.status === 401 ||
            error?.response?.data?.message === "Invalid token."
          ) {
            clearCookie(CookieKey.TOKEN);
            googleLogout();
          }
          onError && onError(error.response);
        });
    },
    [doRequest, onError, onRequestStart, onResponse, onSuccess],
  );

  return {
    doRequest: doRequestEnhanced,
    responseData,
    success,
    error,
    isLoading,
  };
};
