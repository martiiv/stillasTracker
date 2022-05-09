import React, {useState} from "react";
import "./projects.css"
import CardElement from './elements/card'
import {Route, Routes} from "react-router-dom";
import {PROJECTS_WITH_SCAFFOLDING_URL} from "../../modelData/constantsFile";
import {GetDummyData} from "../../modelData/addData";
import {SpinnerDefault} from "../Spinner";



import 'react-date-range/dist/styles.css'; // main style file
import 'react-date-range/dist/theme/default.css'; // theme css file


import { DateRangePicker } from 'react-date-range';


import "react-dates/initialize";
//import { DateRangePicker } from "react-dates";
import "react-dates/lib/css/_datepicker.css";


/**
 Class that will create an overview of the projects
 */
export function Project(){
    const [fromSize, setFromSize] = useState(0)
    const [toSize, setToSize] = useState(0)

    const [searchName, setSearchName] = useState("")
    const [selectedOption, setSelectedOption] = useState("")


    const [startDate, setStartDate] = useState(null);
    const [endDate, setEndDate] = useState(null);
    const [focusedInput, setFocusedInput] = useState(null);
    const handleDatesChange = ({ startDate, endDate }) => {
        setStartDate(startDate);
        setEndDate(endDate);
    };



    const formatDate = (inputDate) => {
        const dateArray = inputDate.split('-')
        return new Date(dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0])
    }



    const {isLoading, data} = GetDummyData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)
    if (isLoading) {
        return(
          <SpinnerDefault />
        )
    } else {
        return (
            <div className={"main-project-window"}>

                    <div className={"main-sidebar"} >
                        <div className={"search-filter"}>
                            <p className = {"input-sorting-text"}>Status</p>
                            <select className={"form-select options"} onChange={(e) =>
                                setSelectedOption(e.target.value)}>
                                <option defaultValue="">Velg her</option>
                                <option value={"Active"}>Aktiv</option>
                                <option value={"Inactive"}>Inaktiv</option>
                                <option value={"Upcoming"}>Kommende</option>
                            </select>
                        </div>
                        <div className={"search-filter"}>
                            <p className = {"input-sorting-text"}>Prosjekt navn: </p>

                            <input
                                className={"form-control"}
                                type="text"
                                   placeholder={"Søk prosjekt navn"}
                                   value={searchName}
                                   onChange={e => {
                                       setSearchName(e.target.value)
                                   }}/>
                        </div>
                        <div className={"search-filter"}>
                            <p className = {"input-sorting-text"}>Stillsmengde: </p>
                            <div className={"search-filter size"}>
                                <input
                                    className={"form-control size-search"}
                                    type="number"
                                    placeholder={"Fra"}
                                    min={0}
                                    onWheel={(e) => e.prototype}
                                    onChange={e => setFromSize(Number(e.target.value))}
                                />
                                <input
                                    className={"form-control size-search"}
                                    type="number"
                                    placeholder={"Til"}
                                    min={0}
                                    onChange={e => {
                                        setToSize(Number(e.target.value))
                                    }}
                                />
                            </div>
                        </div>
                    </div>
                <div>
                    <div className={"projects-display"}>
                        {data.filter(data => (data.projectName.toLowerCase()).includes(searchName.toLowerCase()))
                            .filter(data => {
                                if (fromSize !== 0) {
                                    console.log(fromSize)
                                    return data.size > fromSize
                                } else {
                                    return true
                                }
                            })
                            .filter(data => {
                                console.log(startDate)
                                if (startDate !== null ) {
                                    return formatDate(data.period.startDate) >= startDate._d
                                } else {
                                    return true
                                }
                            })
                            .filter(data => {
                                if (endDate !== null) {
                                    return formatDate(data.period.endDate) <= endDate
                                } else {
                                    return true
                                }
                            })
                            .filter(data => {
                                if (toSize !== 0) {
                                    return data.size < toSize
                                } else {
                                    return true
                                }
                            })
                            .filter(data => {
                                if (!(selectedOption.length === 0) && !(selectedOption === "Velg her")) {
                                    return data.state === selectedOption
                                } else {
                                    return true
                                }
                            })
                            .map((e) => {
                                return (
                                    <div key={e.projectID}>
                                        <Routes>
                                            <Route path="/project/:id" element={<CardElement data={e}/>}/>
                                        </Routes>
                                        <CardElement key={e.projectID}
                                                     id={e.projectID}
                                                     name={e.projectName}
                                                     state={e.state}
                                                     rentPeriod={e.period.startDate}
                                                     size={e.size}
                                                     contactPerson={e.customer.name}
                                                     contactNumber={e.customer.number}
                                                     address_Street={e.address.street}
                                                     address_Municipality={e.address.municipality}
                                                     address_zip={e.address.zipcode}
                                        />
                                    </div>
                                );
                            })}

                    </div>
                </div>
            </div>

        );
    }
}



