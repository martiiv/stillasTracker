import React, {useState} from "react";
import mapboxgl from "mapbox-gl";
import "./preView.css"
import Tabs from "../tabView/Tabs"
import ScaffoldingCardProject from "../../scaffolding/elements/scaffoldingCardProject";
import InfoModal from "./Modal";




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


    //todo refactor dette
    async componentDidMount() {
        const path = this.getProjectID()
        console.log(path)
        const url ="http://10.212.138.205:8080/stillastracking/v1/api/project?id=" + path + "&scaffolding=true";
        await fetch(url)
            .then(res => res.json())
            .then(
                (result) => {
                    sessionStorage.setItem('project', (JSON.stringify(result[0])))
                   this.setState({
                       isLoaded: true,

                       data: result[0],
                   })
                    const map = new mapboxgl.Map({
                        container: this.mapContainer.current,
                        style: 'mapbox://styles/mapbox/streets-v11',
                        center: [result[0].longitude, result[0].latitude],
                        zoom: 15
                    });

                    // Add markers to the map.
                    for (const marker of result) {
                        // Create a DOM element for each marker.
                        const el = document.createElement('div');
                        const width = result.size;
                        const height = result.size;
                        el.className = 'marker';
                        el.style.backgroundImage = ("src/components/mapPage/mapbox-marker-icon-20px-orange.png");
                        el.style.width = `${width}px`;
                        el.style.height = `${height}px`;
                        el.style.backgroundSize = '100%';

                        el.addEventListener('click', () => {
                            window.alert(marker.properties.message);
                        });

                        // Add markers to the map.
                        new mapboxgl.Marker(el)
                            .setLngLat([marker.longitude, marker.latitude])
                            .addTo(map);
                    }

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
    }


    contactInformation(){
        const {data} = this.state
        let project
        if (sessionStorage.getItem('project') != null){
            const scaffold = sessionStorage.getItem('project')
            console.log('From Storage')
            project = (JSON.parse(scaffold))
        }else {
            console.log('From API')
            project = data
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
        const {data} = this.state
        let project
        if (sessionStorage.getItem('project') != null){
            const scaffold = sessionStorage.getItem('project')
            console.log('From Storage')
            project = (JSON.parse(scaffold))
        }else {
            console.log('From API')
            project = data
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
        const {isLoaded, data} = this.state
        let project
        if (sessionStorage.getItem('project') != null){
            const scaffold = sessionStorage.getItem('project')
            console.log('From Storage')
            project = (JSON.parse(scaffold))
        }else {
            console.log('From API')
            project = data
        }


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




