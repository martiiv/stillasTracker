import React from "react";
import mapboxgl from "mapbox-gl";
import "./preView.css"
import Tabs from "../tabView/Tabs"

mapboxgl.accessToken = 'pk.eyJ1IjoiYWxla3NhYWIxIiwiYSI6ImNrbnFjbms1ODBkaWEyb3F3OTZiMWd6M2gifQ.vzOmLzHH3RXFlSsCRrxODQ';



class PreView extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            data:[]
        }
        this.mapContainer = React.createRef();

    }


    getProjectID(){
        const pathSplit = window.location.href.split("/")
        return pathSplit[pathSplit.length - 1]
    }

    async fetchData() {
        const path = this.getProjectID()
        console.log(path)
        const url ="http://10.212.138.205:8080/stillastracking/v1/api/project?id=" + path + "&scaffolding=true";
        fetch(url)
            .then(res => res.json())
            .then(
                (result) => {
                    sessionStorage.setItem('project', (result))
                    this.setState({
                        isLoaded: true,
                        data: result
                    });
                },
                (error) => {
                    this.setState({
                        isLoaded: true,

                    });
                }
            )
    }

    //todo refactor dette
    async componentDidMount() {
        const path = this.getProjectID()
        console.log(path)
        await this.fetchData()
        const url ="http://10.212.138.205:8080/stillastracking/v1/api/project?id=" + path + "&scaffolding=true";
        fetch(url)
            .then(res => res.json())
            .then(
                (result) => {
                    const map = new mapboxgl.Map({
                        container: this.mapContainer.current,
                        style: 'mapbox://styles/mapbox/streets-v11',
                        center: [result[0].longitude, result[0].latitude],
                        zoom: 15
                    });

                    // Add markers to the map.
                    for (const marker of result) {
                        // Create a DOM element for each marker.
                        const el = document.createElement('div');
                        const width = result.size;
                        const height = result.size;
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


    }



    kontakt(){
        const {data } = this.state;

        return(
            <div className={"informasjon"}>
                {JSON.stringify(data)}
            </div>

        )
    }


    render() {

        return(
            <div className={"preView-Project-Main"}>
                <div ref={this.mapContainer} className="map-container-project"/>
                <div className={"tabs"}>
                    <Tabs>
                        <div label="Kontakt">
                            {this.kontakt()}
                        </div>
                        <div label="Stillas komponenter">
                            After 'while, <em>Crocodile</em>!
                        </div>
                    </Tabs>
                </div>
            </div>

        )
    }

}


export default PreView




