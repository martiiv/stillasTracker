import React from "react";
import "./scaffolding.css"
import CardElement from "./elements/scaffoldingCard";
import {PROJECTS_WITH_SCAFFOLDING_URL, SCAFFOLDING_URL, STORAGE_URL} from "../../modelData/constantsFile";
import {GetDummyData} from "../../modelData/addData";
import {SpinnerDefault} from "../Spinner";
import {InternalServerError} from "../error/error";

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


    /**
     * Function that will count numbers of occurrences different types of scaffolding.
     *
     * @param arr is the array we are iterating.
     * @param key is the field we want to count.
     * @returns {*[]}
     */
    countObjects(arr, key){
        let arr2 = [];
        arr?.forEach((x)=>{
            // Checking if there is any object in arr2
            // which contains the key value
            if(arr2?.some((val)=>{return val[key] === x[key]})){
                // If yes! then increase the occurrence by 1
                arr2?.forEach((k)=>{
                    if(k[key] !== x[key]){
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
                arr2?.push(a);
            }
        })

        return arr2
    }


    /**
     * Function to add occurrences of type scaffolding in desired body.
     *
     * @param scaffold count of occurrences
     * @param storage objects in an array.
     * @returns {{scaffolding: *[]}}
     */
    scaffoldingAndStorage(scaffold, storage){
        const scaffoldVar = {
            scaffolding: []
        };
        for(const scaffoldIndex of scaffold) {
            const scaff = scaffoldIndex;
            for (const storageIndex of storage){
                const stor = storageIndex;
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

        //If user would like to sort based on scaffolding
        if (selectedOption === "ascending") {
            result[0]?.sort((a, b) => (a.scaffolding < b.scaffolding) ? 1 : -1)
        } else if (selectedOption === "descending") {
            result[0]?.sort((a, b) => (a.scaffolding > b.scaffolding) ? 1 : -1)
        } else {
            result[0]?.sort((a, b) => (a.type > b.type))
        }
        return (
            <div className={"scaffolding"}>
                <div className={"all-scaffolding"}>
                    <div className={"sorting"}>
                        <p className = {"input-sorting-text"}>Sorter på:</p>
                        <select className={"form-select"} onChange={(e) =>
                            this.setState({selectedOption: e.target.value})}>
                            <option value={"alphabetic"}>Alfabetisk(A-Å)</option>
                            <option value={"ascending"}>Stigende</option>
                            <option value={"descending"}>Synkende</option>
                        </select>
                    </div>

                    <div className={"grid-container"}>
                    {result[0]?.map((e) => {

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

            </div>

        )
    }

}


/**
 * Function to display information about scaffolding
 * @returns {JSX.Element}
 * @constructor
 */
export const Scaffolding = () => {
    const {isLoading: LoadingScaffolding, data: Scaffolding, isError: scaffoldingError} = GetDummyData("scaffolding", SCAFFOLDING_URL)
    const {isLoading: LoadingStorage, data: Storage, isError: storageError} = GetDummyData("storage", STORAGE_URL)
    const {isLoading: LoadingAll, data: Project, isError: allProjectError} = GetDummyData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)



    //If loading
    if (LoadingScaffolding || LoadingStorage || LoadingAll) {
        return <SpinnerDefault />
    } else if(scaffoldingError || storageError || allProjectError) //If loading error
    {
        return <InternalServerError />
    } else { //On success
        const scaffoldingData = JSON.parse(Scaffolding.text)
        const storageData = JSON.parse(Storage.text)
        const projectData = JSON.parse(Project.text)

        return <ScaffoldingClass scaffolding = {scaffoldingData}
                                 storage = {storageData}
                                 projects = {projectData}
        />
    }
}
