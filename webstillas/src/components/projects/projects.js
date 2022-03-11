import React from "react";
import "./projects.css"
import PreView from "./elements/preView"
import {Container, Grid} from "@material-ui/core";
/**
 Class that will create an overview of the projects
 */

class Projects extends React.Component {
    render() {
        return(
            <Container>
                <Grid lg={2} xs={25} md={30}>
                    <PreView />
                </Grid>
¨            </Container>
        );
    }
}

export default Projects;
