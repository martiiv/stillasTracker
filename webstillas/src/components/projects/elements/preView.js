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
        const url ="http://localhost:8080/exchange/v1/diag?limit=heipågeddd" ;
        fetch(url)
            .then(res => res.json())
            .then(
                (result) => {
                    console.log(result)

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
        return(
            <div className="container">
                <div>
                    <h1>Hello World</h1>
                </div>

            </div>
        )
    }

}


export default PreView




