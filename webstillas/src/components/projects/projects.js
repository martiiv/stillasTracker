import React from "react";
import "./projects.css"
import  CardElement from './elements/card'

/**
 Class that will create an overview of the projects
 */

class Projects extends React.Component {
    render() {
        return(
            <div className={"main"}>
                <CardElement />

            </div>
        );
    }
}

export default Projects;
