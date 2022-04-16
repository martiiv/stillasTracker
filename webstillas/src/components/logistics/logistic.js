import React from "react";
import AddProject from "./project/addProject";
import AddScaffolding from './scaffold/addScaffolding'
import AddUser from "./user/addUser";

class Logistic extends React.Component{
    constructor(props) {
        super(props);

    }

    render() {
        return(
            <AddUser />
        )
    }

}

export default Logistic
