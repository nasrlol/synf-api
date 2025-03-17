import "./login.css";

export function loginForms() {
	return (
		<>
			<label>
				<input type="email" name="email" />
				<input type="password" name="password" />
			</label>
		</>
	);
}

function loginPage() {
	return (
		<>
		 { loginForms() }
		</>
	);
}

export default loginPage;
