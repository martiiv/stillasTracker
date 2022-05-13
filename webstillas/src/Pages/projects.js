import React, {useState} from "react";
import "../Assets/Styles/projects.css"
import CardElement from '../components/projects/mainProjectCard'
import {Route, Routes} from "react-router-dom";
import {PROJECTS_WITH_SCAFFOLDING_URL} from "../Constants/apiURL";
import {GetCachingData} from "../Middleware/addData";
import {SpinnerDefault} from "../components/Indicators/Spinner";
import {InternalServerError} from "../components/Indicators/error";


/**
 Class that will create an overview of the projects
 */
export function Project() {
    const [fromSize, setFromSize] = useState(0)
    const [toSize, setToSize] = useState(0)

    const [searchName, setSearchName] = useState("")
    const [selectedOption, setSelectedOption] = useState("")

    const [startDate, setStartDate] = useState(null);
    const [endDate, setEndDate] = useState(null);


    const {isLoading, data, isError} = GetCachingData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)


    //If data is loading
    if (isLoading) {
        return (
            <SpinnerDefault/>
        )
    } else if (isError) //If an Indicators while fetching data has occurred
        {
        return <InternalServerError/>
    } else {
        const allProjects = (JSON.parse(data.text))
        return (
            <div className={"main-project-window"}>

                <div className={"main-sidebar"}>
                    <div className={"search-filter"}>
                        <p className={"input-sorting-text"}>Status</p>
                        <select className={"form-select options"} onChange={(e) =>
                            setSelectedOption(e.target.value)}>
                            <option defaultValue="">Velg her</option>
                            <option value={"Active"}>Aktiv</option>
                            <option value={"Inactive"}>Inaktiv</option>
                            <option value={"Upcoming"}>Kommende</option>
                        </select>
                    </div>
                    <div className={"search-filter"}>
                        <p className={"input-sorting-text"}>Prosjekt navn: </p>
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
                        <p className={"input-sorting-text"}>Stillsmengde: </p>
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
                    <div className={"date-filter"}>
                        <p className={"input-sorting-text"}>Fra dato: </p>
                        <input
                            className={"form-control"}
                            type="date"
                            onChange={e => {
                                setStartDate(formatDateToString(e.target.value))
                            }}/>
                    </div>
                    <div className={"search-filter"}>
                        <p className={"input-sorting-text"}>Til dato: </p>
                        <input
                            className={"form-control"}
                            type="date"
                            onChange={e => {
                                setEndDate(formatDateToString(e.target.value))
                            }}/>
                    </div>
                </div>
                <div>
                    <div className={"projects-display"}>
                        {allProjects.filter(data => (data.projectName.toLowerCase()).includes(searchName.toLowerCase()))
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
                                if (startDate !== null) {
                                    return formatDate(data.period.startDate) >= formatDate(startDate)
                                } else {
                                    return true
                                }
                            })
                            .filter(data => {
                                if (endDate !== null) {
                                    return formatDate(data.period.endDate) <= formatDate(endDate)
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


/**
 * Function to format date from "mm-dd-yyyy" to "dd-mm-yyyy"
 * @param inputDate in format "mm-dd-yyyy"
 * @returns {Date} in format "dd-mm-yyyy"
 */
export function formatDate(inputDate) {
    const dateArray = inputDate.split('-')
    return new Date(dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0])
}

/**
 * Function to format date from "mm-dd-yyyy" to "dd-mm-yyyy"
 * @param inputDate in format "mm-dd-yyyy"
 * @returns {string} in format "dd-mm-yyyy"
 */
export function formatDateToString(inputDate) {
    const dateArray = inputDate.split('-')
    return (dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0])
}
