import "./public.css";
import { Link } from "react-router-dom";

function signin() {
	return (
		<>
			<section className="reg">
				<form>
					<input type="text" name="txt" placeholder="User name" required />
					<input type="email" name="email" placeholder="Email" required />
					<input type="password" name="pswd" placeholder="Password" required />
					<button>
						<Link to="/signup">Sign up</Link>
					</button>
				</form>
				<button>
					<Link to="/login">Already have an account?</Link>
				</button>
			</section>
		</>
	);
}

export default signin;
