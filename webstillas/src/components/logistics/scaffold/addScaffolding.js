import React from 'react'
import {FormSelect} from "react-bootstrap";
import Form from "react-bootstrap/Form";

class AddScaffolding extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            scaffolding :
                {
                    id: 0,
                    type: "",
                    batteryLevel: 100,
                    location: null
                }
            ,

            location: {
                longitude: 0,
                latitude: 0,
                address: ""
            }
        }
    }

    scaffoldingInformation(){
        return(
            <div className={"input-fields"}>
                <div>
                    <input type={"text"}  className={"input-text-add"} onChange={event =>
                    {const scaffolding = {...this.state.scaffolding};
                        scaffolding.id = Number(event.target.value);
                        this.setState({scaffolding})}}/>
                    <p>Enter ID</p>
                </div>
                <div>
                    <span>Overfør til prosjekt:</span>
                        <Form.Select value={"Test"} onChange={(e) =>
                        {const scaffolding = {...this.state.scaffolding};
                            scaffolding.type = e.target.value;
                            this.setState({scaffolding})}}>
                            <option value={"Bunnskrue"}>Bunnskrue</option>
                            <option value={"Spire"}>Spir</option>
                            <option value={"Diagonalstang"}>Diagonalstang</option>
                            <option value={"Enrørsbjelke"}>Enrørsbjelke</option>
                            <option value={"Lengdebjeke"}>Lengdebjeke</option>
                            <option value={"Plank"}>Plank</option>
                            <option value={"Rekkverksramme"}>Rekkverksramme</option>
                            <option value={"Stillaslem"}>Stillaslem</option>
                            <option value={"Trapp"}>Trapp</option>
                        </Form.Select>
                </div>
            </div>
        )
    }

    postRequest(){
        this.state.scaffolding.location = this.state.location
        const body = [
            this.state.scaffolding
        ]

        console.log(body)
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(body)
        };
        fetch('http://localhost:8080/stillastracking/v1/api/unit', requestOptions)
            .then(response => response.json())
            .then(data => console.log("Added new Project"))
            .catch(err => console.log(err));    }


    render() {
        return(
            <div>
                {this.scaffoldingInformation()}
                <button onClick={e => this.postRequest()}>Next</button>
            </div>


    )
    }
}

export default AddScaffolding
