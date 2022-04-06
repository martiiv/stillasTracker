import React from "react";
import "./projects.css"
import  CardElement from './elements/card'
import data from "bootstrap/js/src/dom/data";

/**
 Class that will create an overview of the projects
 */

class Projects extends React.Component {
    constructor(props) {
        super(props);
        this.state ={
            projectData: null
        }
    }


    async componentDidMount() {
        const url = 'http://localhost:8080/stillastracking/v1/api/project/'

       //const url ="http://10.212.138.205:8080/stillastracking/v1/api/project/";
        const response = await fetch(url);
        const data = await response.json();
        this.setState({ projectData: data })
        console.log(this.state.projectData)
    }


    render() {
        return(
            <div className={"grid-container"}>
                <div>Data is:</div>
                {/* <CardElement />*/}
            </div>
        );
    }
}

export default Projects;
