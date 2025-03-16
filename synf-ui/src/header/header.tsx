import { Link } from "react-router-dom"
import  link from "../components/link/link.tsx"
import "./header.css";

function header() {
	return (
		<header>
			<ul>
				<li>
					{link("/"), "Home"}		
				</li>
				<li>
					<Link className="link" to="../pages/installation_page/install.tsx">Install</Link>
				</li>
				<li>
					<Link className="link" to="../pages/login_page/login.tsx">mySynf</Link>
				</li>
			</ul>
		</header>
	);
}

export default header;
