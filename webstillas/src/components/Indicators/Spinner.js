import {Spinner} from "react-bootstrap";
import React from "react";
import "../../Assets/Styles/Spinner.css"

/**
 * Function that will return a spinner
 *
 * @returns {JSX.Element}
 * @constructor
 */
export const SpinnerDefault = () =>{
    return(
        <Spinner className={"spinner"} animation="border" style={{ width: '5rem', height: '5rem' }} role="status">
            <span className="visually-hidden">Loading...</span>
        </Spinner>
    )
}
