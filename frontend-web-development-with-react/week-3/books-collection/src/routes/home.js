import { Link } from "react-router-dom"

export function HomeRoute() {
    return (
        <>
            <div>
                Welcome to my Books Collection
            </div>
            <div>
                <Link to="/book">Enter</Link>
            </div>
        </>
    )
}
