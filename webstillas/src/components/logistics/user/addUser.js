import React from "react";

class AddUser extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            employee : {
                employeeID: 1001,
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

        console.log(this.state.employee)

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(this.state.employee)
        };
        fetch('http://localhost:8080/stillastracking/v1/api/user', requestOptions)
            .then(response => response.json())
            .then(data => console.log("Added new Project"))
            .catch(err => console.log(err));
    }




    inputInformation(){
        //todo bytt fra forms.select til kun select på alle select i appen.
        return(
            <div className={"input-fields"}>
                <div>
                    <input type={"text"}  className={"input-text-add"} onChange={event =>
                    {
                        const name = {...this.state.name};
                        name.firstName = event.target.value;
                        this.setState({name})}}/>
                    <p>Enter First name</p>
                </div>
                <div>
                    <input type={"text"}  className={"input-text-add"} onChange={event =>{
                        const name = {...this.state.name};
                        name.lastName = event.target.value;
                        this.setState({name})}}/>
                    <p>Enter Last name</p>
                </div>
                <div>
                    <input type={"date"}  className={"input-date-add"} onChange={event =>
                    {const employee = {...this.state.employee};
                        employee.dateOfBirth = this.dateFormat(event.target.value);
                        this.setState({employee})}}/>
                    <p>Enter Birthdate</p>
                </div>
                <div>
                    <input type={"number"}  className={"input-text-add"} onChange={event =>
                    {const employee = {...this.state.employee};
                        employee.phone = Number(event.target.value);
                        this.setState({employee})}}/>
                    <p>Enter User phonenumber</p>
                </div>
                <div>
                    <input type={"text"}  className={"input-text-add"} onChange={event =>
                    {const employee = {...this.state.employee};
                        employee.email = event.target.value;
                        this.setState({employee})}}/>
                    <p>Enter User Email</p>
                </div>
                <div>
                    <select onChange={(e) =>
                    {const employee = {...this.state.employee};
                        employee.role = e.target.value;
                        this.setState({employee})}}>
                        <option defaultValue="" >Choose here</option>
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
                        this.setState({employee})}}>
                        <option defaultValue={""}>Choose here</option>
                        <option value={"true"}>Ja</option>
                        <option value={"false"}>Nei</option>
                    </select>
                    <p>Skal brukeren ha admin tillatelser?</p>
                </div>
                <button onClick={e => this.addPostRequest()}>Add User</button>
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
