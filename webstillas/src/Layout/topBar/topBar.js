import React from "react";
import '../../Assets/Styles/topBar.css';
import {
    AppBar, Toolbar, Button
} from '@material-ui/core';
import {Link} from "react-router-dom";
import {Dropdown } from "react-bootstrap";
import DropdownItem from "react-bootstrap/DropdownItem";
import {useUserAuth} from "../../Config/UserAuthContext";
import {auth} from "../../Config/firebase"
import {GetCachingData} from "../../Middleware/addData";
import {USER_URL} from "../../Constants/apiURL";
import {SpinnerDefault} from "../../components/Indicators/Spinner";
import "bootstrap/dist/css/bootstrap.min.css";
import {ADD_PROJECT_URL, ADD_SCAFFOLDING_URL, MAP_URL, PROJECT_URL, SCAFFOLDING_URL, USERINFO_URL} from "../../Constants/webURL";
import DropdownToggle from "react-bootstrap/DropdownToggle";
import DropdownMenu from "react-bootstrap/DropdownMenu";
import profileImg from "../../Assets/Images/profile-png-icon-2.png"

/**
 Component that will be used as a top bar for the user to navigate throughout the application.
 */

const TopBar = () => {
    const {logOut} = useUserAuth();

    let loading, user


    //If the user is authenticated, fetch data from database
    if (auth.currentUser){
        const {isLoading, data} = GetCachingData("user", USER_URL + auth.currentUser.uid)
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

                    <Dropdown>
                        <DropdownToggle className={"dropdown-toggle-topbar"}  variant=" primary" id="dropdown-basic">
                            Logistikk
                        </DropdownToggle>
                        <DropdownMenu>
                            <DropdownItem className={"dropdown-item-topbar"}>
                                <Link className={"dropdown-item"}
                                      to={ADD_PROJECT_URL}>Legg til prosjekt </Link>
                            </DropdownItem>
                            <DropdownItem>
                                <Link className={"dropdown-item"}
                                      to={ADD_SCAFFOLDING_URL}>Legg til stillas</Link>
                            </DropdownItem>
                        </DropdownMenu>
                    </Dropdown>

                    <Dropdown>
                        <DropdownToggle className={"dropdown-toggle-topbar user-button-topbar"} variant=" primary"  id="dropdown-basic">
                            <img  src={profileImg} alt={""} style={{height: "30px"}}/>
                            <p>{userData?.name.firstName}</p>
                        </DropdownToggle>
                        <DropdownMenu >
                            <DropdownItem >
                                <Link className={"dropdown-item"}
                                      to={USERINFO_URL}>
                                    Bruker Informasjon</Link>
                            </DropdownItem>
                            <div className="dropdown-divider"></div>
                            <DropdownItem
                                className={"dropdown-item"}
                                onClick={logOut}>Logg ut</DropdownItem>
                        </DropdownMenu>
                    </Dropdown>
                </Toolbar>
            </AppBar>
        );
    }
}

export default TopBar;
