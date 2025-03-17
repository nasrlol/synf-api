import "./home.css"
import button from "../../components/button/button.tsx" 

function home()
{
	return (
	<>
		<section>
			<h1>SYNF</h1>
		</section>
		<section>
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
		</section>
		{button("Get Started")}
		</> 
	)
}

export default home;
