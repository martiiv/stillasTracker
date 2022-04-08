import React from "react";
import "./scaffolding.css"
import CardElement from "./elements/scaffoldingCard";

/**
 Class that will create an overview of the scaffolding parts
 */

class Scaffolding extends React.Component {
    constructor(props) {
        super(props);
        this.state={
            isLoaded1: false,
            isLoaded2: false,
            scaffolding: [],
            storage:[],
            items: []
        }
    }

    async componentDidMount() {
        const urlScaffolding ="http://10.212.138.205:8080/stillastracking/v1/api/unit/";
        fetch(urlScaffolding)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded1: true,
                        scaffolding: result
                    });
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    this.setState({
                        isLoaded1: true,
                    });
                }
                )

        const urlStorage ="http://10.212.138.205:8080/stillastracking/v1/api/storage/";
        fetch(urlStorage)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded2: true,
                        storage: result
                    });
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    this.setState({
                        isLoaded2: true,
                    });
                }
            )
    }


    countObjects(arr, key){
        let arr2 = [];

        arr.forEach((x)=>{
            // Checking if there is any object in arr2
            // which contains the key value
            if(arr2.some((val)=>{ return val[key] === x[key] })){

                // If yes! then increase the occurrence by 1
                arr2.forEach((k)=>{
                    if(k[key] === x[key]){
                        k["occurrence"]++
                    }
                })

            }else{
                // If not! Then create a new object initialize
                // it with the present iteration key's value and
                // set the occurrence to 1
                let a = {}
                a[key] = x[key]
                a["occurrence"] = 1
                arr2.push(a);
            }
        })

        return arr2
    }


    scaffoldingAndStorage(scaffold, storage){
        const scaffoldVar = {
            scaffolding: []
        };


        for(var i in scaffold) {
            var scaff = scaffold[i];
            for (var j in storage){
                var stor = storage[j];

                if (stor.type.toLowerCase() === scaff.type.toLowerCase()){
                    scaffoldVar.scaffolding.push({
                        "type"          :scaff.type,
                        "scaffolding"   :scaff.occurrence,
                        "storage"       :stor.Quantity.expected
                    });
                }
            }

        }
        return scaffoldVar
    }




  render() {
      const {scaffolding, storage, isLoaded1,isLoaded2 } = this.state;
      const objectArr = this.countObjects(scaffolding, "type")
      const scaffoldingObject = this.scaffoldingAndStorage(objectArr, storage)
      const result = Object.keys(scaffoldingObject).map((key) => scaffoldingObject[key]);

      if (!isLoaded1 && !isLoaded2) {
          return <h1>Is Loading Data....</h1>
      } else {
          return (
              <div className={"grid-container"}>
                  {result[0].map((e) => {
                      console.log(e)
                      return (
                          <CardElement key={e.type}
                                       type={e.type}
                                       total={e.scaffolding}
                                       storage={e.storage}
                          />

                      )
                  })}
              </div>
          )
      }
  }
}

export default Scaffolding;
