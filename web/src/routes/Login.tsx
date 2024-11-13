export function Login() {
  return (
    <>
      <form action="">
        <legend>Log in</legend>
        <div>
          <label htmlFor="email">Email</label>
          <input type="text" name="email" id="email" />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input type="password" name="password" id="password" />
        </div>
        <button>Log in</button>
        <div className="">
          <span>
            Don't have an account? <a href="/signup">Sign up</a>
          </span>
        </div>
      </form>
    </>
  );
}
