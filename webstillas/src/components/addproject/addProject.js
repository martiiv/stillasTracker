import React, {useState} from 'react'
import {MapClass} from "../../Pages/addProjectMap";
import MapboxAutocomplete from "react-mapbox-autocomplete";
import 'mapbox-gl/dist/mapbox-gl.css'
import "../../Assets/Styles/addProject.css"
import {Alert} from "react-bootstrap";
import {MapBoxAPIKey} from "../../Config/firebaseConfig";


/**
 Function that will allow the user to add a new project to the system.
 */
export default function AddProjectFunc() {
    //Access token to the mapbox api.
    const mapAccess = {
        // Thanks to SomeSoftwareTeam (https://github.com/SomeSoftwareTeam/some-react-app/blob/acd17860b8b1f51edefa4e18486cc1fb07afff70/src/components/SomeComponent.js)
        mapboxApiAccessToken: MapBoxAPIKey
    };

    /*
   Initialise variables the user must fill in order to add a new project.
     */
    const [address, setAddress] = useState({street: "", zipcode: 0, municipality: "", county: ""})
    const [period, setPeriod] = useState({startDate: "", endDate: ""})
    const [customer, setCustomer] = useState({name: "", email: ""})
    const [customerNumber, setCustomerNumber] = useState({number: 0})
    const [projectDetails, setProjectDetails] = useState({
        projectID: Math.round(Math.random() * 1000000),
        projectName: '',
        latitude: 0,
        longitude: 0,
        state: ""
    })
    const [size, setSize] = useState({size: 0})

    /*
    Error messages that will be displayed to the user, if the user has entered invalid input.
     */
    const [errors, setErrors] = useState({
        projectName: '',
        address: "",
        name: "",
        number: "",
        email: "",
        size: "",
        date: ""
    })

    /*
    Checks if the input is valid.
     */
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


    /**
     * Function that will handle and check the user input
     *
     * @param e the input field name and value.
     */
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

    /**
     * Function that will handle and check the user input
     *
     * @param e the input field name and value.
     */
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

    /**
     * Function that will format input from "input date" to required API format
     *
     * @param date as "input date" format "mm-dd-yyyy"
     * @returns {string} return to API format "dd-mm-yyyy"
     */
    const dateFormat = (date) => {
        const dateArray = date.split('-')
        return dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0]
    }

    /**
     * Function that will validate the customer input
     * If input is not valid, then a predefined Indicators message is set.
     *
     * @param fieldName is the object field that is going to be set
     * @param value is the object value to be set.
     */
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

    /**
     * Function to validate date format.
     * If input is not valid, then a predefined Indicators message is set.
     *
     * @param fieldName is the object field that is going to be set
     * @param value is the object value to be set.
     */
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

    /**
     * Function to validate projectDetails
     * If input is not valid, then a predefined Indicators message is set.
     *
     * @param fieldName is the object field that is going to be set
     * @param value is the object value to be set.
     */
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


    /**
     * Function that will set period fields
     *
     * @param e is value and name of input fields.
     */
    const handleUserInputPeriod = (e) => {
        const name = e.target.name;
        const value = e.target.value;
        const date = dateFormat(value)
        validateFieldDate(name, value)
        setPeriod({...period, [name]: date});
    }



    // Query that will only allow norwegian addresses
    const queryParams = {
        country: "no",
        place_type: "address"
    };

    /**
     * Function that will fetch data of a spesific longitude and latitude, to set address/poi, postcode, county and municipality
     *
     * @param lat latitude of the place we would like to get information
     * @param long longitude of the place we would like to get information
     * @returns {Promise<void>}
     */
    const parseReverseGeo = async (lat, long) => {
        let street, postcode, region, place
        try {
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

                            //If region is oslo the municipality is not set.
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
        }catch (e) {
            console.log(e)

        }

    }

    /**
     *Function that will redirect to parseReverseGeo
     *
     * @param lat latitude from the suggested input of user
     * @param long longitude from the suggested input of user
     * @returns {Promise<void>}
     */
    const _suggestionSelect = async (_, lat, long) => {
        await parseReverseGeo(lat, long)
    }


    /**
     * Struct of adding project that will set the variables the user has set.
     *
     */
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


    /**
     * Checks if all the fields are valid.
     *
     * @type {boolean} is set true if all the fields are valid.
     */
    let formsValid = false
    if (valid.countyValid && valid.streetValid && valid.zipcodeValid && valid.municipalityValid
        && valid.dateValid && valid.sizeValid && valid.projectNameValid && valid.emailValid && valid.nameValid
        && valid.numberValid
    ) {
        formsValid = true
    }


    /**
     *
     * @returns {JSX.Element} with input that can set the customer information.
     */
    const contactInformation = () => {
        return (
            <div>
                <h3>Contact Information</h3>
                <hr/>
                <div className={"input-with-text"}>
                    <p className={"input-field-text"}>Name</p>
                    <input
                        className={"form-control"}
                        type={"text"}
                        required
                        name={"name"}
                        placeholder={"Enter Customer Name"}
                        onChange={handleUserInputCustomer}
                    />
                    <p className={"error-message"}>
                        {(errors.name === "" || valid.nameValid) ? null : <Alert variant="danger">{errors.name}</Alert>}

                    </p>
                </div>
                <div className={"input-with-text"}>
                    <p className={"input-field-text"}>Number</p>
                    <input
                        className={"form-control"}

                        type={"number"}
                        min={0}
                        required
                        name={"number"}
                        placeholder={"Enter Customer Number"}
                        onChange={handleUserInputCustomer}
                    />
                    <p className={"error-message"}>
                        {(errors.number === "" || valid.numberValid) ? null :
                            <Alert variant="danger">{errors.number}</Alert>}
                    </p>

                </div>
                <div className={"input-with-email"}>
                    <p className={"input-field-text"}>Email</p>
                    <input
                        className={"form-control"}
                        type={"email"}
                        required
                        name={"email"}
                        placeholder={"Enter Customer Email"}
                        onChange={handleUserInputCustomer}
                    />
                    <p className={"error-message"}>
                        {(errors.email === "" || valid.emailValid) ? null :
                            <Alert variant="danger">{errors.email}</Alert>}
                    </p>

                </div>
            </div>
        )
    }


    return (
        <div className={"add-card"}>
            <article className={"information"}>
                <h1>Legg til prosjekt</h1>
                <h2>Generelt</h2>
                <hr/>
                <div>
                    <div className={"test"}>
                        <div className={"address-name"}>
                            <div className={"input-with-text"}>
                                <p className={"input-field-text"}>Project Name </p>
                                <input
                                    className={"form-control name"}
                                    type={"text"}
                                    required
                                    name={"projectName"}
                                    placeholder={"Project Name"}
                                    onChange={handleUserInputProjectDetails}
                                />
                                <p className={"error-message"}>
                                    {(errors.projectName === "" || valid.projectNameValid) ? null :
                                        <Alert variant="danger">{errors.projectName}</Alert>}
                                </p>

                            </div>

                            <div className={"input-with-text"}>
                                <p className={"input-field-text"}>Project size</p>
                                <input type={"number"}
                                       min={0}
                                       required
                                       name={"size"}
                                       placeholder={"Size"}
                                       className={"form-control number"}
                                       onChange={handleUserInputProjectDetails}/>
                                <p className={"error-message"}>
                                    {(errors.size === "" || valid.sizeValid) ? null :
                                        <Alert variant="danger">{errors.size}</Alert>}
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
                                        {(errors.date === "" || valid.dateValid) ? null :
                                            <Alert variant="danger">{errors.date}</Alert>}

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
                                        {(errors.date === "" || valid.dateValid) ? null :
                                            <Alert variant="danger">{errors.date}</Alert>}
                                    </p>

                                </div>


                                <div className="col">
                                    <p className={"input-field-text"}>Prosjekt status</p>
                                    <select className={"form-select options"} onChange={(e) =>
                                        setProjectDetails({...projectDetails, state: e.target.value})}>
                                        <option defaultValue="">Velg her</option>
                                        <option value={"Active"}>Aktiv</option>
                                        <option value={"Inactive"}>Inaktiv</option>
                                        <option value={"Upcoming"}>Kommende</option>
                                    </select>
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
                                    resetSearch={false}
                                    country = "no"
                                    placeholder="Search Address..."
                                    queryParams={queryParams}
                                />

                                <p className={"error-message"}>
                                    {(errors.address === "" || valid.streetValid) ? null :
                                        <Alert variant="danger">{errors.address}</Alert>}
                                </p>
                            </div>
                        </div>
                        {contactInformation()}

                    </div>
                </div>
            </article>
            <MapClass props={finalProject}
                      valid={formsValid}

            />
        </div>
    )

}





