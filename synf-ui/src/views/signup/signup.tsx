import "./signup.css";
import { Link } from "react-router-dom";

function signin() {
	return (
		<>
			<section className="signup">
				<form>
					<input type="text" name="txt" placeholder="User name" required />
					<input type="email" name="email" placeholder="Email" required />
					<input type="password" name="pswd" placeholder="Password" required />
					<button className={"loginButton"}>Sign up</button>
				</form>
				<button>
					<Link to="/login">Already have an account?</Link>
				</button>
			</section>
		</>
	);
}

export default signin;
