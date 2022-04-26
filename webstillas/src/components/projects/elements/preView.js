import React from "react";
import mapboxgl from "mapbox-gl";
import "./preView.css"
import Tabs from "../tabView/Tabs"
import ScaffoldingCardProject from "../../scaffolding/elements/scaffoldingCardProject";
import InfoModal from "./Modal";
import {MAP_STYLE_V11, PROJECTS_URL_WITH_ID, WITH_SCAFFOLDING_URL} from "../../../modelData/constantsFile";
import img from "./../../mapPage/mapbox-marker-icon-20px-orange.png"
import {GetDummyData} from "../getDummyData";

mapboxgl.accessToken = 'pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ';

class PreViewClass extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            data: props.data
        }
        this.mapContainer = React.createRef();
    }

    async componentDidMount() {
        const {data} = this.state
        console.log(data)
        try {
            const map = new mapboxgl.Map({
                container: this.mapContainer.current,
                style: MAP_STYLE_V11,
                center: [data.longitude, data.latitude],
                zoom: 15
            });

            // Create a DOM element for each marker.
            const el = document.createElement('div');
            const width = 50;
            const height = 50;
            el.className = 'marker';
            el.style.backgroundImage = (img);
            el.style.width = `${width}px`;
            el.style.height = `${height}px`;
            el.style.backgroundSize = '100%';

            el.addEventListener('click', () => {
                window.alert("Project: " + data.projectName)
            });

            // Add markers to the map.
            new mapboxgl.Marker(el)
                .setLngLat([data.longitude, data.latitude])
                .addTo(map);

        }catch (e) {
            console.log(e)
        }

    }



    contactInformation(){
        const {data} = this.state
        const project = data
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

    getProjectID(){
        const pathSplit = window.location.href.split("/")
        return pathSplit[pathSplit.length - 1]
    }

    scaffoldingComponents(){
        const {data} = this.state
        return(
            <div className={"grid-container-project-scaffolding"}>
                {data.scaffolding.map((e) => {
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

        return (
            <div className={"preView-Project-Main"}>
                <div ref={this.mapContainer} className="map-container-project"/>
                <div className={"tabs"}>
                    <Tabs>
                        <div label="Kontakt">
                            {this.contactInformation()}
                        </div>
                        <div label="Stillas-komponenter">
                            <InfoModal id={this.getProjectID()}/>
                            {this.scaffoldingComponents()}
                        </div>
                    </Tabs>
                </div>
            </div>
        )
    }

}

function getProjectID(){
    const pathSplit = window.location.href.split("/")
    return pathSplit[pathSplit.length - 1]
}

export const PreView = () => {
    const {isLoading, data} = GetDummyData(["project", getProjectID()], PROJECTS_URL_WITH_ID + getProjectID() + WITH_SCAFFOLDING_URL)
    console.log(data)
    if (isLoading) {
        return <h1>Loading</h1>
    } else {
        return <PreViewClass data={data[0]}/>
    }
}


