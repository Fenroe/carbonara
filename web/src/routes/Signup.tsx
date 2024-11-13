import { useState } from "react";

export function Signup() {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [confirmPassword, setConfirmPassword] = useState<string>("");
  return (
    <>
      <section className="max-w-lg">
        <form action="" className="w-full p-6">
          <legend>Sign up</legend>
          <div className="w-full flex flex-col">
            <label htmlFor="email">Email</label>
            <input type="text" name="email" id="email" value={email} />
          </div>
          <div>
            <label htmlFor="password">Password</label>
            <input
              type="password"
              name="password"
              id="password"
              value={password}
            />
          </div>
          <div>
            <label htmlFor="confirmPassword">Confirm Password</label>
            <input
              type="password"
              name="confirmPassword"
              id="confirmPassword"
              value={confirmPassword}
            />
          </div>
          <button>Sign up</button>
          <div className="">
            <span>
              Already have an account? <a href="/login">Log in</a>
            </span>
          </div>
        </form>
      </section>
    </>
  );
}
