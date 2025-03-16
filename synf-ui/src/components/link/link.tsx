import "./link.css"
import { Link } from "react-router-dom"


function link(destination: string, content: string)
{
	return <Link className="link" to={ destination }>{ content }</Link>;
}

export default link;
