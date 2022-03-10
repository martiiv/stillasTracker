import React from "react";
import {Card, CardContent, CardHeader, Typography} from "@material-ui/core";


class PreView extends React.Component {
    render() {
        return (
         <div>
             <Card elevation={5} >
                 <CardHeader
                     title="Hello World"
                     subheader="Test"
                 />
                 <CardContent>
                     <Typography>
                         Her skal det så detaljer om prosjektet
                     </Typography>
                 </CardContent>
             </Card>
          </div>
        )
    }
}

export default PreView
