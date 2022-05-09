import React from "react";
import "./mapPage.css"
import {PROJECTS_WITH_SCAFFOLDING_URL} from "../../modelData/constantsFile";
import {GetDummyData} from "../../modelData/addData";
import ReactMapboxGl, {ScaleControl, Source, Layer, Marker, ZoomControl} from "react-mapbox-gl";
import img from "./mapbox-marker-icon-20px-orange.png"
import {NavigationControl} from "react-map-gl";


const Map = ReactMapboxGl({
    accessToken:
        "pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ"
});


/**
 Class that will create the map-page of the application
 */
//Kode hentet fra https://docs.mapbox.com/help/tutorials/use-mapbox-gl-js-with-react/
function MapPageClass(props) {
    const projectData = props.data
    const lng = 10.69155
    const lat = 60.79574
    const zoom = 9


    const onClick = (data) =>{
        window.alert(data.projectName)
    }


    return (
        <Map
            style="mapbox://styles/mapbox/streets-v10"
            containerStyle={{
                height: '100vh',
                width: '100vw'
            }}
            center={[lng, lat]}
        >


            {projectData.map(res => {
                return(
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


export const MapPage = () => {
    const {isLoading, data} = GetDummyData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)
    if (isLoading) {
        return <h1>Loading</h1>
    } else {
        return <MapPageClass data={data}/>
    }
}
