import "./install.css"

function install()
{
	return (
		<>

			<section>
				<h1>Installation Guide</h1>
				<hr/>
				<h2>Prerequisites</h2>
				<p>Before you begin, ensure you have the following installed on your system:</p>
				<ul>
					<li><strong>Node.js (&gt;= 18.x.x)</strong> & npm (for the React frontend)</li>
					<li><strong>Go (&gt;= 1.19)</strong> (for the API server)</li>
					<li><strong>GCC (&gt;= 11.x.x)</strong> (for the C backend)</li>
					<li><strong>MariaDB ({">"}= 10.5)</strong> (for remote user data storage)</li>
					<li><strong>Git</strong> (to clone the repository)</li>
					<li><strong>Docker</strong> (optional, for containerized deployment)</li>
				</ul>

				<h2>Clone the Repository</h2>
				<pre><code>git clone https://github.com/your-repo/project-name.git
cd project-name</code></pre>

				<h2>Setup MariaDB</h2>
				<p>Start MariaDB Server (If not installed, install it first):</p>
				<pre><code>sudo systemctl start mariadb
sudo systemctl enable mariadb</code></pre>

				<p>Log in to the MariaDB shell:</p>
				<pre><code>mysql -u root -p</code></pre>

				<p>Execute the following SQL commands:</p>
				<pre><code>CREATE DATABASE system_stats;
CREATE USER 'app_user'@'%' IDENTIFIED BY 'secure password';
GRANT ALL PRIVILEGES ON system_stats.* TO 'app_user'@'%';
FLUSH PRIVILEGES;
exit;</code></pre>

				<h2>Set up the Backend (C Server)</h2>
				<pre><code>cd backend
make
./backend_server</code></pre>
				<p>By default, the backend listens on <strong>port 5000</strong>.</p>

				<h2>Set up the API (Go Server)</h2>
				<pre><code>cd api
go mod tidy
go run main.go</code></pre>
				<p>By default, the API runs on <strong>port 8080</strong>.</p>

				<h2>Set up the Frontend (React)</h2>
				<pre><code>cd frontend
npm install
npm start</code></pre>
				<p>By default, the frontend is available at <strong>http://localhost:3000</strong>.</p>

				<h2>Configuration</h2>
				<p>Create a <code>.env</code> file in the root directory and configure it:</p>
				<pre><code>DATABASE_URL=mariadb://app_user:securepassword@localhost/system_stats
API_URL=http://localhost:8080
BACKEND_URL=http://localhost:5000</code></pre>

				<h2>Access the Application</h2>
				<ul>
					<li>Open <strong>http://localhost:3000</strong> in your browser to access the React frontend.</li>
					<li>The backend and API are accessible on ports <strong>5000</strong> and <strong>8080</strong> respectively.</li>
				</ul>

				<h2>Optional: Docker Deployment</h2>
				<pre><code>docker-compose up --build</code></pre>

				<h2>Troubleshooting</h2>
				<ul>
					<li>Ensure MariaDB is running: <code>sudo systemctl status mariadb</code></li>
					<li>Verify Go API is accessible: <code>curl http://localhost:8080/health</code></li>
					<li>Check backend logs for errors: <code>./backend_server --debug</code></li>
					<br/>
				</ul>

				<hr/>
				<h2>Conclusion</h2>
				<p>Your system statistics and user management application should now be up and running!</p>
			</section>

		</>
	)
}

export default install;
