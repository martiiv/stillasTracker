import React from 'react'
import MapClass from "./map";


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
                    size: 0
                },


            projectID: 94328328,
                projectName: "",
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


    dateFormat(date){
        const dateArray = date.split('-')
        return dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0]
    }




    generalInformation(){



        //todo integrere med api, slik at brukeren ikke trenger å skrive inn hele addressen.
        return(
            <div className={"general-information"}>
                <h1>Add project</h1>
                <hr/>
                <form>
                    <div className={"input-fields"}>
                    <div>
                        <input type={"text"}
                               className={"input-text-add"}
                               onChange={event => {this.setState({projectName: event.target.value})}}

                        />
                        <p>Enter Project Name</p>
                    </div>
                    <div>
                        <input type={"text"}
                               className={"input-text-add"}
                               onChange={event => {
                                const address = {...this.state.address};
                                address.street = event.target.value;
                                this.setState({address})}}
                        />
                        <p>Enter Address</p>
                    </div>
                    <div>
                        <input type={"number"}
                               className={"input-text-add"}
                               onChange={event => {
                                   const address = {...this.state.address};
                                    address.zipcode = (event.target.value);
                                    this.setState({address})}}
                        />
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
                        <input type={"number"}  className={"input-text-add"} onChange={event => this.setState({size: Number(event.target.value)})}/>
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
                        <input type={"number"}  className={"input-text-add"} onChange={event =>
                        {const customer = {...this.state.customer};
                            customer.number = Number(event.target.value);
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
                <button onClick={e => this.setState({mapPage: true})}>Next</button>
                </form>
            </div>
        )

    }







    render() {
        console.log(this.state)
        const project = ({
                projectID: 999,
                projectName: this.state.projectName,
                latitude: 60.79077759591496,
                longitude: 10.683249543160402,
                period: {
                    startDate: this.state.period.startDate,
                    endDate: this.state.period.endDate
                },
                size: this.state.size,
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
        console.log(mapPage)

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
