import React from "react";
import "./mapPage.css"
import {MAP_STYLE_V11, PROJECTS_WITH_SCAFFOLDING_URL} from "../../modelData/constantsFile";
import {GetDummyData} from "../../modelData/addData";
import ReactMapboxGl, {ScaleControl, Marker, ZoomControl} from "react-mapbox-gl";
import {MapBoxAPIKey} from "../../firebaseConfig";
import img from "./marker.png"
import {InternalServerError} from "../error/error";
import {SpinnerDefault} from "../Spinner";

const Map = ReactMapboxGl({
    accessToken: MapBoxAPIKey
});


//Kode hentet fra https://docs.mapbox.com/help/tutorials/use-mapbox-gl-js-with-react/
/**
 * Function that will display a map on the website, with markers on lat, long.
 *
 * @param props information of project
 * @returns {JSX.Element}
 */
function MapPageClass(props) {

    const projectData = props.data
    const lng = 10.69155
    const lat = 60.79574

    const onClick = (data) => {
        window.alert(data.projectName)
    }

    //Returns a map centered at desired longitude and latitude.
    return (
        <Map
            style={MAP_STYLE_V11}
            containerStyle={{
                height: '100vh',
                width: '100vw'
            }}
            center={[lng, lat]}
        >


            {projectData.map(res => {
                return (
                    <Marker
                        offsetTop={-48}
                        offsetLeft={-24}
                        coordinates={[res.longitude, res.latitude]}
                        onClick={() => onClick(res)}
                    >
                        <img src={img} alt={""}/>
                    </Marker>
                )
            })}

            <ZoomControl
                position="top-right"
            />

            <ScaleControl/>
        </Map>

    );

}


/**
 * Function that will fetch the information from API/Cache.
 * If loading a spinner will be displayed.
 *
 * @returns {JSX.Element}
 */
export const MapPage = () => {
    const {isLoading, data, isError} = GetDummyData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)
    if (isLoading) {
        return <SpinnerDefault />
    } else if(isError){
        return <InternalServerError />
    } else {
        const projects = JSON.parse(data.text)
        return <MapPageClass data={projects}/>
    }
}
