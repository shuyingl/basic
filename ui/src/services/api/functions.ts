import { AxiosResponse } from "axios";

export const getApiMessage = (
  response: AxiosResponse | undefined,
): string | undefined =>
  response?.data?.message ?? response?.statusText ?? undefined;
