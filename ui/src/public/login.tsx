import "./public.css";
import { Link } from "react-router-dom";

function loginPage() {
	return (
		<>
			<section className="reg">
				<form>
					<input type="email" name="email" placeholder="Email" required />
					<input type="password" name="pswd" placeholder="Password" required />
					<br />
					<button>
						<Link to="/Login/Dashboard">Login</Link>
					</button>
				</form>
				<button>
					<Link to="/Signup">Create an account</Link>
				</button>
			</section>
		</>
	);
}

export default loginPage;
