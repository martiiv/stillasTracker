import React from "react";
import "./scaffolding.css"
import CardElement from "./elements/scaffoldingCard";
import fetchModel from "../../modelData/fetchData";
import {
    PROJECTS_URL_WITH_SCAFFOLDING,
    PROJECTS_WITH_SCAFFOLDING_URL,
    SCAFFOLDING_URL,
    STORAGE_URL
} from "../../modelData/constantsFile";
import {GetDummyData} from "../../modelData/addData";

/**
 Class that will create an overview of the scaffolding parts
 */



class ScaffoldingClass extends React.Component {
    constructor(props) {
        super(props);
        this.state={
            isLoaded1: false,
            isLoaded2: false,
            scaffolding: props.scaffolding,
            storage:props.storage,
            items: [],
            selectedOption: ""
        }
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
      const {scaffolding, storage, isLoaded1,isLoaded2, selectedOption } = this.state;

      let scaffoldingArray
      if (sessionStorage.getItem('allScaffolding') != null){
          const scaffold = sessionStorage.getItem('allScaffolding')
          console.log('From Storage')
          scaffoldingArray = (JSON.parse(scaffold))
      }else {
          console.log('From API')
          scaffoldingArray = scaffolding
      }

      const objectArr = this.countObjects(scaffoldingArray, "type")

      let storageArray
      if (sessionStorage.getItem('fromStorage') != null){
          const storage = sessionStorage.getItem('fromStorage')
          storageArray = (JSON.parse(storage))
      }else {
          console.log('From API')
          storageArray = storage
      }
      const scaffoldingObject = this.scaffoldingAndStorage(objectArr, storageArray)

      const result = Object.keys(scaffoldingObject).map((key) => scaffoldingObject[key]);
      if (!isLoaded1 && !isLoaded2) {
          return <h1>Is Loading Data....</h1>
      } else {
          if (selectedOption === "ascending") {
              result[0].sort((a, b) => (a.scaffolding < b.scaffolding) ? 1 : -1)
          }else if (selectedOption === "descending") {
              result[0].sort((a, b) => (a.scaffolding > b.scaffolding) ? 1 : -1)
          }else {
              result[0].sort((a, b) => (a.type > b.type))
          }
          return (
              //todo only scroll the scaffolding not the map
              <div>
                  <div>
                      <select onChange={(e) =>
                          this.setState({selectedOption: e.target.value})}>
                          <option value={"alphabetic"}>Alfabetisk(A-Å)</option>
                          <option value={"ascending"}>Stigende</option>
                          <option value={"descending"}>Synkende</option>
                      </select>
                      <p>Sorter</p>
                  </div>



                 <div className={"grid-container"}>
                      {result[0].map((e) => {
                          return (
                              <CardElement key={e.type}
                                           type={e.type}
                                           total={e.scaffolding}
                                           storage={e.storage}
                              />
                          )
                      })}
                  </div>
              </div>

          )
      }
  }
}

export const Scaffolding = () => {
    const {isLoading: LoadingScaffolding, data: Scaffolding} = GetDummyData("scaffolding", SCAFFOLDING_URL)
    const {isLoading: LoadingStorage, data: Storage} = GetDummyData("storage", STORAGE_URL)

    if (isLoading) {
        return <h1>Loading</h1>
    } else {
        return <ScaffoldingClass scaffolding = {Scaffolding}
                                 storage = {Storage} />
    }
}
