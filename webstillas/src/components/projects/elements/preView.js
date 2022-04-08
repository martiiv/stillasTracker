import React from "react";
import {useLocation, useParams} from "react-router";



class PreView extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            data:[]
        }
    }


    getProjectID(){
        const pathSplit = window.location.href.split("/")
        return pathSplit[pathSplit.length - 1]
    }


    async componentDidMount() {
        const path = this.getProjectID()
        console.log(path)
        const url ="http://10.212.138.205:8080/stillastracking/v1/api/project/?id=" + path;
        fetch(url)
            .then(res => res.json())
            .then(
                (result) => {
                    sessionStorage.setItem('project', (result))
                    this.setState({
                        isLoaded: true,
                        projectData: result
                    });
                },
                (error) => {
                    this.setState({
                        isLoaded: true,

                    });
                }
            )

    }

    render() {
        const {projectData} = this.state
        {console.log(projectData)}
        return(
            <div className="container">
                <div>
                    <h1>Hello World</h1>
                    <div>{JSON.stringify(projectData)}</div>
                </div>

            </div>
        )
    }

}


export default PreView




