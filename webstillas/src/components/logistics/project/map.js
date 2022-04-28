import React, {useState} from "react";
import ReactMapboxGl from "react-mapbox-gl";
import DrawControl from "react-mapbox-gl-draw";
import "@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css";
import postModel from "../../../modelData/postModel";
import {PROJECTS_URL} from "../../../modelData/constantsFile";
import {useQueryClient} from "react-query";

const Map = ReactMapboxGl({
    accessToken:
        "pk.eyJ1IjoiZmFrZXVzZXJnaXRodWIiLCJhIjoiY2pwOGlneGI4MDNnaDN1c2J0eW5zb2ZiNyJ9.mALv0tCpbYUPtzT7YysA2g"
});

export function MapClass(props) {
    const queryClient = useQueryClient()

    console.log(props.props)
    const [mapInfo, setMapInfo] = useState([])

    const onDrawCreate = ({ features }) => {
        if(features[0].geometry.coordinates[0].length !== 5){
            console.log("length is invalid")
            window.alert("Invalid geo format. Only valid is 4 points")
            console.log(features)
        } else {
            console.log(features)
            setMapInfo(features)
        }

    };


    let geofence = null
    let project = (props.props)
    for (const mapInfoElement of mapInfo) {
        if ((mapInfoElement).hasOwnProperty('geometry')){
            geofence = {
                "w-position":{
                    latitude: mapInfoElement.geometry.coordinates[0][0][0],
                    longitude: mapInfoElement.geometry.coordinates[0][0][1]
                },
                "x-position":{
                    latitude: mapInfoElement.geometry.coordinates[0][1][0],
                    longitude: mapInfoElement.geometry.coordinates[0][1][1]
                },
                "y-position":{
                    latitude: mapInfoElement.geometry.coordinates[0][2][0],
                    longitude: mapInfoElement.geometry.coordinates[0][2][1]
                },
                "z-position":{
                    latitude: mapInfoElement.geometry.coordinates[0][3][0],
                    longitude: mapInfoElement.geometry.coordinates[0][3][1]
                },
            }
        }
        console.log(geofence)
        project = {...project, geofence}

    }


    console.log(JSON.stringify(project))

    const AddProjectRequest = async () => {
        try {
            await postModel(PROJECTS_URL, JSON.stringify(project))
            await queryClient.removeQueries("allProjects")

        }catch (e){
            console.log(e)
        }

    }


    return (
        <div className="App">
            <Map
                style="mapbox://styles/mapbox/streets-v9" // eslint-disable-line
                containerStyle={{
                    height: "100vh",
                    width: "100vw"
                }}
                zoom={[16]}
                center={[-73.9757752418518, 40.69144210646147]}
            >
                <DrawControl
                    position="top-left"
                    displayControlsDefault={"polygon"}
                    onDrawCreate={onDrawCreate}
                />
            </Map>
            <button onClick={AddProjectRequest} >Add Project</button>

        </div>
    );

}


