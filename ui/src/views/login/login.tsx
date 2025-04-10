import "./login.css";
import { Link } from "react-router-dom";

function loginPage() {
	return (
		<>
			<section className="login">
				<form>
					<input type="email" name="email" placeholder="Email" required />
					<input type="password" name="pswd" placeholder="Password" required />
					<br />
					<button className={"loginButton"}>Login</button>
				</form>
				<button>
					<Link to="/Signup">Create an account</Link>
				</button>
			</section>
		</>
	);
}

export default loginPage;
