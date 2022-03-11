import React from "react";
import {Card, CardContent, CardHeader, CardMedia, Typography} from "@material-ui/core";

class PreView extends React.Component {
    render() {
        return (
         <div>
             <Card elevation={5} >
                 <CardHeader
                     title="Spire (3m)"
                     subheader="Test"
                 />
                 <CardMedia>
                     <img src="AK_SPI_0_01100.png" alt="recipe thumbnail"/>
                 </CardMedia>
                 <CardContent>
                     <Typography className="text">
                         Her skal det så detaljer om prosjektet
                     </Typography>
                 </CardContent>
             </Card>
          </div>
        )
    }
}

export default PreView
