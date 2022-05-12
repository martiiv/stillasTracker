import React from "react";
import { auth } from "../../firebase";
import {GetDummyData} from "../../modelData/addData";
import { USER_URL} from "../../modelData/constantsFile";
import {SpinnerDefault} from "../Spinner";
import "./userInfo.css"
import profileImg from "./profile-png-icon-2.png"
import {InternalServerError} from "../error/error";


/**
 * Function that will return information about the user.
 * @returns {JSX.Element}
 * @constructor
 */
export function UserInfo(){
    //Todo add if error

    let isLoadingUser, userData, isErrorUser

    //If user is authenticated load user data
    if (auth.currentUser){
        const {isLoading, data, isError} = GetDummyData("user", USER_URL + auth.currentUser.uid)
        isLoadingUser = isLoading
        userData = data
        isErrorUser = isError
    }

    if (isLoadingUser) {
        return (<SpinnerDefault/>)
    } else if( isErrorUser){
        return <InternalServerError />
    }
    else {
        const user = JSON.parse(userData.text)
        return (
            <div className={"main-userinfo"}>
                <div className={"info-card"}>
                    <div className={"image-frame"}>
                        <img src={profileImg} alt={""} className={"profile-image"}/>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {user?.name.firstName} {user?.name.lastName}
                        </h4>
                        <h4 className={"under-information"}>
                            Navn
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {user?.phone}
                        </h4>
                        <h4 className={"under-information"}>
                            Telefonnummer
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {user?.email}
                        </h4>
                        <h4 className={"under-information"}>
                            Email
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {user?.employeeID}
                        </h4>
                        <h4 className={"under-information"}>
                            Ansatt ID
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {user?.dateOfBirth}
                        </h4>
                        <h4 className={"under-information"}>
                            Fødselsdato
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {user?.role}
                        </h4>
                        <h4 className={"under-information"}>
                            Stilling
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {user?.admin.toString()}
                        </h4>
                        <h4 className={"under-information"}>
                            Administrerende rettigheter
                        </h4>
                    </div>
                </div>
            </div>

        )
    }
}


