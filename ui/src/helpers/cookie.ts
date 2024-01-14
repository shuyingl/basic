import { CookieKey } from "src/types/cookie.types";

export const clearCookie = (name: CookieKey) => {
  document.cookie = `${name}= `;
};
