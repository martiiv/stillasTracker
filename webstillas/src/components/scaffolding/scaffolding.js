import React from "react";
import "./scaffolding.css"
import CardElement from "./elements/scaffoldingCard";
import {PROJECTS_WITH_SCAFFOLDING_URL, SCAFFOLDING_URL, STORAGE_URL} from "../../modelData/constantsFile";
import {GetDummyData} from "../projects/getDummyData";
import {useQueryClient} from "react-query";
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
        const {scaffolding, storage, selectedOption} = this.state;

        const objectArr = this.countObjects(scaffolding, "type")


        const scaffoldingObject = this.scaffoldingAndStorage(objectArr, storage)

        const result = Object.keys(scaffoldingObject).map((key) => scaffoldingObject[key]);

        if (selectedOption === "ascending") {
            result[0].sort((a, b) => (a.scaffolding < b.scaffolding) ? 1 : -1)
        } else if (selectedOption === "descending") {
            result[0].sort((a, b) => (a.scaffolding > b.scaffolding) ? 1 : -1)
        } else {
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
                                         projects = {this.props.projects}
                            />
                        )
                    })}
                </div>
            </div>

        )
    }

}

export const Scaffolding = () => {
    const {isLoading: LoadingScaffolding, data: Scaffolding} = GetDummyData("scaffolding", SCAFFOLDING_URL)
    const {isLoading: LoadingStorage, data: Storage} = GetDummyData("storage", STORAGE_URL)
    const queryClient = useQueryClient()
    const dataProjects = queryClient.getQueryData("allProjects")
    let LoadingAll
    let ProjectsData
    if(dataProjects === undefined){
        const {isLoading: LoadingAllProjects, data: ProjectData} = GetDummyData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)
        LoadingAll = LoadingAllProjects
        ProjectsData = ProjectData
    }

    if (LoadingScaffolding || LoadingStorage || LoadingAll) {
        return <h1>Loading</h1>
    } else {
        return <ScaffoldingClass scaffolding = {Scaffolding}
                                 storage = {Storage}
                                 projects = {ProjectsData}
        />
    }
}
