import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import header from "./layout/header.tsx";
import footer from "./layout/footer.tsx";
import Home from "./home/home.tsx";
import Login from "./public/login.tsx";
import About from "./pages/about.tsx";
import Install from "./pages/install.tsx";
import Signup from "./public/signup.tsx";
import Dashboard from "./data/dashboard.tsx";

function App() {
	return (
		<Router>
			<>
				{header()}
				<main>
					<Routes>
						<Route path="/" element={<Home />} />
						<Route path="/install" element={<Install />} />
						<Route path="/about" element={<About />} />
						<Route path="/login" element={<Login />} />
						<Route path="/Signup" element={<Signup />} />
						<Route path="/Dashboard" element={<Dashboard />} />
					</Routes>
				</main>

				{footer()}
			</>
		</Router>
	);
}

export default App;
