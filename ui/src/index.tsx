import "./index.css";

import { GoogleOAuthProvider } from "@react-oauth/google";
import React from "react";
import ReactDOM from "react-dom/client";

import App from "./App";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement,
);
const app = (
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
root.render(
  <GoogleOAuthProvider
    clientId={process.env.REACT_APP_GOOGLE_CLIENT_ID as string}
  >
    {app}
  </GoogleOAuthProvider>,
);
