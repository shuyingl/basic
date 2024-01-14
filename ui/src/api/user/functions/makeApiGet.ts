import { URLS } from "src/constants/urls";

import { api } from "src/services/api/axios";

import { CurrentUserResponse } from "src/types/api/user.types";

export const makeApiGetUser = () => () =>
  api.get<CurrentUserResponse>(URLS.user.getCurrentUser);
