import "./App.css";

function App() {
	return (
		<>
			<header>
				<ul>
					<li>
						<a href={"/"}>home</a>
					</li>
					<li>
						<a href={"install_guide"}>install</a>
					</li>
					<li>
						<a href={"mysynf_dashboard"}>mySYNF</a>
					</li>
				</ul>
			</header>

			<main>
				<h1>
					SYNF <br /> your device's health anywhere and anytime
				</h1>
				<p>Track your device's health and get instant alerts on issues where ever you are!</p>
				<button>learn more</button>
			</main>

			<footer>
				<p>&copy; 2025 SYNF Inc. All rights reserved</p>
			</footer>
		</>
	);
}

export default App;
