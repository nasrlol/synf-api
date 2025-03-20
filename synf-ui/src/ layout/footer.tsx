import { Link } from "react-router-dom";

export default function footer() {
	return (
		<>
			<footer>
				<ul>
					<li>
						<Link to="/">SYNF &copy;</Link>
					</li>
					<li>
						<Link to="https://www.github.com/nasrlol">GITHUB</Link>
					</li>
					<li>
						<Link to="https://www.linkedin.com/in/abdellah-el-morabit-578a51324/">LINKEDIN</Link>
					</li>
				</ul>
			</footer>
		</>
	);
}
