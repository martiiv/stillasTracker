import React from "react";
import mapboxgl from "mapbox-gl";
import "./preView.css"
import Tabs from "../tabView/Tabs"
import ScaffoldingCardProject from "../../scaffolding/elements/scaffoldingCardProject";
import InfoModal from "./Modal";
import fetchModel from "../../../modelData/fetchData";
import {MAP_STYLE_V11, PROJECTS_URL_WITH_ID, WITH_SCAFFOLDING_URL} from "../../../modelData/constantsFile";
import img from "./../../mapPage/mapbox-marker-icon-20px-orange.png"

mapboxgl.accessToken = 'pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ';

class PreView extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            isLoaded: false,
            data:null
        }
        this.mapContainer = React.createRef();
    }


    getProjectID(){
        const pathSplit = window.location.href.split("/")
        return pathSplit[pathSplit.length - 1]
    }

    async componentDidMount() {
        const path = this.getProjectID()
        try {
            const projectResult = await fetchModel(PROJECTS_URL_WITH_ID + path + WITH_SCAFFOLDING_URL)
            sessionStorage.setItem('project', (JSON.stringify(projectResult[0])))
            this.setState({
                isLoaded: true,
                data: projectResult[0],
            })
            const map = new mapboxgl.Map({
                container: this.mapContainer.current,
                style: MAP_STYLE_V11,
                center: [projectResult[0].longitude, projectResult[0].latitude],
                zoom: 15
            });
            for (const marker of projectResult) {
                // Create a DOM element for each marker.
                const el = document.createElement('div');
                const width = projectResult[0].size/100;
                const height = projectResult[0].size/100;
                el.className = 'marker';
                el.style.backgroundImage = (img);
                el.style.width = `${width}px`;
                el.style.height = `${height}px`;
                el.style.backgroundSize = '100%';

                el.addEventListener('click', () => {
                    window.alert("Project: " + marker.projectName)
                });

                // Add markers to the map.
                new mapboxgl.Marker(el)
                    .setLngLat([marker.longitude, marker.latitude])
                    .addTo(map);
            }
        }catch (e) {
            console.log(e)
        }

    }



    contactInformation(){
        let project
        if (sessionStorage.getItem('project') != null){
            const scaffold = sessionStorage.getItem('project')
            console.log('From Storage')
            project = (JSON.parse(scaffold))
        }else {
            console.log('From API')
            project = this.state.data
        }


        return(
            <section className={"contact-highlights-cta"}>
                <div className={"information-highlights"}>
                    <ul className={"contact-list"}>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Navn/Bedrift</span>
                            <span className={"right-contact-text"}>{project.customer.name}</span>
                        </li>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Telefon nummer</span>
                            <span className={"right-contact-text"}>{project.customer.number}</span>
                        </li>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Adresse</span>
                            <span className={"right-contact-text"}>{project.address.street}, {project.address.zipcode} {project.address.municipality}</span>
                        </li>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>E-mail</span>
                            <span className={"right-contact-text"}>{project.customer.email}</span>
                        </li>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Periode</span>
                            <span className={"right-contact-text"}>{project.period.startDate} to {project.period.endDate}  </span>
                        </li>
                    </ul>
                </div>
            </section>
        )
    }



    scaffoldingComponents(){
        let project
        if (sessionStorage.getItem('project') != null){
            const scaffold = sessionStorage.getItem('project')
            console.log('From Storage')
            project = (JSON.parse(scaffold))
        }else {
            console.log('From API')
            project = this.state.data
        }

        return(
            <div className={"grid-container-project-scaffolding"}>
                {project.scaffolding.map((e) => {
                    return (
                        <ScaffoldingCardProject
                            key={e.type}
                            type={e.type}
                            expected={e.Quantity.expected}
                            registered={e.Quantity.registered}
                        />
                    )
                })}
            </div>
        )
    }


    render() {
        const {isLoaded} = this.state
        if (!isLoaded){
            return <h1>Is Loading Data....</h1>
        }else{
            return (
                <div className={"preView-Project-Main"}>
                    <div ref={this.mapContainer} className="map-container-project"/>
                    <div className={"tabs"}>
                        <Tabs>
                            <div label="Kontakt">
                                {this.contactInformation()}
                            </div>
                            <div label="Stillas-komponenter">
                                <InfoModal />
                                {this.scaffoldingComponents()}
                            </div>
                        </Tabs>
                    </div>
                </div>
            )
        }
    }
}

export default PreView




