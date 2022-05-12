import React from "react";
import './topBar.css';
import {
    AppBar, Toolbar, Button
} from '@material-ui/core';
import {Link} from "react-router-dom";
import {DropdownButton, NavDropdown} from "react-bootstrap";
import DropdownItem from "react-bootstrap/DropdownItem";
import {useUserAuth} from "../../context/UserAuthContext";
import {auth} from "../../firebase"
import {GetDummyData} from "../../modelData/addData";
import {USER_URL} from "../../modelData/constantsFile";
import {SpinnerDefault} from "../Spinner";
import "bootstrap/dist/css/bootstrap.min.css";
import {ADD_PROJECT_URL, ADD_SCAFFOLDING_URL, MAP_URL, PROJECT_URL, SCAFFOLDING_URL, USERINFO_URL} from "../constants";

/**
 Component that will be used as a top bar for the user to navigate throughout the application.
 */

const TopBar = () => {
    const {logOut} = useUserAuth();

    let loading, user


    //If the user is authenticated, fetch data from database
    if (auth.currentUser){
        const {isLoading, data} = GetDummyData("user", USER_URL + auth.currentUser.uid)
        loading = isLoading
        user = data
    }

    /*
    If the user is not authenticated, the topbar will be empty.
     */
    if (!auth.currentUser) {
        return (
            <AppBar position="sticky">
                <Toolbar className="toolbar">
                </Toolbar>
            </AppBar>
        )
    } else if (loading) {
        //If data is loading, the user will get a spinner displayed
        return <SpinnerDefault/>
    } else {
        const userData = JSON.parse(user.text)
        //Top bar with interactive buttons to navigate.
        return (
            <AppBar position="sticky">
                <Toolbar className="toolbar">
                    <Link className="link" to={PROJECT_URL}>
                        <Button className="button">Prosjekter</Button>
                    </Link>
                    <Link className="link" to={SCAFFOLDING_URL}>
                        <Button className="button">Stillasdeler</Button>
                    </Link>
                    <Link className="link" to={MAP_URL}>
                        <Button className="button">Kart</Button>
                    </Link>

                    <NavDropdown id="basic-nav-dropdown1"
                                 title={"Logistikk"}
                                 size="sm"
                                 className={"dropDownMenu"}
                    >
                        <DropdownItem>
                            <Link className={"link"}
                                to={ADD_PROJECT_URL}>Legg til prosjekt </Link>
                        </DropdownItem>
                        <DropdownItem>
                            <Link className={"link"}
                                to={ADD_SCAFFOLDING_URL}>Legg til stillas</Link>
                        </DropdownItem>
                    </NavDropdown>
                    <DropdownButton id="dropdown-button"
                                    title={userData?.name.firstName}
                                    size="sm"
                                    className={"dropDownMenu"}
                    >
                        <DropdownItem>
                            <Link className={"link"}
                                  to={USERINFO_URL}>
                                Bruker Informasjon</Link>
                        </DropdownItem>
                        <DropdownItem onClick={logOut}>Logg ut</DropdownItem>
                    </DropdownButton>
                </Toolbar>
            </AppBar>
        );
    }
}

export default TopBar;
