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
            isLoaded: false,
            projectData: []
        }
    }


    async componentDidMount() {
       const url ="http://10.212.138.205:8080/stillastracking/v1/api/project/";
        fetch(url)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        projectData: result
                    });
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    this.setState({
                        isLoaded: true,

                    });
                }
            )
    }



    render() {
        const {projectData } = this.state;
        return(
            <div className={"grid-container"}>
                {projectData.map((e) =>{
                return(
                    <CardElement key = {e.projectID}
                                 name = {e.projectName}
                                 state = {e.state}
                                 rentPeriod = {e.period.startDate}
                                 size = {e.size}
                                 contactPerson = {e.customer.name}
                                 contactNumber = {e.customer.number}
                                 address_Street = {e.address.street}
                                 address_Municipality = {e.address.municipality}
                                 address_zip = {e.address.zipcode}
                    />
                );
                })}
            </div>
        );
    }
}

export default Projects;
