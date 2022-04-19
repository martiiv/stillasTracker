import React from "react";
import AddProject from "./project/addProject";
import AddScaffolding from './scaffold/addScaffolding'
import AddUser from "./user/addUser";
import Tabs from "../projects/tabView/Tabs";

class Logistic extends React.Component{
    constructor(props) {
        super(props);

    }

    render() {
        return(
            <Tabs>
                <div label="Legg til Prosjekt">
                    <AddProject />
                </div>
                <div label="Legg til Bruker">
                    <AddUser />
                </div>
                <div label="Leggt til Stillasdel ">
                    <AddScaffolding />
                </div>
            </Tabs>

        )
    }

}

export default Logistic
