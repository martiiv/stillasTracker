import React, {useState} from "react";
import ReactMapboxGl, {ZoomControl} from "react-mapbox-gl";
import DrawControl from "react-mapbox-gl-draw";
import "@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css";
import postModel from "../../../modelData/postModel";
import {PROJECTS_URL} from "../../../modelData/constantsFile";
import {useQueryClient} from "react-query";
import polygon from "@mapbox/mapbox-gl-draw/src/feature_types/polygon";
import "./map.css"

const Map = ReactMapboxGl({
    accessToken:
        "pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ"
});

export function MapClass(props) {
    const queryClient = useQueryClient()
    const [ok, setOk] = useState(false)

    console.log(props.props)
    const [mapInfo, setMapInfo] = useState([])

    const onDrawCreate = ({features}) => {
        if (features[0].geometry.coordinates[0].length !== 5) {
            console.log("length is invalid")
            window.alert("Invalid geo format. Only valid is 4 points")
            console.log(features)
        } else {
            console.log(features)
            setMapInfo(features)
            setOk(true)

        }

    };


    let geofence = null
    let project = (props.props)
    for (const mapInfoElement of mapInfo) {
        if ((mapInfoElement).hasOwnProperty('geometry')) {
            geofence = {
                "w-position": {
                    latitude: mapInfoElement.geometry.coordinates[0][0][0],
                    longitude: mapInfoElement.geometry.coordinates[0][0][1]
                },
                "x-position": {
                    latitude: mapInfoElement.geometry.coordinates[0][1][0],
                    longitude: mapInfoElement.geometry.coordinates[0][1][1]
                },
                "y-position": {
                    latitude: mapInfoElement.geometry.coordinates[0][2][0],
                    longitude: mapInfoElement.geometry.coordinates[0][2][1]
                },
                "z-position": {
                    latitude: mapInfoElement.geometry.coordinates[0][3][0],
                    longitude: mapInfoElement.geometry.coordinates[0][3][1]
                },
            }
        }
        console.log(geofence)
        project = {...project, geofence}

    }


    const AddProjectRequest = async () => {
        try {
            await postModel(PROJECTS_URL, JSON.stringify(project))
            await queryClient.refetchQueries("allProjects")

        } catch (e) {
            console.log(e)
        }

    }


    return (
        <div className="App">
            <div className={"map"}>
                <Map
                    style="mapbox://styles/mapbox/streets-v9" // eslint-disable-line
                    containerStyle={{
                        height: "80vh",
                        width: "50vw"
                    }}
                    zoom={[17]}
                    center={[Number(project.longitude), Number(project.latitude)]}
                >
                    <DrawControl
                        position="top-left"
                        displayControlsDefault={"polygon"}
                        onDrawCreate={onDrawCreate}
                        controls={{
                            point: false,
                            line_string: false,
                            combine_features: false,
                            uncombine_features: false
                        }}
                        default_mode={"polygon"}
                        clickBuffer={5}
                        onDrawDelete={() => setOk(false)}

                    />
                    <ZoomControl
                        position="bottom-right"
                    />
                </Map>
            </div>
            <button disabled={!ok} onClick={AddProjectRequest}>Add Project</button>
        </div>
    );


}


