import React from "react";
import mapboxgl from "mapbox-gl";
import "./preView.css"
import Tabs from "../tabView/Tabs"
import ScaffoldingCardProject from "../../scaffolding/elements/scaffoldingCardProject";
import InfoModal from "./Modal";
import {
    PROJECTS_URL_WITH_ID,
    PROJECTS_WITH_SCAFFOLDING_URL,
    WITH_SCAFFOLDING_URL
} from "../../../modelData/constantsFile";
import img from "./../../mapPage/mapbox-marker-icon-20px-orange.png"
import {GetDummyData} from "../../../modelData/addData";
import {useQueryClient} from "react-query";
import {SpinnerDefault} from "../../Spinner";
import ReactMapboxGl, {Marker} from "react-mapbox-gl";


const Map = ReactMapboxGl({
    accessToken:
        "pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ"
});


//mapboxgl.accessToken = 'pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ';
/*
//Todo refactor class to function
class PreViewClass extends React.Component {
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

        } catch (e) {
            console.log(e)
        }

    }


    getProjectID() {
        const pathSplit = window.location.href.split("/")
        return pathSplit[pathSplit.length - 1]
    }


    render() {
        return (
            <div className={"preView-Project-Main"}>
                <div ref={this.mapContainer} className="map-container-project"/>
            </div>
        )
    }
}*/


function PreViewFunction(props) {
    const data = props.data


    return (
        <div className={"preView-Project-Main"}>
            <Map
                style="mapbox://styles/mapbox/streets-v9" // eslint-disable-line
                containerStyle={{
                    height: "93vh",
                    width: "40vw"
                }}
                zoom={[17]}
                center={[data.longitude, data.latitude]}
            >
                <Marker
                    offsetTop={-48}
                    offsetLeft={-24}
                    coordinates={[data.longitude, data.latitude]}
                >
                    <img src={img} alt={""}/>
                </Marker>
            </Map>
        </div>


    )

}


function getProjectID() {
    const pathSplit = window.location.href.split("/")
    return pathSplit[pathSplit.length - 1]
}


function scaffoldingComponents(data) {
    return (
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


function contactInformation(project) {
    return (
        <div>
            <section className={"info-card"}>
                <div className={"information-highlights preview-info"}>
                    <h3>Prosjekt Informasjon</h3>
                    <ul className={"contact-list"}>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Kunde</span>
                            <span className={"right-contact-text"}>{project[0].customer.name}</span>
                        </li>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Størrelse</span>
                            <span className={"right-contact-text"}>{project[0].size} &#13217;</span>
                        </li>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Status</span>
                            <span className={"right-contact-text"}>{project[0].state}</span>
                        </li>

                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Periode</span>
                            <span
                                className={"right-contact-text"}>{project[0].period.startDate} - {project[0].period.endDate}  </span>
                        </li>
                    </ul>
                </div>
            </section>
            <section className={"info-card"}>
                <div className={"information-highlights preview-info"}>
                    <h3>Kontakt Informasjon</h3>
                    <ul className={"contact-list"}>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Kontakt person</span>
                            <span className={"right-contact-text"}>{project[0].customer.name}</span>
                        </li>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Telefon nummer</span>
                            <span className={"right-contact-text"}>{project[0].customer.number}</span>
                        </li>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>E-mail</span>
                            <span className={"right-contact-text"}>{project[0].customer.email}</span>
                        </li>
                        <li className={"horizontal-list-contact"}>
                            <span className={"left-contact-text"}>Adresse</span>
                            <span
                                className={"right-contact-text"}>{project[0].address.street}, {project[0].address.zipcode} {project[0].address.municipality}</span>
                        </li>
                    </ul>
                </div>
            </section>


        </div>

    )
}


export const PreView = () => {
    const queryClient = useQueryClient()

    const {
        isLoading: projectLoad,
        data: project
    } = GetDummyData(["project", getProjectID()], PROJECTS_URL_WITH_ID + getProjectID() + WITH_SCAFFOLDING_URL)
    let projects
    let allProjectsLoading
    if (queryClient.getQueryData("allProjects") !== undefined) {
        projects = queryClient.getQueryData("allProjects")
    }
    const {isLoading: allProjects, data} = GetDummyData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)
    projects = data
    allProjectsLoading = allProjects


    if (allProjectsLoading || projectLoad) {
        return <SpinnerDefault/>

    } else {
        return (
            <div className={"preView-Project-Main"}>
                <div className={"map-preview"}>
                    <PreViewFunction data={project[0]}/>
                </div>
                <div className={"tabs"}>
                    <Tabs>
                        <div label="Kontakt">
                            {contactInformation(project)}
                        </div>
                        <div label="stillas-komponenter">
                            <InfoModal id={getProjectID()}/>
                            {scaffoldingComponents(project[0])}
                        </div>
                    </Tabs>
                </div>
            </div>
        )
    }
}




