import React from "react";
import './topBar.css';
import {
    AppBar, Toolbar, Button
} from '@material-ui/core';
import {Link} from "react-router-dom";

/**
 Class that will create a topbar for the application.
 */

//Todo make list instead of toolbar
    //se hva andre nettsider har gjort
class TopBar extends React.Component {
    render() {
        return(
            <AppBar position="relative">
                <Toolbar className="toolbar" >
                    <Link className="link" to="/prosjekt">
                        <Button className="button">Prosjekter</Button>
                    </Link>
                    <Link className="link" to="/stillas">
                        <Button className="button">Stillasdeler</Button>
                    </Link>
                    <Link className="link" to="/kart">
                        <Button className="button">Kart</Button>
                    </Link>
                    <Link className="link" to="/logistics">
                        <Button className="button">Logistikk</Button>
                    </Link>
                </Toolbar>
            </AppBar>
        );
    }
}

export default TopBar;
