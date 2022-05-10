import React from "react";
import './topBar.css';
import {
    AppBar, Toolbar, Button
} from '@material-ui/core';
import {Link, NavLink} from "react-router-dom";
import {DropdownButton, NavDropdown} from "react-bootstrap";
import DropdownItem from "react-bootstrap/DropdownItem";
import {useUserAuth} from "../../context/UserAuthContext";
import {auth} from "../../firebase"

/**
 Class that will create a topbar for the application.
 */

//Todo make list instead of toolbar
    //se hva andre nettsider har gjort
const TopBar = () => {
        const {logOut} = useUserAuth();


        if (!auth.currentUser) {
            return (
                <AppBar position="sticky">
                    <Toolbar className="toolbar">

                    </Toolbar>
                </AppBar>
            )
        } else {
            return (
                <AppBar position="sticky">
                    <Toolbar className="toolbar">
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

                        <NavDropdown id="basic-nav-dropdown"
                                        title={"Logistikk"}
                                        size="sm"
                                     style={{textDecorationColor: "black"}}


                        >
                            <DropdownItem>
                                <Link to={"/addproject/"}>Legg til prosjekt </Link>
                            </DropdownItem>
                            <DropdownItem>
                                <Link to={"/addscaffolding/"}>Legg til stillas</Link>
                            </DropdownItem>


                        </NavDropdown>
                        <DropdownButton id="dropdown-button"
                                        title={"Bruker"}
                                        size="sm"

                        >
                            <DropdownItem>
                                <Link to={"/addscaffolding/"}>Bruker Informasjon</Link>
                            </DropdownItem>

                            <DropdownItem onClick={logOut}>Logg ut</DropdownItem>
                        </DropdownButton>
                    </Toolbar>
                </AppBar>
            );
        }


    }

export default TopBar;
