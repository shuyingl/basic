import { AxiosError } from "axios";

export type ApiError = {
  message: string;
};

export type ApiErrorResponse = AxiosError<ApiError>["response"];

export type BaseUseApiProps<ResponseData> = {
  onSuccess?: () => void;
  onResponse?: (data: ResponseData) => void;
  onError?: (error: ApiErrorResponse) => void;
};
