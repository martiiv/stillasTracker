import React, {useState} from "react";
import "./projects.css"
import CardElement from './elements/card'
import {Route, Routes} from "react-router-dom";
import {PROJECTS_WITH_SCAFFOLDING_URL} from "../../modelData/constantsFile";
import {GetDummyData} from "../../modelData/addData";
import {SpinnerDefault} from "../Spinner";

/**
 Class that will create an overview of the projects
 */
export function Project(){
    const [fromSize, setFromSize] = useState(0)
    const [toSize, setToSize] = useState(0)
    const [fromDate, setFromDate] = useState("")
    const [toDate, setToDate] = useState("")
    const [searchName, setSearchName] = useState("")
    const [selectedOption, setSelectedOption] = useState("")
    const reverseDate = (inputDate) => {
        const dateArray = inputDate.split('-')
        return dateArray[2] + '-' + dateArray[1] + '-' + dateArray[0]
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
                    <div>
                        <select onChange={(e) =>
                            setSelectedOption(e.target.value)}>
                            <option defaultValue="">Velg her</option>
                            <option value={"Active"}>Aktiv</option>
                            <option value={"Inactive"}>Inaktiv</option>
                            <option value={"Upcoming"}>Kommende</option>
                        </select>
                        <p>Status</p>

                    </div>
                    <form className={"filter-content-search"}>
                        <p>Prosjekt Navn</p>
                        <input type="text"
                               placeholder={"Søk prosjekt navn"}
                               value={searchName}
                               onChange={e => {
                                   setSearchName(e.target.value)
                               }}/>
                    </form>
                    <form className={"filter-content-input"}>
                        <p>Stillsmengde: </p>
                        <input type="number" placeholder={"Fra"}
                               min={0}
                               onWheel={(e) => e.prototype}
                               onChange={e => setFromSize(Number(e.target.value))}
                               className={"input-fieldNumber-filter"}/>
                        <input type="number" placeholder={"Til"}
                               min={0}
                               onChange={e => {
                                   setToSize(Number(e.target.value))
                               }}
                               className={"input-fieldNumber-filter"}/>
                    </form>
                    <form className={"filter-content-input"}>
                        <p>Tidsperiode: </p>
                        <input type="date" value={fromDate} onChange={e => setFromDate(e.target.value)}
                               className={"input-field-filter"}/>
                        <input type="date" value={toDate} onChange={e => setToDate(e.target.value)}
                               className={"input-field-filter"}/>
                    </form>
                </div>
                <div className={"grid-container"}>
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
                            if (fromDate !== "") {
                                return reverseDate(data.period.startDate) >= fromDate
                            } else {
                                return true
                            }
                        })
                        .filter(data => {
                            if (toDate !== "") {
                                return reverseDate(data.period.endDate) <= toDate
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

        );
    }
}



