import React from "react";
import './topBar.css';
import {
    AppBar, Toolbar, Button
} from '@material-ui/core';
import {Link} from "react-router-dom";
import {DropdownButton} from "react-bootstrap";
import DropdownItem from "react-bootstrap/DropdownItem";
import {useUserAuth} from "../../context/UserAuthContext";
import {auth} from "../../firebase"

/**
 Class that will create a topbar for the application.
 */

//Todo make list instead of toolbar
    //se hva andre nettsider har gjort
const TopBar = () => {
    const { logOut } = useUserAuth();


    if (!auth.currentUser){
        return (
            <AppBar position="sticky">
                <Toolbar className="toolbar" >

                </Toolbar>
            </AppBar>
        )
    }else {
        return(
            <AppBar position="sticky">
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
                   {/* <Link className="link" to="/logistics">
                        <Button className="button">Logistikk</Button>
                    </Link>*/}

                    <DropdownButton id="dropdown-button"
                                    title= {"Logistikk"}
                                    size="sm"

                    >
                        <DropdownItem href="/addproject/">Legg til prosjekt</DropdownItem>
                        <DropdownItem href="/addscaffolding">Legg til stillasdel</DropdownItem>
                    </DropdownButton>
                    <DropdownButton id="dropdown-button"
                                    title= {"Bruker"}
                                    size="sm"

                    >
                        <DropdownItem href="#/action-1">Bruker Informasjon</DropdownItem>
                        <DropdownItem onClick={logOut}>Logg ut</DropdownItem>
                    </DropdownButton>
                </Toolbar>
            </AppBar>
        );
    }







}

export default TopBar;
