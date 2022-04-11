import React from "react";
import "./projects.css"
import CardElement from './elements/card'
import {Route, Routes} from "react-router-dom";


/**
 Class that will create an overview of the projects
 */

class Projects extends React.Component {
    constructor(props) {
        super(props);
        this.state ={
            isLoaded: false,
            projectData: [],
            fromSize: 0,
            toSize: 1000000,
            fromDate: new Date().toLocaleString(),
            toDate: new Date().toLocaleString(),
            selectedOption: '',
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
                    // Note: it's important to handle errors here
                    // instead of a catch() block so that we don't swallow
                    // exceptions from actual bugs in components.
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
        const {isLoaded, fromSize, toSize, fromDate, toDate, searchName ,selectedOption} = this.state;

        if (!isLoaded){
            return <h1>Is Loading data....</h1>
        }else {
            return (
                <div className={"main-sidebar"}>

                    <form className={"filter-content"}>
                        <input type="radio" value="active"
                               onChange={(e) => this.setState({selectedOption: e.target.value})}/> Aktiv
                        <input type="radio" value="inactive"
                               onChange={(e) => this.setState({selectedOption: e.target.value})}/> Inaktiv
                        <input type="radio" value="upcoming"
                               onChange={(e) => this.setState({selectedOption: e.target.value})}/> Kommende
                    </form>
                    <form className={"filter-content-search"}>
                        <input type="text" value={searchName}
                               onChange={(e) => this.setState({searchName: e.target.value})}/>
                    </form>
                    <form className={"filter-content-input"}>
                        <p>Stillsmengde: </p>
                        <input type="number" value={fromSize} onChange={e => this.setState({fromSize: e.target.value})}
                               className={"input-field-filter"}/>
                        <input type="number" value={toSize} onChange={e => this.setState({toSize: e.target.value})}
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

    isAfter(date1, date2) {
        return date1 > date2;
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
                {this.SideBarFunction()}
                <div className={"grid-container"}>



                    {allProjects.filter(data => (data.projectName.toLowerCase()).includes(searchName.toLowerCase()))
                        .filter(data => data.size > fromSize)
                        //todo add date filter
                        .filter(data => data.size < toSize)
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
