import "./App.css";

import React, { useEffect, useState } from "react";

import { useApiLogout } from "./api/user/hooks/useApiAuth";
import { useApiGetCurrentUser } from "./api/user/hooks/useApiGet";
import useGoogleLogin from "./hooks/UseGoogleLogin";
import { User } from "./types/dao/user.types";
function App() {
  const [currentUser, setCurrentUser] = useState<User>();
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [attemptedLogin, setAttemptedLogin] = useState(false);

  const { googleLogin } = useGoogleLogin({
    onResponse: (response) => {
      setCurrentUser(response);
      setIsLoggedIn(true);
      setAttemptedLogin(true);
    },
    onErrorLogin: () => setAttemptedLogin(true),
  });

  const { doRequest: getCurrentUser } = useApiGetCurrentUser({
    onResponse: (response) => {
      setCurrentUser(response);
      setIsLoggedIn(true);
    },
    onError: () => setIsLoggedIn(false),
  });

  const { doRequest: logout } = useApiLogout({
    onSuccess: () => {
      setCurrentUser({} as User);
      setIsLoggedIn(false);
    },
    onError: () => {
      setCurrentUser({} as User);
      setIsLoggedIn(false);
    },
  });

  const handleLogout = () => {
    logout();
  };

  useEffect(() => {
    if (!attemptedLogin) {
      getCurrentUser();
      setAttemptedLogin(true);
    }
  }, [getCurrentUser, attemptedLogin]);

  return (
    <div className="App">
      <h1>Basic Login</h1>
      <div className="App-auth-area">
        {isLoggedIn ? (
          <button className="App-button" onClick={handleLogout}>
            Logout
          </button>
        ) : (
          <button className="App-button" onClick={googleLogin}>
            Login
          </button>
        )}
        <div className="email-placeholder">
          {isLoggedIn && currentUser?.email}
        </div>
      </div>
    </div>
  );
}

export default App;
