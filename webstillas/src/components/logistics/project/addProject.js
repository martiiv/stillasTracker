import React from 'react'


class AddProject extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            projectID: 3223,
            projectName: "",
            latitude: 60.79077759591496,
            longitude:  10.683249543160402,
            period: {
            startDate: "",
                endDate: ""
        },
            size: 0,
            state: "Upcoming",
            address:{
            street: "",
                zipcode: "",
                municipality: "",
                county: ""
        },
            customer: {
            name: "",
                number: 0,
                email: ""
        },
            geofence: {
            "w-position": {
                "latitude": 60.79077759591496,
                "longitude":  10.683249543160402
            },
            "x-position": {
                "latitude": 60.79077759591496,
                "longitude":  10.683249543160402
            },
            "y-position":{
                "latitude": 60.79077759591496,
                "longitude":  10.683249543160402
            },
            "z-position":{
                "latitude": 60.79077759591496,
                "longitude":  10.683249543160402
            }
        }

        }
    }


    dateFormat(date){
        const dateArray = date.split('-')
        return dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0]
    }

    addProjectRequest(){

    }

    generalInformation(){
        //todo integrere med api, slik at brukeren ikke trenger å skrive inn hele addressen.
        return(
            <div className={"general-information"}>
                <h1>Add project</h1>
                <hr/>
                <div className={"input-fields"}>
                    <div>
                        <input type={"text"}  className={"input-text-add"} onChange={event => this.setState({projectName: event.target.value})}/>
                        <p>Enter Project Name</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"} onChange={event =>
                        {const address = {...this.state.address};
                            address.street = event.target.value;
                            this.setState({address})}}/>
                        <p>Enter Address</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"} onChange={event =>
                        {const address = {...this.state.address};
                            address.zipcode = event.target.value;
                            this.setState({address})}}/>
                        <p>Enter Zip Code</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"} onChange={event =>
                        {const address = {...this.state.address};
                            address.municipality = event.target.value;
                            this.setState({address})}}/>
                        <p>Enter Municipality</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"} onChange={event =>
                        {const address = {...this.state.address};
                            address.county = event.target.value;
                            this.setState({address})}}/>
                        <p>Enter County</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"} onChange={event => this.setState({size: event.target.value})}/>
                        <p>Enter size</p>
                    </div>
                    <div>
                        <input type={"date"}  className={"input-date-add"} onChange={event =>
                        {const period = {...this.state.period};
                            period.startDate = this.dateFormat(event.target.value);
                            this.setState({period})}}/>
                        <p>Enter Start date</p>
                    </div>
                    <div>
                        <input type={"date"}   className={"input-date-add"} onChange={event =>
                        {const period = {...this.state.period};
                            period.endDate = this.dateFormat(event.target.value);
                            this.setState({period})}}/>
                        <p>Enter endDate</p>
                    </div>
                    <div>
                        <input type={"text"} className={"input-text-add"} onChange={event =>
                        {const customer = {...this.state.customer};
                            customer.name = event.target.value;
                            this.setState({customer})}}/>
                        <p>Enter Customer Name</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"} onChange={event =>
                        {const customer = {...this.state.customer};
                            customer.number = event.target.value;
                            this.setState({customer})}}/>
                        <p>Enter Customer Number</p>
                    </div>
                    <div>
                        <input type={"text"}  className={"input-text-add"} onChange={event =>
                        {const customer = {...this.state.customer};
                            customer.email = event.target.value;
                            this.setState({customer})}}/>
                        <p>Enter Customer Email</p>
                    </div>
                </div>
                <button>Add Project</button>

            </div>
        )

    }





    render() {
        console.log(this.state)
        return(
            <div>
                {this.generalInformation()}
            </div>
        )
    }

}

export default AddProject
