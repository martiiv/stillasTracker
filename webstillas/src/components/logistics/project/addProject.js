import React from 'react'
import MapClass from "./map";
import {FormErrors} from "./FormErrors";


class AddProject extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            mapInfo: [],
            mapPage : false,
            formsErrors:
                {
                    projectName: '',
                    street: '',
                    zipcode: "",
                    municipality: "",
                    county: "",
                    name: "",
                    number: 0,
                    email: "",
                    size: 0,
                    date: ""
                },

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
            dateValid: false,

            projectID: 0,
            projectName: '',
            latitude: 60.79077759591496,
            longitude: 10.683249543160402,
            period: {
                    startDate: "",
                    endDate: ""
                },
            size: 0,
            state: "Upcoming",
            address: {
                    street: "",
                    zipcode: "",
                    municipality: "",
                    county: ""
                },
            customer: {
                    name: "",
                    number: 0,
                    email: ""
                }
        }

    }

    handleUserInput = (e) => {
        const name = e.target.name;
        const value = e.target.value;
        this.setState({[name]: value},
            () => { this.validateField(name, value) });
    }

    dateFormat(date){
        const dateArray = date.split('-')
        return dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0]
    }

    dateFormatReturn(date){
        const dateArray = date.split('-')
        return dateArray[0] + '/' + dateArray[1] + '/' + dateArray[2]
    }

    //Todo add all fylker
    countyArray = ["viken", "innlandet"]


    validateField(fieldName, value) {
        let fieldValidationErrors = this.state.formsErrors;
        let projectNameValid = this.state.projectNameValid;
        let streetValid = this.state.streetValid;
        let zipcodeValid = this.state.zipcodeValid;
        let municipalityValid = this.state.municipalityValid;
        let countyValid = this.state.countyValid;
        let sizeValid = this.state.sizeValid;
        let dateValid = this.state.dateValid;
        let nameValid = this.state.nameValid;
        let numberValid = this.state.numberValid
        let emailValid = this.state.emailValid


        switch(fieldName) {
            case 'projectName':
                projectNameValid = value.length >= 2;
                fieldValidationErrors.projectName = projectNameValid ? '': ' is too short';
                break;
            case 'street':
                streetValid = value.street.length >=2
                fieldValidationErrors.street = streetValid ? '': ' is too short';
                break;
            case 'zipcode':
                zipcodeValid = value.zipcode.length === 4
                fieldValidationErrors.zipcode = zipcodeValid  ? '': ' needs to be of length 4';
                break;
            case 'municipality':
                for (const municipalityValidElement of this.countyArray) {
                    municipalityValid = municipalityValidElement === value.municipality
                    if (municipalityValid) break
                }
                fieldValidationErrors.municipality = municipalityValid  ? '': 'Not valid';
                break;
            case 'county':
                for (const municipalityValidElement of this.countyArray) {
                    countyValid = municipalityValidElement === value.municipality
                    if (countyValid) break
                }
                fieldValidationErrors.county = countyValid  ? '': 'Not valid';
                break;
            case 'size':
                sizeValid = value >= 1
                fieldValidationErrors.size = sizeValid ? '': ' cannot be 0';
                break;
            case 'startDate':
                dateValid = value.startDate !== ""
                fieldValidationErrors.date = dateValid  ? '': 'No Valid date';
                break;
            case 'endDate':
                dateValid = value.endDate !== ""
                console.log(this.dateFormatReturn(value.endDate))
                if (Date.parse(value.endDate) < Date.parse(value.startDate)) {
                    dateValid = false
                }
                fieldValidationErrors.date = dateValid  ? '': 'No Valid date';

                break;
            case 'name':
                nameValid = value.name.length >= 2
                fieldValidationErrors.name = nameValid  ? '': 'No valid name';
                break;
            case 'number':
                numberValid =(value.number.toString().length === 8
                    && (value.number.toString().charAt(0) === "4"
                    || value.number.toString().charAt(0) === "9")
                )
                fieldValidationErrors.number = numberValid  ? '': 'Not a valid number';
                break;
            case 'email':
                const validRegex = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/;
                emailValid = value.email.match(validRegex)
                fieldValidationErrors.email = emailValid  ? '': 'Not a valid';
                break;
            default:
                break;
        }
        this.setState({formErrors: fieldValidationErrors,
                            projectNameValid: projectNameValid,
                             streetValid: streetValid,
                            zipcodeValid:zipcodeValid,
                             municipalityValid:municipalityValid,
                             countyValid:countyValid,
                             sizeValid:sizeValid,
                             dateValid :dateValid,
                             nameValid :nameValid,
                             numberValid :numberValid,
                             emailValid :emailValid
        }, this.validateForm);
    }

    validateForm() {
        this.setState({formValid:
                this.state.projectNameValid
                && this.state.streetValid
                && this.state.zipcodeValid
            && this.state.municipalityValid
            && this.state.countyValid
            && this.state.nameValid
            && this.state.numberValid
            && this.state.emailValid
            && this.state.sizeValid
        });
    }

    errorClass(error) {
        console.log("error " + error)
        return(error.length === 0 ? '' : 'has-error');
    }

    generalInformation(){
        //todo integrere med api, slik at brukeren ikke trenger å skrive inn hele addressen.
        return(
        <div className={"general-information"}>
                <h1>Add project</h1>
                <hr/>
            <FormErrors formErrors={this.state.formsErrors} />

            <form>
                    <div className={"input-fields"}>
                    <div className={`form-group ${this.errorClass(this.state.formsErrors.projectName)}`}>
                        <input type={"text"}
                               required
                               name={"projectName"}
                               placeholder={"Project Name"}
                               className={"input-text-add"}
                               value={this.state.projectName}
                               onChange={this.handleUserInput}

                        />
                        <p>Enter Project Name</p>
                    </div>
                    <div>
                        <input type={"text"}
                               className={"input-text-add"}
                               onChange={(event) => {
                                   const address = {...this.state.address};
                                   address.street = event.target.value;
                                   this.setState({address})
                                   this.validateField("street", address)}
                                }

                        />
                        <p>Enter Address</p>
                    </div>
                    <div>
                        <input type={"number"}
                               className={"input-text-add"}
                               onChange={event => {
                                   const address = {...this.state.address};
                                   address.zipcode = (event.target.value);
                                   this.setState({address})
                                   this.validateField("zipcode", address)}
                               }
                        />
                        <p>Enter Zip Code</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"}
                               onChange={event => {
                                   const address = {...this.state.address};
                                   address.municipality = event.target.value;
                                   this.setState({address})
                                   this.validateField("municipality", address)}
                               }/>
                        <p>Enter Municipality</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"}
                               onChange={event => {
                                   const address = {...this.state.address};
                                   address.county = event.target.value;
                                   this.setState({address})
                                   this.validateField("county", address)}
                               }/>
                        <p>Enter County</p>
                    </div>
                    <div>
                        <input type={"number"}
                               className={"input-number-add"}
                               name={"size"}
                               value={this.state.size}
                               onChange={this.handleUserInput}/>
                        <p>Enter size</p>
                    </div>
                    <div>
                        <input type={"date"}  className={"input-date-add"}
                               onChange={event => {
                                const period = {...this.state.period};

                                   period.startDate = this.dateFormat(event.target.value);
                                   this.setState({period})
                                this.validateField("startDate", period)
                               }
                        }/>
                        <p>Enter Start date</p>
                    </div>
                    <div>
                        <input type={"date"}
                               className={"input-date-add"}
                               onChange={event => {
                                   const period = {...this.state.period};
                                   period.endDate = this.dateFormat(event.target.value);
                                   this.setState({period})
                                   this.validateField("endDate", period)
                               }
                        }/>
                        <p>Enter end Date</p>
                    </div>
                    <div>
                        <input type={"text"}
                               className={"input-text-add"}
                               onChange={event =>
                                 {const customer = {...this.state.customer};
                                customer.name = event.target.value;
                                this.setState({customer})
                               this.validateField("name", customer)}}

                        />
                        <p>Enter Customer Name</p>
                    </div>
                    <div>
                        <input type={"number"}  className={"input-text-add"} onChange={event =>
                        {const customer = {...this.state.customer};
                            customer.number = Number(event.target.value);
                            this.setState({customer})
                            this.validateField("number", customer)}}
                        />
                        <p>Enter Customer Number</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"} onChange={event =>
                        {const customer = {...this.state.customer};
                            customer.email = event.target.value;
                            this.setState({customer})
                            this.validateField("email", customer)
                        }}
                        />
                        <p>Enter Customer Email</p>
                    </div>
                </div>
                <button type={"submit"} disabled={!this.state.formValid} onClick={() => this.setState({mapPage: true})}>Next</button>
                </form>
            </div>
        )

    }







    render() {
        //todo check projectID number
        const project = ({
                projectID: Math.round(Math.random() * 1000),
                projectName: this.state.projectName,
                latitude: 60.79077759591496,
                longitude: 10.683249543160402,
                period: {
                    startDate: this.state.period.startDate,
                    endDate: this.state.period.endDate
                },
                size: Number(this.state.size),
                state: "Upcoming",
                address: {
                    street: this.state.address.street,
                    zipcode: this.state.address.zipcode,
                    municipality: this.state.address.municipality,
                    county: this.state.address.county
                },
                customer: {
                    name: this.state.customer.name,
                    number: this.state.customer.number,
                    email: this.state.customer.email
                }
            }
        )

        const {mapPage} = this.state;
        if (!mapPage){
            return(
                this.generalInformation()
            )
        }else {
            return (
                <MapClass project = {(project)}/>
            )
        }
    }
}

export default AddProject
