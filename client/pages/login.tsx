import { SetCookie } from '@/helpers/cookie';
import Link from 'next/link';
import { useRouter } from 'next/router';
import React from 'react';
import { useForm, SubmitHandler } from "react-hook-form"

type Inputs = {
  email: string
  password: string
}

function LoginPage() {
  const router = useRouter()

  const [response, setResponse] = React.useState<any>()

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Inputs>()
  const onSubmit: SubmitHandler<Inputs> = (data) => {
    console.log(data)

    fetch("http://localhost:4000/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data)
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data)
        setResponse(data)
        if (data.message === "User logged in successfully") {
          SetCookie("token", data.token, 86400)
          router.push("/")
        }
      })
      .catch((error) => {
        console.error("Error:", error)
        setResponse(error)
      })
  }

  return (
    <div className="flex flex-col gap-1 m-4">
      <h1>Login</h1>
      <form onSubmit={handleSubmit(onSubmit)} className="border border-gray-400 rounded w-[370px] p-[10px]">
        <div className='flex flex-col w-[350px] mb-1'>
          <label>Email</label>
          <input
            {...register("email", { required: true })} type="email"
            className="border border-gray-400 px-2 py-1 rounded"
          />
          {errors.email && <span>This field is required</span>}
        </div>
        <div className='flex flex-col w-[350px] mb-3'>
          <label>Password</label>
          <input
            {...register("password", { required: true })} type="password"
            className="border border-gray-400 px-2 py-1 rounded"
          />
          {errors.password && <span>This field is required</span>}
        </div>
        {response && response.error && <div className="my-2 text-red-600">{response.error}</div>}
        <button className="border border-gray-400 px-2 py-1 hover:cursor-pointer hover:bg-gray-100" type="submit">Submit</button>
      </form>
      <div className="mt-2 flex gap-1">
        Don't have an account?
        <Link href="/register" className="text-blue-600 underline">Register</Link>
      </div>
    </div>
  );
}

export default LoginPage;
