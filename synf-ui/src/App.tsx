import "./App.css";
import header from "./header/header.tsx";
import footer from "./footer/footer.tsx";
import home from "./views/home/home.tsx"

function App() {
    return (
        <>
            {header()}
            <main>
                {home()}
            </main>
            {footer()}
        </>
    )
}

export default App;
