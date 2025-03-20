import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import header from "./ layout/header.tsx";
import footer from "./ layout/footer.tsx";
import Home from "./views/home/home.tsx";
import Login from "./views/login/login.tsx";
import About from "./views/about/about.tsx";
import Install from "./views/install/install.tsx";
import Signup from "./views/signup/signup.tsx";
import Dashboard from "./views/dashboard/dashboard.tsx";

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
