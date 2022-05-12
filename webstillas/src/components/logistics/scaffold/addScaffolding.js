import React, {useState} from 'react'
import postModel from "../../../modelData/postModel";
import {SCAFFOLDING_URL} from "../../../modelData/constantsFile";
import "./addScaffolding.css"
import {Alert} from "react-bootstrap";


/**
 * Function that will allow the user to add a new scaffolding unit
 * @returns {JSX.Element}
 */
function AddScaffolding() {

    //Defining the json body to add a new unit
    const [scaffolding, setScaffolding] = useState({
        id: "",
        type: "",
        batteryLevel: 100,
        location: {
            longitude: 0,
            latitude: 0,
            address: ""
        }
    })

    //Verification of
    const [buttonPress, setButtonPress] = useState(false)


    /**
     * Returns card to write id scaffolding type.
     * @returns {JSX.Element}
     */
    const scaffoldingInformation = () => {
        return (
            <div className={"input-information"}>
                <div className={"input-fields-add"}>
                    <p className={"input-sorting-text"}>Enter ID</p>

                    <input type={"text"} className={"form-control scaffolding-input"} onChange={event => {
                        //Setting the id
                        setScaffolding({...scaffolding, id: event.target.value})
                    }}/>

                </div>
                <div className={"input-fields-add"}>
                    <p className={"input-sorting-text"}>Stillasdel:</p>
                    <select
                        className={"form-select scaffolding-input"}
                        value={scaffolding.type}
                        onChange={(e) => {
                            //setting the type
                            setScaffolding({...scaffolding, type: e.target.value})
                        }}>
                        <option value={"Bunnskrue"}>Bunnskrue</option>
                        <option value={"Spir"}>Spir</option>
                        <option value={"Diagonalstang"}>Diagonalstang</option>
                        <option value={"Enrørsbjelke"}>Enrørsbjelke</option>
                        <option value={"Lengdebjeke"}>Lengdebjeke</option>
                        <option value={"Plank"}>Plank</option>
                        <option value={"Gelender"}>Gelender</option>
                        <option value={"Rekkverksramme"}>Rekkverksramme</option>
                        <option value={"Stillaslem"}>Stillaslem</option>
                        <option value={"Trapp"}>Trapp</option>
                    </select>
                </div>
            </div>
        )
    }


    const [type, setType] = useState("")


    /**
     * Function that will add the new scaffolding object
     *
     * @returns {Promise<void>}
     */
    const postRequest = async () => {
        setButtonPress(true)
        const body = [
            scaffolding
        ]
        try {
            await postModel(SCAFFOLDING_URL, (body))
            setType("success")

        } catch (e) {
            setType("danger")
        }
    }

    return (
        <div className={"main-add-scaffolding"}>
            {(type === "success") ?
                (<Alert className={"alert-success"}
                        key={"success"} variant={"success"}>
                    Stillasdel har blitt registrert
                </Alert>) : null}
            {(type === "danger" && buttonPress) ?
                (<Alert className={"alert-success"}
                        key={"danger"} variant={"danger"}>
                    Stillasdel har ikke blitt registrert
                </Alert>) : null}
            <div className={"info-card"}>
                {scaffoldingInformation()}
                <div className={"btn-add-scaffolding"}>
                    <button className={"btn"} onClick={() => postRequest()}>Legg til</button>
                </div>
            </div>
        </div>
    )


}

export default AddScaffolding
