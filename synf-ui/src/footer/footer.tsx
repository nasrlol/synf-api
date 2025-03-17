import "./footer.css";
import { Link } from "react-router-dom"

export default function footer() {
	return (
		<>
			<footer>
				<p><Link to="https://www.github.com">GITHUB</Link> | <Link to="https://www.linkedin.com/in/abdellah-el-morabit-578a51324/">LINKEDIN</Link> | <Link to="/">SYNF</Link></p>
			</footer>
		</>
	);
}
