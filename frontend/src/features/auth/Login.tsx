import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { InferType } from "yup";
import api from "../../api/api";
import useAuth from "./useAuth";

const loginSchema = yup.object({
  email: yup.string().email().required(),
  password: yup.string().required(),
});

type FormData = InferType<typeof loginSchema>;

const Login = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: yupResolver(loginSchema),
  });

  const { login } = useAuth();

  const onSubmit = async (data: FormData) => {
    try {
      const response = await api.post("/auth/login", data);
      console.log(response.data);
      login(response.data);
    } catch (error) {
      throw new Error("failed login");
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <label htmlFor="email">email: </label>
      <input type="email" id="email" {...register("email")} />
      <p>{errors.email?.message}</p>
      <label htmlFor="password">password: </label>
      <input type="password" id="password" {...register("password")} />
      <p>{errors.password?.message}</p>
      <button type="submit">login</button>
    </form>
  );
};
export default Login;
