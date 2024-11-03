import AuthContext from "./authContext";

type LoginResponse = {
  message: string;
  data: any;
  token: string;
};

import { ReactNode } from "react";

const AuthProvider = ({ children }: { children: ReactNode }) => {
  const login = (user: LoginResponse) => {
    localStorage.setItem("token", user.token);
  };

  return (
    <AuthContext.Provider value={{ login }}>{children}</AuthContext.Provider>
  );
};

export default AuthProvider;
