import React from "react";
import "./mapPage.css"
import {PROJECTS_WITH_SCAFFOLDING_URL} from "../../modelData/constantsFile";
import {GetDummyData} from "../../modelData/addData";
import ReactMapboxGl, {GeoJSONLayer, Source, Layer, Marker} from "react-mapbox-gl";
import img from "./mapbox-marker-icon-20px-orange.png"


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

        </Map>




        /*<Map
            style="mapbox://styles/mapbox/streets-v9" // eslint-disable-line
            containerStyle={{
                height: "93vh",
                width: "100vw"
            }}
            zoom={[17]}
            center={[lng, lat]}
            >
            <Layer
                type="symbol"
                layout={{ "icon-image": img }}>
                <Feature coordinates={[-0.13235092163085938,51.518250335096376]}/>
            </Layer>
        </Map>*/
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
