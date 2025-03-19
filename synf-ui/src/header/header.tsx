import "./header.css";
import {Link} from "react-router-dom"

function header() {
    return (
        <header>
            <nav>
                <ul>
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/install">Install</Link></li>
                    <li><Link to="/login">MYSYNF</Link></li>
                </ul>
            </nav>
        </header>
    );
}

export default header;
