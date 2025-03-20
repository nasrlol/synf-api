import hero from "../hero/hero.tsx";
import "./home.css";

function home() {
	return (
		<>
			<div className="home">
				{hero()}
				<section className="introduction">
					<h2>Why Choose Our Monitoring Solution?</h2>
					<p>
						Our system monitoring tool provides real-time insights into CPU, memory, storage, and network
						usage. With smart alerts and strong security, it ensures optimal performance and privacy.
					</p>
				</section>

				<section className="features">
					<article>
						<h3>Real-Time Monitoring</h3>
						<p>Track CPU, RAM, storage, and network usage live.</p>
					</article>
					<article>
						<h3>Smart Alerts</h3>
						<p>Receive instant notifications for unusual device activity.</p>
					</article>
					<article>
						<h3>Secure & Private</h3>
						<p>Your data is encrypted, ensuring your privacy.</p>
					</article>
					<article>
						<h3>Detailed Reports</h3>
						<p>View historical logs and trends for better system insights.</p>
					</article>
					<article>
						<h3>Lightweight & Fast</h3>
						<p>Designed to run efficiently without slowing down your device.</p>
					</article>
				</section>

				<section className="benefits">
					<h2>What Makes This Unique?</h2>
					<ul>
						<li>️Open-source and customizable</li>
						<li>️Supports multiple operating systems</li>
						<li>️Integrates easily with existing infrastructure</li>
						<li>️Cloud and on-premise deployment options</li>
					</ul>
				</section>

				<section className="how-it-works">
					<h2>How It Works</h2>
					<ol>
						<li>Install the monitoring agent on your device.</li>
						<li>Connect it to the dashboard for live insights.</li>
						<li>Customize alerts and notifications based on thresholds.</li>
						<li>View detailed reports and optimize performance.</li>
					</ol>
				</section>

				<button className="cta-button">Get Started</button>
			</div>
		</>
	);
}

export default home;
