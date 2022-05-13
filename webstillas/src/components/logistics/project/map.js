import React, {useState} from "react";
import ReactMapboxGl, {ZoomControl} from "react-mapbox-gl";
import DrawControl from "react-mapbox-gl-draw";
import "@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css";
import postModel from "../../../modelData/postModel";
import {MAP_STYLE_V11, PROJECTS_URL} from "../../../modelData/constantsFile";
import {useQueryClient} from "react-query";
import polygon from "@mapbox/mapbox-gl-draw/src/feature_types/polygon";
import "./map.css"
import {MapBoxAPIKey} from "../../../firebaseConfig";
import {AlertCatch} from "../../error/error";

const Map = ReactMapboxGl({
    accessToken: MapBoxAPIKey
});

/**
 * Function that will display a map, and allow a user to draw a polygon.
 *
 * @param props variables sent from previous view.
 * @returns {JSX.Element} Map with draw controllers.
 */
export function MapClass(props) {
    //Query client that will manage the caching.
    const queryClient = useQueryClient()
    //Setting variables
    const [ok, setOk] = useState(false)
    const [mapInfo, setMapInfo] = useState([])

    /**
     * Function that will validate the polygon the user has drawn.
     * If the user has drawn a polygon more than 4 points, the user will get a warning.
     *
     * @param features is the return value from the polygon.
     */
    const onDrawCreate = ({features}) => {
        if (features[0].geometry.coordinates[0].length !== 5) {
            console.log("length is invalid")
            window.alert("Invalid geo format. Only valid is 4 points")
        } else {
            setMapInfo(features)
            setOk(true)
        }
    };



    //Defines the geofence and project object
    let geofence = null
    let project = (props.props)
    for (const mapInfoElement of mapInfo) {
        //if the user has added a valid polygon for the system. the geofence object is initialized.
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
        //The geofence is added to project object.
        project = {...project, geofence}
    }


    /**
     * Function that will add the project to the API.
     *
     * @returns {Promise<void>}
     * @constructor
     */
    const AddProjectRequest = async () => {
        try {
            console.log(JSON.stringify(project))
            await postModel(PROJECTS_URL, JSON.stringify(project))
            await queryClient.refetchQueries("allProjects")
        } catch (e) {
            AlertCatch()
        }
    }



    return (
        <div className="App">
            <div className={"map"}>
                <Map
                    style={MAP_STYLE_V11}
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
                        onDrawDelete={() => setOk(false)}
                    />
                    <ZoomControl
                        position="bottom-right"
                    />
                </Map>
            </div>
            <button className={"confirm-btn"} disabled={!ok || !props.valid} onClick={AddProjectRequest}>Add Project</button>
        </div>
    );
}


