import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { Home, Login, Signup } from "./routes";

function App() {
  return (
    <>
      <Router>
        <header className="h-16 bg-carb-black flex items-center justify-between px-12 text-carb-white shadow-lg sticky top-0">
          <Link to="/" className="text-carb-green-light font-bold text-xl">
            carbonara.
          </Link>
          <div className="flex gap-3 font-bold items-center">
            <Link to="login">Log in</Link>
            <Link
              to="signup"
              className="bg-carb-green-light px-6 py-2 rounded-md"
            >
              Get started
            </Link>
          </div>
        </header>
        <main className="bg-carb-black min-h-screen w-full text-carb-white">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<Signup />} />
          </Routes>
        </main>
      </Router>
    </>
  );
}

export default App;
