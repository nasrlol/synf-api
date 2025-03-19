import "./login.css";

function loginPage() {
    return (
        <>

            <section className="signup">
                <form>
                    <input type="text" name="txt" placeholder="User name" required/>
                    <input type="email" name="email" placeholder="Email" required/>
                    <input type="number" name="broj" placeholder="BrojTelefona" required/>
                    <input type="password" name="pswd" placeholder="Password" required/>
                    <button className={"loginButton"}>Sign up</button>
                </form>
            </section>

            <section className="login">
                <form>
                    <input type="email" name="email" placeholder="Email" required/>
                    <input type="password" name="pswd" placeholder="Password" required/>
                    <button className={"loginButton"}>Login</button>
                </form>
            </section>
        </>
    );
}

export default loginPage;
