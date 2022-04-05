import React from "react";
import "./mapPage.css"
import mapboxgl from 'mapbox-gl';

mapboxgl.accessToken = 'pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ';



/**
Class that will create the map-page of the application
 */
//Kode hentet fra https://docs.mapbox.com/help/tutorials/use-mapbox-gl-js-with-react/

class MapPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            lng: 10.69163,
            lat: 60.79543,
            zoom: 9
        };
        this.mapContainer = React.createRef();
    }



    componentDidMount() {
        const { lng, lat, zoom } = this.state;
        const geojson = {
            'type': 'FeatureCollection',
            'features': [
                {
                    'type': 'Feature',
                    'properties': {
                        'message': 'Lillehammer',
                        'iconSize': [60, 60]
                    },
                    'geometry': {
                        'type': 'Point',
                        'coordinates': [10.46628, 61.11514]
                    }
                },
                {
                    'type': 'Feature',
                    'properties': {
                        'message': 'Gjøvik',
                        'iconSize': [50, 50]
                    },
                    'geometry': {
                        'type': 'Point',
                        'coordinates': [10.69155, 60.79574]
                    }
                },
                {
                    'type': 'Feature',
                    'properties': {
                        'message': 'Hamar',
                        'iconSize': [40, 40]
                    },
                    'geometry': {
                        'type': 'Point',
                        'coordinates': [11.06798, 60.7945]
                    }
                }
            ]
        };
        const map = new mapboxgl.Map({
            container: this.mapContainer.current,
            style: 'mapbox://styles/mapbox/streets-v11',
            center: [lng, lat],
            zoom: zoom
        });

        // Add markers to the map.
        for (const marker of geojson.features) {
            // Create a DOM element for each marker.
            const el = document.createElement('div');
            const width = marker.properties.iconSize[0];
            const height = marker.properties.iconSize[1];
            el.className = 'marker';
            el.style.backgroundImage = ("src/components/mapPage/mapbox-marker-icon-20px-orange.png");
            el.style.width = `${width}px`;
            el.style.height = `${height}px`;
            el.style.backgroundSize = '100%';

            el.addEventListener('click', () => {
                window.alert(marker.properties.message);
            });

            // Add markers to the map.
            new mapboxgl.Marker(el)
                .setLngLat(marker.geometry.coordinates)
                .addTo(map);
        }
    }




    render() {
        return(
          <div ref={this.mapContainer} className="map-container"/>
        );
    }
}

export default MapPage;
