import { useMemo } from "react";

import { useApi } from "src/services/api/hooks/useApi";

import { BaseUseApiProps } from "src/types/api/api.types";
import { CurrentUserResponse } from "src/types/api/user.types";

import { makeApiGetUser } from "../functions/makeApiGet";

export const useApiGetCurrentUser = ({
  onSuccess,
  onResponse,
  onError,
}: BaseUseApiProps<CurrentUserResponse> = {}) => {
  const doRequest = useMemo(() => makeApiGetUser(), []);
  return useApi({
    doRequest,
    onSuccess,
    onResponse,
    onError,
  });
};
