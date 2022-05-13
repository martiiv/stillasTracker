import React, {useState} from "react";
import ReactMapboxGl, {ZoomControl} from "react-mapbox-gl";
import DrawControl from "react-mapbox-gl-draw";
import "@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css";
import postModel from "../Middleware/postModel";
import {MAP_STYLE_V11, PROJECTS_URL} from "../Constants/apiURL";
import {useQueryClient} from "react-query";
import polygon from "@mapbox/mapbox-gl-draw/src/feature_types/polygon";
import "../Assets/Styles/map.css"
import {MapBoxAPIKey} from "../Config/firebaseConfig";
import {AlertCatch} from "../components/Indicators/error";
import {Button, Spinner} from "react-bootstrap";
import {PROJECT_URL} from "../Constants/webURL";
import {useNavigate} from "react-router-dom";

const AddProjectMap = ReactMapboxGl({
    accessToken: MapBoxAPIKey
});

/**
 * Function that will display a map, and allow a user to draw a polygon.
 *
 * @param props variables sent from previous view.
 * @returns {JSX.Element} AddProjectMap with draw controllers.
 */
export function MapClass(props) {
    //Query client that will manage the caching.
    const queryClient = useQueryClient()
    let navigate = useNavigate();

    //Setting variables
    const [ok, setOk] = useState(false)
    const [mapInfo, setMapInfo] = useState([])
    const [buttonPressed, setButtonPressed] = useState(false)


    /**
     * Function that will validate the polygon the user has drawn.
     * If the user has drawn a polygon more than 4 points, the user will get a warning.
     *
     * @param features is the return value from the polygon.
     */
    const onDrawCreate = ({features}) => {
        if (features[0].geometry.coordinates[0].length !== 5) {
            window.alert("Format ikke godkjent! Kun fire punkter er tillatt ")
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
            setButtonPressed(true)
            await postModel(PROJECTS_URL, (project))
            await queryClient.refetchQueries("allProjects").then(
                navigate(PROJECT_URL)
            )
        } catch (e) {
            setButtonPressed(true)
            AlertCatch()
        }
    }



    return (
        <div className="App">
            <div className={"map"}>
                <AddProjectMap
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
                </AddProjectMap>
            </div>

            {buttonPressed ? <Button className={"confirm-btn"}
                    disabled>
                    <Spinner
                        as="span"
                        animation="grow"
                        size="sm"
                        role="status"
                        aria-hidden="true"
                    />
                    Legger til
                </Button> :
                <button className={"confirm-btn"} disabled={!ok || !props.valid} onClick={AddProjectRequest}>Add Project</button>}



        </div>
    );
}


