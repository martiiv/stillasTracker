import React from "react";
import "../../Assets/Styles/preView.css"
import Tabs from "../../Layout/tabView/Tabs"
import ScaffoldingCardProject from "../../components/projects/scaffoldingCardProject";
import InfoModal from "./transferScaffoldingModal";
import {
    MAP_STYLE_V11,
    PROJECTS_URL_WITH_ID,
    WITH_SCAFFOLDING_URL
} from "../../Constants/apiURL";
import img from "../../Assets/Images/marker.png"
import {GetCachingData} from "../../Middleware/addData";
import {SpinnerDefault} from "../Indicators/Spinner";
import ReactMapboxGl, {Marker} from "react-mapbox-gl";
import {MapBoxAPIKey} from "../../Config/firebaseConfig";
import {InternalServerError} from "../Indicators/error";


const Map = ReactMapboxGl({
    accessToken: MapBoxAPIKey
});

/**
 * Function that will display a map where the project is located
 * @param props the project to be displayed
 * @returns {JSX.Element}
 */
function PreViewFunction(props) {
    const data = props.data

    return (
        <div className={"preView-Project-Main"}>
            <Map
                style= {MAP_STYLE_V11} // eslint-disable-line
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

/**
 * Function that will get the project id form the url.
 * @returns {string} project id
 */
function getProjectID() {
    const pathSplit = window.location.href.split("/")
    return pathSplit[pathSplit.length - 1]
}


/**
 * Function that will display all the different types of scaffolding
 * @param data scaffolding parts
 * @returns {JSX.Element}
 */
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


/**
 * Will return cards with project information
 * @param project to be displayed
 * @returns {JSX.Element}
 */
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
            <section className={"info-card contact-information"}>
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


/**
 * Function that displays the whole page.
 * @returns {JSX.Element}
 */
export const PreViewSite = () => {
    //Fetching the data of specific project.
    const {isLoading: projectLoad, data, isError } = GetCachingData(["project", getProjectID()], PROJECTS_URL_WITH_ID + getProjectID() + WITH_SCAFFOLDING_URL)

    if (projectLoad) {
        return <SpinnerDefault/>
    } else if(isError){
        return <InternalServerError />
    } else {
        const project = JSON.parse(data.text)
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
                        <div label="Stillasdeler">
                            <InfoModal id={getProjectID()}/>
                            {scaffoldingComponents(project[0])}
                        </div>
                    </Tabs>
                </div>
            </div>
        )
    }
}




