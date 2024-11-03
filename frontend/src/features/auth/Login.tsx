import { useForm } from "react-hook-form";

const Login = () => {
  const { register, handleSubmit } = useForm<{
    email: string;
    password: string;
  }>();
  const onSubmit = (data: { email: string; password: string }): void => {
    console.log(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <label htmlFor="email"></label>
      <input type="email" id="email" {...register("email")} />
      <input type="password" id="password" {...register("password")} />
      <button type="submit">login</button>
    </form>
  );
};
export default Login;
