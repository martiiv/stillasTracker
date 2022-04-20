import React from "react";
import "./projects.css"
import CardElement from './elements/card'
import {Route, Routes} from "react-router-dom";


/**
 Class that will create an overview of the projects
 */

//Todo refactor the fetching components from all classes.

class Projects extends React.Component {
    constructor(props) {
        super(props);
        this.state ={
            isLoaded: false,
            projectData: [],
            fromSize: 0,
            toSize: 0,
            fromDate: "",
            toDate: "",
            selectedOption: "",
            searchName: ""
        }
    }


    async componentDidMount() {
        if (sessionStorage.getItem('allProjects') == null){
            const url ="http://10.212.138.205:8080/stillastracking/v1/api/project?scaffolding=true";
            fetch(url)
                .then(res => res.json())
                .then(
                    (result) => {
                        sessionStorage.setItem('allProjects',JSON.stringify(result))
                        console.log('API Kjores')
                        this.setState({
                            isLoaded: true,
                            projectData: result,
                        });
                    },

                    (error) => {
                        this.setState({
                            isLoaded: true,

                        });
                    }
                )
        }else{
           console.log('API Kjøres ikke')
            this.setState({
                isLoaded: true,
            });
        }
    };




    SideBarFunction(){
        const {isLoaded, fromSize, toSize, fromDate, toDate, searchName} = this.state;


        //Todo autocomplete
        if (!isLoaded){
            return <h1>Is Loading data....</h1>
        }else {
            return (
                <div className={"main-sidebar"}>
                    <div>
                        <select onChange={(e) =>
                            this.setState({selectedOption: e.target.value})}>
                            <option defaultValue="" >Velg her</option>
                            <option value={"Active"}>Aktiv</option>
                            <option value={"Inactive"}>Inaktiv</option>
                            <option value={"Upcoming"}>Kommende</option>
                        </select>
                        <p>Status</p>

                    </div>
                    <form className={"filter-content-search"}>
                        <p>Prosjekt Navn</p>

                        <input type="text" value={searchName}
                               onChange={(e) => this.setState({searchName: e.target.value})}/>
                    </form>
                    <form className={"filter-content-input"}>
                        <p>Stillsmengde: </p>
                        <input type="number"  onChange={e => this.setState({fromSize: Number(e.target.value)})}
                               className={"input-field-filter"}/>
                        <input type="number" onChange={e => this.setState({toSize: Number(e.target.value)})}
                               className={"input-field-filter"}/>
                    </form>
                    <form className={"filter-content-input"}>
                        <p>Tidsperiode: </p>
                        <input type="date" value={fromDate} onChange={e => this.setState({fromDate: e.target.value})}
                               className={"input-field-filter"}/>
                        <input type="date" value={toDate} onChange={e => this.setState({toDate: e.target.value})}
                               className={"input-field-filter"}/>
                    </form>
                </div>
            )
        }

    }


    reverseDate(inputDate){
        const dateArray = inputDate.split('-')
        return dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0]
    }


    render() {
        const {projectData, fromSize, toSize, fromDate, toDate, searchName ,selectedOption } = this.state;

        let allProjects
       if (sessionStorage.getItem('allProjects') != null){
             allProjects = sessionStorage.getItem('allProjects')
            console.log('From Storage')
           allProjects = (JSON.parse(allProjects))
        }else {
            console.log('From API')
           allProjects = projectData
        }


        return(
            <div className={"main-project-window"}>
                <div className={"main-sidebar"}>
                    {this.SideBarFunction()}
                </div>
                <div className={"grid-container"}>
                    {allProjects.filter(data => (data.projectName.toLowerCase()).includes(searchName.toLowerCase()))
                        .filter(data => {
                            if (fromSize !== 0){
                                console.log(fromSize)
                                return data.size > fromSize
                            }else {
                                return true
                            }
                        })
                        .filter(data => {
                            if (fromDate !== ""){
                                return this.reverseDate(data.period.startDate) >= fromDate
                            }else {
                                return true
                            }
                        })
                        .filter(data => {
                            if (toDate !== ""){
                                return this.reverseDate(data.period.endDate) <= toDate
                            }else {
                                return true
                            }
                        })
                        .filter(data => {
                            if (toSize !== 0) {
                                return data.size < toSize
                            } else {
                                return true
                            }
                        })
                        .filter(data => {
                            if (!(selectedOption.length === 0 ) && !(selectedOption === "Velg her")){
                                return data.state === selectedOption
                            }else {return true}
                        })
                        .map((e) =>{

                    return(
                        <div>
                            <Routes>
                                <Route path="/project/:id" element={<CardElement data={e} />} />
                            </Routes>
                            <CardElement key = {e.projectID}
                                         id = {e.projectID}
                                         name = {e.projectName}
                                         state = {e.state}
                                         rentPeriod = {e.period.startDate}
                                         size = {e.size}
                                         contactPerson = {e.customer.name}
                                         contactNumber = {e.customer.number}
                                         address_Street = {e.address.street}
                                         address_Municipality = {e.address.municipality}
                                         address_zip = {e.address.zipcode}
                            />
                        </div>

                    );
                    })}


                </div>
            </div>

        );
        }
}

export default Projects;
