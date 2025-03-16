import "./login.css";
import header from "../../header/header.tsx";
import footer from "../../footer/footer.tsx";

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

export function loginPage() {
	return (
		<>
			{header()}
			<main>loginForms()</main>
			{footer()}
		</>
	);
}
