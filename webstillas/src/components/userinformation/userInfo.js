import React from "react";
import { auth } from "../../firebase";
import {GetDummyData} from "../../modelData/addData";
import { USER_URL} from "../../modelData/constantsFile";
import {SpinnerDefault} from "../Spinner";
import "./userInfo.css"
import profileImg from "./profile-png-icon-2.png"

export function UserInfo(){
    const {isLoading, data} = GetDummyData("user", USER_URL + auth.currentUser.uid)

    if (isLoading) {
        return (<SpinnerDefault/>)
    } else {
        return (
            <div className={"main-userinfo"}>
                <div className={"info-card"}>
                    <div className={"image-frame"}>
                        <img src={profileImg} alt={""} className={"profile-image"}/>

                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {data.name.firstName} {data.name.lastName}
                        </h4>
                        <h4 className={"under-information"}>
                            Navn
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {data.phone}
                        </h4>
                        <h4 className={"under-information"}>
                            Telefonnummer
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {data.email}
                        </h4>
                        <h4 className={"under-information"}>
                            Email
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {data.employeeID}
                        </h4>
                        <h4 className={"under-information"}>
                            Ansatt ID
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {data.dateOfBirth}
                        </h4>
                        <h4 className={"under-information"}>
                            Fødselsdato
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {data.role}
                        </h4>
                        <h4 className={"under-information"}>
                            Stilling
                        </h4>
                    </div>
                    <div className={"information-text"}>
                        <h4 className={"header-information"}>
                            {data.admin.toString()}
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


