/*
Error function that will be displayed if the user requests an invalid url.
 */
export function NotFound() {
    return (
        <h1>
            404 Page Not Found
        </h1>
    )
}

/*
Error function that will be displayed if a server Indicators occurs.
 */
export function InternalServerError() {
    return (
        <h1>
            500 Internal Server Error
        </h1>
    )
}


export function AlertCatch(){
    window.alert("En feil oppstod! Prøv igjen senere.")
}
