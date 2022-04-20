import React from "react";
import "./mapPage.css"
import mapboxgl from 'mapbox-gl';
import postModel from "../../modelData/postModel";
import {PROJECT_URL, PROJECTS_URL} from "../../modelData/constantsFile";
import fetchModel from "../../modelData/fetchData";

mapboxgl.accessToken = 'pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ';



/**
Class that will create the map-page of the application
 */
//Kode hentet fra https://docs.mapbox.com/help/tutorials/use-mapbox-gl-js-with-react/
class MapPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isLoaded: false,
            projectData: [],
            lng: 10.69163,
            lat: 60.79543,
            zoom: 9
        };
        this.mapContainer = React.createRef();
    }


    async componentDidMount() {
        const { lng, lat, zoom} = this.state;
        try {
            const projectResult = await fetchModel(PROJECTS_URL)
            const map = new mapboxgl.Map({
                container: this.mapContainer.current,
                style: 'mapbox://styles/mapbox/streets-v11',
                center: [lng, lat],
                zoom: zoom
            });
            for (const marker of projectResult) {
                // Create a DOM element for each marker.
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

export default MapPage;
