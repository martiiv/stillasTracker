import React  from "react";
import ReactMapboxGl from "react-mapbox-gl";
import DrawControl from "react-mapbox-gl-draw";
import "@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css";
import addProject from "./addProject";
import AddProject from "./addProject";

const Map = ReactMapboxGl({
    accessToken:
        "pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ"
});

class MapClass extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            mapInfo: []
        }
    }
    onDrawCreate = ({ features }) => {
        this.setState({mapInfo: features})
    };

    onDrawUpdate = ({ features }) => {
        console.log({ features });
    };

    addProjectRequest = (inputBody) => {
        console.log(inputBody)
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(inputBody)
        };
        fetch('http://localhost:8080/stillastracking/v1/api/project', requestOptions)
            .then(response => response.json())
            .then(data => console.log("Added new Project"))
            .catch(err => console.log(err));
    }


    render() {


        const {mapInfo} = this.state;

        let geofence = null
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
        }


        const project = (this.props.project)
        project.geofence = geofence

        console.log(JSON.stringify(project))

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
                        onDrawCreate={this.onDrawCreate}
                        onDrawUpdate={this.onDrawUpdate}
                    />
                </Map>
                <button onClick={e => this.addProjectRequest(project)}>Add Project</button>
            </div>
        );
    }
}

export default MapClass
