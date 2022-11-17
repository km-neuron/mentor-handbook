import { Link } from "react-router-dom"

export function BookCard(props) {
    return (
        <div>
            Title: {props.name}
            <Link to={props.isbn}>Detail</Link>
        </div>
    )
}
