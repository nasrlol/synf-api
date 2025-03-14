import "./App.css";
import header from "./header/header.tsx"

function App() {
    return (
        <>
            {header()}
            <main>
                <h1>
                    SYNF <br/> your device's health anywhere and anytime
                </h1>
                <p>
                </p>
                <p>Track your device's health and get instant alerts on issues where ever you are!</p>
                <button>learn more</button>
            </main>

            <footer>
                <p>&copy; 2025 SYNF Inc. All rights reserved</p>
            </footer>
        </>
    );
}

export default App;
