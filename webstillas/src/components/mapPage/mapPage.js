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
            isLoaded: false,
            projectData: [],
            lng: 10.69163,
            lat: 60.79543,
            zoom: 9
        };
        this.mapContainer = React.createRef();
    }


    async fetchData() {
        const url ="http://10.212.138.205:8080/stillastracking/v1/api/project/";
        fetch(url)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        projectData: result
                    });
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    this.setState({
                        isLoaded: true,

                    });
                }
            )
    }



    componentDidMount() {
        const { lng, lat, zoom, projectData} = this.state;

        const url ="http://10.212.138.205:8080/stillastracking/v1/api/project/";
        fetch(url)
            .then(res => res.json())
            .then(
                (result) => {
                    console.log(result)
                    const map = new mapboxgl.Map({
                        container: this.mapContainer.current,
                        style: 'mapbox://styles/mapbox/streets-v11',
                        center: [lng, lat],
                        zoom: zoom
                    });

                    // Add markers to the map.
                    for (const marker of result) {
                        // Create a DOM element for each marker.
                        const el = document.createElement('div');
                        const width = 50;
                        const height = 50;
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
                            .setLngLat([marker.longitude, marker.latitude])
                            .addTo(map);
                    }

                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    this.setState({
                        isLoaded: true,

                    });
                }
            )

        const geojson = {
            'type': 'FeatureCollection',
            'features': [
                {
                    'type': 'Feature',
                    'properties': {
                        'message': 'Lillehammer',
                        'iconSize': [60, 60],
                        'coordinates': [10.46628, 61.11514]

                    },
                    'geometry': {
                        'type': 'Point',
                    }
                },
                {
                    'type': 'Feature',
                    'properties': {
                        'message': 'Gjøvik',
                        'iconSize': [50, 50],
                        'coordinates': [10.69155, 60.79574]

                    },
                    'geometry': {
                        'type': 'Point',
                    }
                },
                {
                    'type': 'Feature',
                    'properties': {
                        'message': 'Hamar',
                        'iconSize': [40, 40],
                        'coordinates': [10.681777071532371, 60.7905060889568]
                    },
                    'geometry': {
                        'type': 'Point',

                    }
                }
            ]
        };



    }




    render() {
        const {projectData} = this.state;
        console.log(projectData)

        return(
          <div ref={this.mapContainer} className="map-container"/>
        );
    }
}

export default MapPage;
