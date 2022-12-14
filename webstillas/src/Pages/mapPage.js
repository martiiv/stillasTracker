import React from "react";
import "../Assets/Styles/mapPage.css"
import {MAP_STYLE_V11, PROJECTS_WITH_SCAFFOLDING_URL} from "../Constants/apiURL";
import {GetCachingData} from "../Middleware/addData";
import ReactMapboxGl, {ScaleControl, Marker, ZoomControl} from "react-mapbox-gl";
import {MapBoxAPIKey} from "../Config/firebaseConfig";
import img from "../Assets/Images/marker.png"
import {InternalServerError} from "../components/Indicators/error";
import {SpinnerDefault} from "../components/Indicators/Spinner";

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
                        key = {res.projectID}
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
    const {isLoading, data, isError} = GetCachingData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)
    if (isLoading) {
        return <SpinnerDefault />
    } else if(isError){
        return <InternalServerError />
    } else {
        const projects = JSON.parse(data.text)
        return <MapPageClass data={projects}/>
    }
}
