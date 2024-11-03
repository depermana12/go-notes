import { createContext } from "react";

export interface LoginResponse {
  message: string;
  data: any;
  token: string;
}

export interface AuthContextType {
  login: (user: LoginResponse) => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export default AuthContext;
