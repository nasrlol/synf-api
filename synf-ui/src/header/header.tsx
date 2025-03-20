import "./header.css";
import {Link} from "react-router-dom"

function header() {
    return (
        <header>
            <nav>
                <ul>
                    <li><Link to="/">SYNF</Link></li>
                    <li><Link to="/install">INSTALL</Link></li>
                    <li><Link to="/about">ABOUT</Link></li>
                    <li><Link to="/login">MYSYNF</Link></li>
                </ul>
            </nav>
        </header>
    );
}

export default header;
