import "./install.css";

function install() {
	return (
		<>
			<section className="install">
				<h1>Installation Guide</h1>
				<h2>Prerequisites</h2>
				<p>Before you begin, ensure you have the following installed on your system:</p>
				<ul>
					<li>
						<strong>Go </strong> (for the API server)
					</li>
					<li>
						<strong>Clang </strong> (for the C backend)
					</li>
					<li>
						<strong>MariaDB </strong> (for remote user data storage)
					</li>
					<li>
						<strong>Git</strong> (to clone the repository)
					</li>
				</ul>

				<h2>Clone the Repository</h2>
				<pre>
					<code>git clone https://github.com/nasrlol/synf.git cd synf</code>
				</pre>
				<h2>Set up the API (Go Server)</h2>
				<pre>
					<code>cd api go mod tidy go run main.go</code>
				</pre>
				<p>
					By default, the API runs on <strong>port 8080</strong>.
				</p>

				<h2>Set up the Frontend (React)</h2>
				<pre>
					<code>cd frontend npm install npm start</code>
				</pre>
				<p>
					By default, the frontend is available at <strong>http://localhost:3000</strong>.
				</p>

				<h2>Access the Application</h2>
				<ul>
					<li>
						Open <strong>http://localhost:3000</strong> in your browser to access the React frontend.
					</li>
					<li>
						The backend and API are accessible on ports <strong>5000</strong> and <strong>8080</strong>{" "}
						respectively.
					</li>
				</ul>
			</section>
		</>
	);
}

export default install;
