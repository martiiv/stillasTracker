import React from "react";
import "./mapPage.css"
import mapboxgl from 'mapbox-gl';
import {MAP_STYLE_V11, PROJECTS_URL, PROJECTS_WITH_SCAFFOLDING_URL} from "../../modelData/constantsFile";
import {GetDummyData} from "../../modelData/addData";
import {useQueryClient} from "react-query";

mapboxgl.accessToken = 'pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ';

/**
 Class that will create the map-page of the application
 */
//Kode hentet fra https://docs.mapbox.com/help/tutorials/use-mapbox-gl-js-with-react/
class MapPageClass extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isLoaded: false,
            projectData: props.data,
            lng: 10.69163,
            lat: 60.79543,
            zoom: 9
        };
        this.mapContainer = React.createRef();
    }

    async componentDidMount() {
        const { lng, lat, zoom, projectData} = this.state;
        console.log(projectData)
        try {
            const projectResult = projectData
            const map = new mapboxgl.Map({
                container: this.mapContainer.current,
                style: MAP_STYLE_V11,
                center: [lng, lat],
                zoom: zoom
            });
            for (const marker of projectResult) {
                const el = document.createElement('div');
                const width = projectResult.size;
                const height = projectResult.size;
                el.className = 'marker';
                el.style.backgroundImage = ("src/components/mapPage/mapbox-marker-icon-20px-orange.png");
                el.style.width = `${width}px`;
                el.style.height = `${height}px`;
                el.style.backgroundSize = '100%';

                el.addEventListener('click', () => {
                    window.alert("Project: " + marker.projectName)
                });

                // Add markers to the map.
                new mapboxgl.Marker(el)
                    .setLngLat([marker.longitude, marker.latitude])
                    .addTo(map);
            }
        }catch (e) {
            console.log(e)
        }
    }

    render() {
        return(
            <div ref={this.mapContainer} className="map-container"/>
        );
    }
}


export const MapPage = () => {
    let projects
    let allProjectsLoading

    const {isLoading: allProjects, data} = GetDummyData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)
    projects = data
    allProjectsLoading = allProjects


    if (allProjectsLoading) {
        return <h1>Loading</h1>
    } else {
        return <MapPageClass data={projects}/>
    }
}
