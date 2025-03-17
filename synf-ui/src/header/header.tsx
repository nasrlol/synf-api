import "./header.css";
import {Link} from "react-router-dom"

function header() {
    return (
        <header>
            <ul>
                <li><Link to="/">Home</Link></li>
                <li><Link to="../pages/installation_page/install_guide.tsx">Install</Link></li>
                <li><Link to="../pages/dashboard_pages/dashboard.tsx">MYSYNF</Link></li>
            </ul>
        </header>
    );
}

export default header;
