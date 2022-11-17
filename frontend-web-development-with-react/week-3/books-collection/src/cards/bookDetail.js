export function BookDetailCard(props) {
    return (
        <>
            <div>
                Title: {props.book.title}
            </div>
            <div>
                Description: {props.book.description}
            </div>
        </>
    )
}
