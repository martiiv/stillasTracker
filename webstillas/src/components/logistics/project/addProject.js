import React, {useEffect, useReducer, useState} from 'react'
import {MapClass} from "./map";
import {FormErrors} from "./FormErrors";



export default function AddProjectFunc(string) {


    const [address, setAddress] = useState({street: "", municipality: "", county: ""})

    const [zipCode, setZipCode] = useState({zipcode: 0})

    const [period, setPeriod] = useState({startDate: "", endDate: ""})

    const [customer, setCustomer] = useState({name: "",email: ""})

    const [customerNumber, setCustomerNumber] = useState({number: 0})

    const [projectDetails, setProjectDetails] = useState({
        projectID: 2122398,
        projectName: '',
        latitude: 60.79077759591496,
        longitude: 10.683249543160402,
        state: "Active"
    })

    const [size, setSize] = useState({size: 0})

    const [errors, setErrors] = useState({
        projectName: '',
        street: '',
        zipcode: "",
        municipality: "",
        county: "",
        name: "",
        number: 0,
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
        formValid: false,
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
            setCustomerNumber({...customerNumber,  [name]: parseInt(value, 10)});
            validateFieldProjectCustomer(name, (value))
        }
        validateFieldProjectCustomer(name, value)
        setCustomer(({...customer, [name]: value}));
    }

    const handleUserInputAddress = (e) => {
        const name = e.target.name;
        const value = e.target.value;
        if (name.toLowerCase() === "zipcode"){
            setZipCode({...zipCode, [name]: parseInt(value, 10)})
            validateFieldProjectAddress(name, parseInt(value, 10))
        }
        validateFieldProjectAddress(name, value)
        setAddress({...address, [name]: value});
    }


    const dateFormat = (date) => {
        const dateArray = date.split('-')
        return dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0]
    }

    const validateFieldProjectAddress = (fieldName, value) => {
        switch (fieldName) {
            case 'street':
                setValid({...valid, streetValid: (value.length >= 2)})
                setErrors({...errors, street: (valid.streetValid ? '' : ' is too short')})
                break;
            case 'zipcode':
                setValid({...valid, zipcodeValid: (value.length === 4)})
                setErrors({...errors, zipcode: (valid.zipcodeValid ? '' : ' needs to be of length 4')});
                break;
            case 'municipality':
                //Todo validate
                setValid({...valid, municipalityValid: true})
                break;
            case 'county':
                //Todo validate
                setValid({...valid, countyValid: true})
                break;
            default:
                break;
        }
    }

    const validateFieldProjectCustomer = (fieldName, value) => {
        switch (fieldName) {
            case 'name':
                setValid({...valid, nameValid: (value.length >= 2)})
                setErrors({...errors, street: (valid.nameValid ? '' : 'No valid name')})
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

                setValid({...valid, sizeValid: (value >= 1)})
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

    const finalProject = {
        projectID: projectDetails.projectID,
        projectName: projectDetails.projectName,
        latitude: projectDetails.latitude,
        longitude: projectDetails.longitude,
        state: projectDetails.state,
        size: (size.size),
        period: {
            startDate: period.startDate,
            endDate: period.endDate
        },
        customer:{
            name: customer.name,
            number: customerNumber.number,
            email: customer.email,
        },
        address:{
            street: address.street,
            zipcode: String(zipCode.zipcode),
            municipality: address.municipality,
            county: address.county,
        }

    }

    console.log(finalProject)
    console.log(valid)

    if (!mapPage) {
        return (
            <div className={"general-information"}>
                <h1>Add project</h1>
                <hr/>
                <form>
                    <div className={"input-fields"}>
                        <div className={`form-group`}>
                            <input type={"text"}
                                   required
                                   name={"projectName"}
                                   placeholder={"Project Name"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputProjectDetails}

                            />
                            <p>Enter Project Name</p>
                        </div>
                        <div>
                            <input type={"text"}
                                   required
                                   name={"street"}
                                   placeholder={"Street"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputAddress}
                            />
                            <p>Enter Address</p>
                        </div>
                        <div>
                            <input type={"number"}
                                   min={0}
                                   required
                                   name={"zipcode"}
                                   placeholder={"ZIP Code"}
                                   onChange={handleUserInputAddress}
                            />
                            <p>Enter Zip Code</p>
                        </div>
                        <div>
                            <input type={"text"}
                                   required
                                   name={"municipality"}
                                   placeholder={"Enter Municipality"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputAddress}
                            />
                            <p>Enter Municipality</p>
                        </div>
                        <div>
                            <input type={"text"}
                                   required
                                   name={"county"}
                                   placeholder={"Enter County"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputAddress}
                            />
                            <p>Enter County</p>
                        </div>
                        <div>
                            <input type={"number"}
                                   min={0}
                                   required
                                   name={"size"}
                                   placeholder={"Size"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputProjectDetails}/>
                            <p>Enter size</p>
                        </div>
                        <div>
                            <input type={"date"}
                                   required
                                   name={"startDate"}
                                   placeholder={"Start Date"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputPeriod}/>
                            <p>Enter Start date</p>
                        </div>
                        <div>
                            <input type={"date"}
                                   required
                                   name={"endDate"}
                                   placeholder={"End Date"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputPeriod}/>
                            <p>Enter end Date</p>
                        </div>
                        <div>
                            <input type={"text"}
                                   required
                                   name={"name"}
                                   placeholder={"Project Name"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputCustomer}
                            />
                            <p>Enter Customer Name</p>
                        </div>
                        <div>
                            <input type={"number"}
                                   min={0}
                                   required
                                   name={"number"}
                                   placeholder={"Project Name"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputCustomer}
                            />
                            <p>Enter Customer Number</p>
                        </div>
                        <div>
                            <input type={"text"}
                                   required
                                   name={"email"}
                                   placeholder={"Project Name"}
                                   className={"input-text-add"}
                                   onChange={handleUserInputCustomer}
                            />
                            <p>Enter Customer Email</p>
                        </div>
                    </div>
                </form>
                <button type={"submit"} onClick={() => nextPage()}>Next</button>
            </div>
        )
    } else {
        return <MapClass props = {finalProject}/>
    }
}





