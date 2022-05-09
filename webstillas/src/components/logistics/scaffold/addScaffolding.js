import React from 'react'
import postModel from "../../../modelData/postModel";
import {SCAFFOLDING_URL} from "../../../modelData/constantsFile";
import "./addScaffolding.css"

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
                },
            location: {
                longitude: 0,
                latitude: 0,
                address: ""
            }
        }
    }

    scaffoldingInformation(){
        return(
            <div className={"input-information"}>
                <div className={"input-fields-add"}>
                    <p className = {"input-sorting-text"}>Enter ID</p>

                    <input type={"text"}  className={"form-control scaffolding-input"} onChange={event =>
                    {const scaffolding = {...this.state.scaffolding};
                        scaffolding.id = Number(event.target.value);
                        this.setState({scaffolding})}}/>
                </div>
                <div className={"input-fields-add"}>
                    <p className = {"input-sorting-text"}>Overfør til prosjekt:</p>
                        <select
                            className={"form-select scaffolding-input"}
                            value={"Test"}
                                onChange={(e) =>
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
                        </select>
                </div>
            </div>
        )
    }

    postRequest(){
        this.state.scaffolding.location = this.state.location
        const body = [
            this.state.scaffolding
        ]

        try {
            postModel(SCAFFOLDING_URL, JSON.stringify(body))
        }catch (e){
            console.log(e)
        }

    }


    render() {
        return(
            <div className={"main-add-scaffolding"}>
                <div className={"info-card"}>
                    {this.scaffoldingInformation()}
                    <div className={"btn-add-scaffolding"}>
                        <button className={"btn"} onClick={() => this.postRequest()}>Legg til</button>
                    </div>
                </div>
            </div>



    )
    }
}

export default AddScaffolding
