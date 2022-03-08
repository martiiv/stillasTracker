import React from "react";
import './topBar.css';
import {
    AppBar, Toolbar, Button
} from '@material-ui/core';

/**
 Class that will create a topbar for the application.
 */

class TopBar extends React.Component {
    render() {
        return(
            <AppBar className="appbar">
                <Toolbar className="toolbar">
                        <Button className="button">Prosjekter</Button>
                        <Button className="button">Stillasdeler</Button>
                        <Button className="button">Kart</Button>
                        <Button className="button">Logistikk</Button>
                </Toolbar>
            </AppBar>


        );
    }
}

export default TopBar;
