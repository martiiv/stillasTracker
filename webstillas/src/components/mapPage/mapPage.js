import React from "react";
import "./mapPage.css"
import mapboxgl from 'mapbox-gl'; // eslint-disable-line import/no-webpack-loader-syntax

mapboxgl.accessToken = "API_KEY"



/**
Class that will create the map-page of the application
 */

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
        new mapboxgl.Map({
            container: this.mapContainer.current,
            style: 'mapbox://styles/mapbox/streets-v11',
            center: [lng, lat],
            zoom: zoom
        });
    }


    render() {
        return(
            <div ref={this.mapContainer} className="map-container"/>
        );
    }
}

export default MapPage;
