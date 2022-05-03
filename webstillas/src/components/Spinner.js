import {Spinner} from "react-bootstrap";
import React from "react";

export const SpinnerDefault = () =>{
    return(
        <Spinner animation="border" style={{ width: '5rem', height: '5rem' }} role="status">
            <span className="visually-hidden">Loading...</span>
        </Spinner>
    )
}
