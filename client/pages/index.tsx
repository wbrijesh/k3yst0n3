import { GetCookie, SetCookie } from "@/helpers/cookie";
import ParseJWT from "@/helpers/jwt";
import Link from "next/link";
import React from "react";

function IndexPage() {
  const [tokenData, setTokenData] = React.useState<any>()

  React.useEffect(() => {
    setTokenData(ParseJWT(GetCookie("token") || "{}"))
  }, []);

  return (
    <div className="flex flex-col gap-1 m-4">
      <h1>K3YST0N3</h1>
      {tokenData
        ? <div>
          <pre>{JSON.stringify(tokenData, null, 2)}</pre>
          <button className="border border-gray-400 rounded px-2 py-1" onClick={() => {
            SetCookie("token", "", -1)
            setTokenData(undefined)
          }}>Logout</button>
        </div>
        : <>
          <Link
            href="/login"
            className="text-blue-600 underline"
          >Login</Link>
          <Link
            href="/register"
            className="text-blue-600 underline"
          >Register</Link>
        </>
      }
    </div>
  );
}

export default IndexPage;
