import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import header from "./header/header.tsx";
import footer from "./footer/footer.tsx";
import home from "./pages/home_page/home.tsx"

function App() {
	return (
		<Router>
			<Routes>
				<Route path = "" element = {
					<>
						{header()}
						<main>
							{home()}	
						</main>
						{footer()}
					</>
					} />
			</Routes>
		</Router>
	);
}

export default App;
