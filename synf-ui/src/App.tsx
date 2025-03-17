import "./App.css";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import header from "./header/header.tsx";
import footer from "./footer/footer.tsx";
import Home from "./views/home/home.tsx";
import Login from "./views/login/login.tsx";
import Install from "./views/install/install.tsx";

function App() {

    return (
        <Router>
            <>
                {header()}
                <main>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/install" element={<Install />} />
                        <Route path="/login" element={<Login />} /> {/* Extra pagina */}
                    </Routes>
                </main>
                {footer()}
            </>
        </Router>

    )
}

export default App;
