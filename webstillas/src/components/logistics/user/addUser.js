import React from "react";
import {FormErrors} from "../project/FormErrors";
import postModel from "../../../modelData/postModel";
import {USER_URL} from "../../../modelData/constantsFile";

class AddUser extends React.Component{
    constructor(props) {
        super(props);
        this.state = {

            formsErrors:
                {
                    firstName: "",
                    lastName: "",
                    number: 0,
                    email: "",
                    date: "",
                    role: "",
                    admin: ""
                },


            firstNameValid: false,
            lastNameValid: false,
            numberValid: false,
            emailValid: false,
            dateValid: false,
            roleValid: false,
            adminValid: false,




            employee : {
                employeeID: Math.round(Math.random() * 1000),
                dateOfBirth: null,
                role: "",
                phone: 0,
                email: "",
                admin: false,
                name: null
            },

            name: {
                firstName: "",
                lastName: ""
            },

        }
    }

    dateFormat(date){
        const dateArray = date.split('-')
        return dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0]
    }


    addPostRequest(){
        this.state.employee.name = this.state.name
        const result = postModel(USER_URL, JSON.stringify(this.state.employee))
        console.log(result)
    }



    validateField(fieldName, value) {
        let fieldValidationErrors = this.state.formsErrors;
        let dateValid = this.state.dateValid;
        let firstNameValid = this.state.firstNameValid;
        let lastNameValid = this.state.lastNameValid;

        let numberValid = this.state.numberValid
        let emailValid = this.state.emailValid
        let roleValid = this.state.roleValid
        let adminValid = this.state.adminValid


        switch(fieldName) {
            case 'startDate':
                dateValid = value.startDate !== ""
                fieldValidationErrors.date = dateValid  ? '': 'No Valid date';
                break;
            case 'firstName':
                firstNameValid = value.firstName.length >= 2
                fieldValidationErrors.firstName = firstNameValid  ? '': 'No valid name';
                break;
            case 'lastName':
                lastNameValid = value.lastName.length >= 2
                fieldValidationErrors.lastName = lastNameValid  ? '': 'No valid name';
                break;
            case 'number':
                numberValid =(value.phone.toString().length === 8
                    && (value.phone.toString().charAt(0) === "4"
                        || value.phone.toString().charAt(0) === "9")
                )
                fieldValidationErrors.number = numberValid  ? '': 'Not a valid number';
                break;
            case 'email':
                const validRegex = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/;
                emailValid = value.email.match(validRegex)
                fieldValidationErrors.email = emailValid  ? '': 'Not a valid';
                break;
            case 'date':
                dateValid = value.dateOfBirth !== ""
                fieldValidationErrors.date = dateValid  ? '': 'No Valid date';
                break;
            case 'role':
                roleValid = value.role !== "Choose here";
                console.log(value.role)
                fieldValidationErrors.role = roleValid  ? '': 'Choose role';
                break;
            case 'admin':
                adminValid = value.admin !== "Choose here";
                console.log(adminValid)
                fieldValidationErrors.admin = adminValid  ? '': 'Choose role';
                break;
            default:
                break;
        }
        this.setState({formErrors: fieldValidationErrors,
            dateValid :dateValid,
            firstNameValid :firstNameValid,
            lastNameValid :lastNameValid,
            numberValid :numberValid,
            emailValid :emailValid,
            roleValid :roleValid,
            adminValid :adminValid


        }, this.validateForm);
    }


    errorClass(error) {
        console.log("error " + error)
        return(error.length === 0 ? '' : 'has-error');
    }

    validateForm() {
        this.setState({formValid:
                 this.state.firstNameValid
                && this.state.lastNameValid
                && this.state.numberValid
                && this.state.emailValid
                && this.state.dateValid
                && this.state.roleValid
                && this.state.adminValid
        });
    }

    inputInformation(){
        //todo bytt fra forms.select til kun select på alle select i appen.

        console.log(this.state)
        return(

            <div className={"input-fields"}>
                <FormErrors formErrors={this.state.formsErrors} />

                <div>
                    <input type={"text"}  className={"input-text-add"} onChange={event =>
                    {
                        const name = {...this.state.name};
                        name.firstName = event.target.value;
                        this.setState({name})
                        this.validateField("firstName", name)}}
                    />
                    <p>Enter First name</p>
                </div>

                <div>
                    <input type={"text"}  className={"input-text-add"}
                           onChange={event =>{
                        const name = {...this.state.name};
                        name.lastName = event.target.value;
                        this.setState({name})
                        this.validateField("lastName", name)}}
                    />
                    <p>Enter Last name</p>
                </div>
                <div>
                    <input type={"date"}  className={"input-date-add"}
                           onChange={event => {
                               const employee = {...this.state.employee};
                             employee.dateOfBirth = this.dateFormat(event.target.value);
                            this.setState({employee})
                               this.validateField("date", employee)
                           }}/>
                    <p>Enter Birthdate</p>
                </div>
                <div>
                    <input type={"number"}  className={"input-text-add"} onChange={event =>
                    {const employee = {...this.state.employee};
                        employee.phone = Number(event.target.value);
                        this.setState({employee})
                        this.validateField("number", employee)
                    }}
                    />
                    <p>Enter User phone number</p>
                </div>
                <div>
                    <input type={"text"}  className={"input-text-add"} onChange={event =>
                    {const employee = {...this.state.employee};
                        employee.email = event.target.value;
                        this.setState({employee})
                        this.validateField("email", employee)
                    }}/>
                    <p>Enter User Email</p>
                </div>
                <div>
                    <select onChange={(e) =>
                    {const employee = {...this.state.employee};
                        employee.role = e.target.value;
                        this.setState({employee})
                        this.validateField("role", employee)

                    }}>
                        <option selected disabled  defaultValue="" >Choose here</option>
                        <option value={"admin"}>Admin</option>
                        <option value={"installer"}>Installer</option>
                        <option value={"storage"}>Storage</option>
                    </select>
                    <p>Enter Role</p>

                </div>

                <div>
                    <select onChange={(e) =>
                    {const employee = {...this.state.employee};
                        employee.admin = Boolean(e.target.value);
                        this.setState({employee})
                        this.validateField("admin", employee)
                    }}>
                        <option selected disabled defaultValue={""}>Choose here</option>
                        <option value={"true"}>Ja</option>
                        <option value={"false"}>Nei</option>
                    </select>
                    <p>Skal brukeren ha admin tillatelser?</p>
                </div>
                <button disabled={!this.state.formValid} onClick={e => this.addPostRequest()}>Add User</button>
            </div>

        )
    }


    render() {

        return(
            this.inputInformation()
        )
    }
}

export default AddUser
