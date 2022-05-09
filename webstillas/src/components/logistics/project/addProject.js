import React, {useCallback, useRef, useState} from 'react'
import {MapClass} from "./map";
import MapboxAutocomplete from "react-mapbox-autocomplete";
import 'mapbox-gl/dist/mapbox-gl.css'
import "./addProject.css"
import {Alert} from "react-bootstrap";

export default function AddProjectFunc() {

    const mapAccess = {
        // Thanks to SomeSoftwareTeam (https://github.com/SomeSoftwareTeam/some-react-app/blob/acd17860b8b1f51edefa4e18486cc1fb07afff70/src/components/SomeComponent.js)
        mapboxApiAccessToken:
            "pk.eyJ1IjoiZmFrZXVzZXJnaXRodWIiLCJhIjoiY2pwOGlneGI4MDNnaDN1c2J0eW5zb2ZiNyJ9.mALv0tCpbYUPtzT7YysA2g"
    };


    const [address, setAddress] = useState({street: "", zipcode: 0, municipality: "", county: ""})


    const [period, setPeriod] = useState({startDate: "", endDate: ""})

    const [customer, setCustomer] = useState({name: "", email: ""})

    const [customerNumber, setCustomerNumber] = useState({number: 0})

    const [projectDetails, setProjectDetails] = useState({
        projectID: Math.round(Math.random() * 1000000),
        projectName: '',
        latitude: 0,
        longitude: 0,
        state: "Inactive"
    })

    const [size, setSize] = useState({size: 0})

    const [errors, setErrors] = useState({
        projectName: '',
        address: "",
        name: "",
        number: "",
        email: "",
        size: "",
        date: ""
    })

    const [valid, setValid] = useState({
        projectNameValid: false,
        streetValid: false,
        zipcodeValid: false,
        municipalityValid: false,
        countyValid: false,
        nameValid: false,
        numberValid: false,
        emailValid: false,
        sizeValid: false,
        dateValid: false
    })


    const handleUserInputProjectDetails = (e) => {
        const name = e.target.name;
        const value = e.target.value;
        if (name.toLowerCase() === "size") {
            const valueNumber = Number(value);
            setSize({...size, [name]: valueNumber});
            validateFieldProjectDetails(name, valueNumber)
        }
        validateFieldProjectDetails(name, value)
        setProjectDetails({...projectDetails, [name]: value});
    }

    const handleUserInputCustomer = (e) => {
        const name = e.target.name;
        const value = e.target.value;
        if (name.toLowerCase() === "number") {
            setCustomerNumber({...customerNumber, [name]: parseInt(value, 10)});
            validateFieldProjectCustomer(name, (value))
        }
        validateFieldProjectCustomer(name, value)
        setCustomer(({...customer, [name]: value}));
    }

    const handleUserInputAddress = (e) => {
        const name = e.target.name;
        const value = e.target.value;
        validateFieldProjectAddress(name, value)
        setAddress({...address, [name]: value});
    }


    const dateFormat = (date) => {
        const dateArray = date.split('-')
        return dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0]
    }

    const validateFieldProjectAddress = (field, value) => {
        switch (field) {
            case "street":
                setValid({...valid, streetValid: value.length > 3})
                setErrors({...errors, street: (valid.streetValid ? '' : ' is too short')})
                break;
            case "zipcode":
                setValid({...valid, zipcodeValid: ((value.length) === 4)})
                setErrors({...errors, zipcode: (valid.zipcodeValid ? '' : ' needs to be of length 4')});
                break;
            case 'municipality':
                setValid({...valid, municipalityValid: (value !== undefined)})
                setErrors({...errors, municipalityErr: (valid.municipalityValid ? '' : 'Not a valid municipality')});

                break;
            case 'county':
                setValid({...valid, countyValid: (value !== undefined)})
                setErrors({...errors, countyErr: (valid.countyValid ? '' : 'Not a valid county')});
                break;
            default:
                break;
        }
    }

    const validateFieldProjectCustomer = (fieldName, value) => {
        switch (fieldName) {
            case 'name':
                setValid({...valid, nameValid: (value.length >= 2)})
                setErrors({...errors, name: (valid.nameValid ? '' : 'No valid name')})
                break;
            case 'number':
                setValid({
                    ...valid, numberValid: (value.toString().length === 8
                        && (value.toString().charAt(0) === "4" || value.toString().charAt(0) === "9"))
                })
                setErrors({...errors, number: (valid.numberValid ? '' : 'Not a valid number')})
                break;
            case 'email':
                const validRegex = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/;
                setValid({...valid, emailValid: validRegex.test(value)})
                setErrors({...errors, email: (valid.emailValid ? '' : 'Not a valid')})
                break;
            default:
                break;
        }
    }

    const validateFieldDate = (fieldName, value) => {
        switch (fieldName) {
            case 'startDate':
                setValid({...valid, dateValid: (value.startDate !== "")})
                setErrors({...errors, date: (valid.dateValid ? '' : 'No Valid date')})
                break;
            case 'endDate':
                setValid({...valid, dateValid: (value.endDate !== "")})
                setErrors({...errors, date: (valid.dateValid ? '' : 'No Valid date')})
                break;
            default:
                break;
        }
    }

    const validateFieldProjectDetails = (fieldName, value) => {
        switch (fieldName) {
            case 'projectName':
                setValid({...valid, projectNameValid: (value.length >= 2)})
                setErrors({...errors, projectName: (valid.projectNameValid ? '' : ' is too short')})
                break;
            case 'size':
                setValid({...valid, sizeValid: (Number(value) > 0)})
                setErrors({...errors, size: (valid.sizeValid ? '' : ' cannot be 0')})
                break;
            default:
                break;
        }
    }


    const handleUserInputPeriod = (e) => {
        const name = e.target.name;
        const value = e.target.value;
        const date = dateFormat(value)
        validateFieldDate(name, value)
        setPeriod({...period, [name]: date});
    }


    const [mapPage, setMapPage] = useState(false)

    const nextPage = () => {
        setMapPage(true)

    }

    const queryParams = {
        country: "no",
        place_type: "address"
    };

    const parseReverseGeo = async (geoData, lat, long) => {
        let street, postcode, region, place
        await fetch("https://api.mapbox.com/geocoding/v5/mapbox.places/" + long + "," + lat + ".json?access_token=pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ")
            .then(res => res.json())
            .then(res => {

                    let validStreet, validZip, validCounty, validMunicipality
                    for (const re of res.features) {
                        console.log((re.place_type[0]))
                        switch (re.place_type[0]) {

                            case "address": {
                                street = re.text
                                if ((re.text.length > 3)) {
                                    validStreet = true
                                }
                            }
                                break;
                            case "poi": {
                                street = re.text
                                if ((re.text.length > 3)) {
                                    validStreet = true
                                }

                            }
                                break;
                            case "postcode": {
                                postcode = re.text
                                if ((re.text.length === 4)) {
                                    validZip = true
                                }

                            }
                                break;
                            case ("region"): {
                                region = re.text
                                if ((re.text.length !== undefined)) {
                                    validCounty = true
                                }
                            }
                                break;
                            case ("place") : {
                                place = re.text
                                if ((re.text.length !== undefined)) {
                                    validMunicipality = true
                                }

                            }
                                break;
                            default:
                                console.log("Ikke validert")
                        }

                    }

                    if (validStreet && validZip && validCounty && validMunicipality) {
                        setValid({
                            ...valid,
                            countyValid: validCounty,
                            municipalityValid: validMunicipality,
                            zipcodeValid: validZip,
                            streetValid: validStreet
                        })
                        if (region === "Oslo") {
                            setAddress({
                                ...address,
                                street: street,
                                county: region,
                                municipality: region,
                                zipcode: postcode
                            })
                        } else {
                            setAddress({...address, street: street, county: region, municipality: place, zipcode: postcode})
                        }
                    } else {
                        setErrors({...errors, address: "You have entered an invalid address"})
                    }


                }
            ).then(() => setProjectDetails({
                ...projectDetails,
                longitude: long,
                latitude: lat
            }))
    }

    const _suggestionSelect = async (result, lat, long, country) => {
        await parseReverseGeo(result, lat, long)
    }

    const finalProject = {
        projectID: projectDetails.projectID,
        projectName: projectDetails.projectName,
        latitude: Number(projectDetails.latitude),
        longitude: Number(projectDetails.longitude),
        state: projectDetails.state,
        size: (size.size),
        period: {
            startDate: period.startDate,
            endDate: period.endDate
        },
        customer: {
            name: customer.name,
            number: customerNumber.number,
            email: customer.email,
        },
        address: {
            street: address.street,
            zipcode: String(address.zipcode),
            municipality: address.municipality,
            county: address.county,
        }

    }


    let formsValid = false
    if (valid.countyValid && valid.streetValid && valid.zipcodeValid && valid.municipalityValid
        && valid.dateValid && valid.sizeValid && valid.projectNameValid && valid.emailValid && valid.nameValid
        && valid.numberValid
    ) {
        formsValid = true
    }



console.log(address)

    const contactInformation = () => {

        return(
            <div>
                <h3>Contact Information</h3>
                <hr/>
                <div className={"input-with-text"}>
                    <p className={"input-field-text"}>Name</p>
                    <input
                        className = {"form-control"}
                        type={"text"}
                        required
                        name={"name"}
                        placeholder={"Enter Customer Name"}
                        onChange={handleUserInputCustomer}
                    />
                    <p className={"error-message"}>
                        {(errors.name === "" || valid.nameValid) ?  null: <Alert variant="danger">{errors.name}</Alert> }

                    </p>
                </div>
                <div className={"input-with-text"}>
                    <p className={"input-field-text"}>Number</p>
                    <input
                        className = {"form-control"}

                        type={"number"}
                           min={0}
                           required
                           name={"number"}
                           placeholder={"Enter Customer Number"}
                           onChange={handleUserInputCustomer}
                    />
                    <p className={"error-message"}>
                        {(errors.number === "" || valid.numberValid) ?  null: <Alert variant="danger">{errors.number}</Alert> }
                    </p>

                </div>
                <div className={"input-with-email"}>
                    <p className={"input-field-text"}>Email</p>
                    <input
                        className = {"form-control"}
                        type={"email"}
                        required
                        name={"email"}
                        placeholder={"Enter Customer Email"}
                        onChange={handleUserInputCustomer}
                    />
                    <p className={"error-message"}>
                        {(errors.email === "" || valid.emailValid) ?  null: <Alert variant="danger">{errors.email}</Alert> }
                    </p>

                </div>
            </div>
        )
    }






    console.log(valid)


    if (!mapPage) {
        return (
            <div className={"add-card"}>
                <article className={"information"}>
                <h1>Add project</h1>
                <h2>Generelt</h2>
                <hr/>
                <div>
                    <div className={"test"}>
                        <div className={"address-name"}>
                                <div className={"input-with-text"}>
                                    <p className={"input-field-text"}>Project Name </p>
                                    <input
                                        className = {"form-control name"}
                                        type={"text"}
                                        required
                                        name={"projectName"}
                                        placeholder={"Project Name"}
                                        onChange={handleUserInputProjectDetails}
                                    />
                                    <p className={"error-message"}>
                                        {(errors.projectName === "" || valid.projectNameValid) ?  null: <Alert variant="danger">{errors.projectName}</Alert> }
                                    </p>

                                </div>

                            <div className={"input-with-text"}>
                                <p className={"input-field-text"}>Project size</p>
                                <input type={"number"}
                                       min={0}
                                       required
                                       name={"size"}
                                       placeholder={"Size"}
                                       className = {"form-control number"}
                                       onChange={handleUserInputProjectDetails}/>
                                <p className={"error-message"}>
                                    {(errors.size === "" || valid.sizeValid) ?  null: <Alert variant="danger">{errors.size}</Alert> }
                                </p>

                            </div>


                        </div>
                        <div className={"date-add-project"}>
                            <div className="row">
                                <div className="col">
                                    <p className={"input-field-text"}>Start date</p>
                                    <input type={"date"}
                                           required
                                           name={"startDate"}
                                           placeholder={"Start Date"}
                                           className={"input-text-add"}
                                           onChange={handleUserInputPeriod}/>
                                    <p className={"error-message"}>
                                        {(errors.date === "" || valid.dateValid) ?  null: <Alert variant="danger">{errors.date}</Alert> }

                                    </p>

                                </div>
                                <div className="col">
                                    <p className={"input-field-text"}>End date</p>
                                    <input type={"date"}
                                           required
                                           name={"endDate"}
                                           placeholder={"End Date"}
                                           className={"input-text-add"}
                                           onChange={handleUserInputPeriod}/>
                                    <p className={"error-message"}>
                                        {(errors.date === "" || valid.dateValid) ?  null: <Alert variant="danger">{errors.date}</Alert> }
                                    </p>

                                </div>
                            </div>
                        </div>
                        <div className="col">
                            <div className={"input-with-text"}>
                                <p className={"input-field-text"}>Address</p>
                                <MapboxAutocomplete
                                    inputClass='form-control address'
                                    publicKey={mapAccess.mapboxApiAccessToken}
                                    onSuggestionSelect={_suggestionSelect}

                                    country="no"
                                    resetSearch={false}
                                    placeholder="Search Address..."
                                    queryParams={queryParams}
                                />

                                <p className={"error-message"}>
                                    {(errors.address === "" || valid.streetValid) ?  null: <Alert variant="danger">{errors.address}</Alert> }
                                </p>
                            </div>
                        </div>
                        {contactInformation()}

                    </div>
                </div>
                </article>
                <MapClass props={finalProject}
                          valid ={formsValid}

                />
            </div>
        )
    }
}





