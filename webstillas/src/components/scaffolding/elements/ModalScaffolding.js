import React, {useState} from "react";
import {Button, Modal} from "react-bootstrap";
import img from "../images/spirstillas_solideq_spir_klasse_5_stillas_135_1.jpg";
import {Link} from "react-router-dom";
import {useQueryClient} from 'react-query'
import {GetDummyData} from "../../../modelData/addData";
import {PROJECTS_WITH_SCAFFOLDING_URL} from "../../../modelData/constantsFile";
import "./Modalscaffolding.css"


/**
 * Function will return information about quantity of scaffolding in a specific project, including the end date of the project.
 *
 * @param type of scaffolding, the user would like more information about
 * @returns {JSX.Element|*}
 */
function ScaffoldingInProject(type) {

    const {isLoading, data} = GetDummyData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)
    if (isLoading){
        return <h1>Loading</h1>
    }else {
        const allProjects = JSON.parse(data.text)
        const result = allProjects.map((element) => {
            return {
                ...element, scaffolding: element.scaffolding.filter((subElement) =>
                    subElement.type.toLowerCase() === type.toLowerCase() && subElement.Quantity.expected !== 0)
            }
        })
        const results = result.filter(element => Object.keys(element.scaffolding).length !== 0)

        return (
            results.map(e => {
                return (
                    <div key={e.projectID} className={"card-scaffolding"}>
                        <div className={"img-and-name"}>
                            <h3>{e.projectName.toUpperCase()}</h3>
                            <img className={"img"} src={img} alt={""}/>
                        </div>
                        <div className={"list-and-btn"}>
                            <ul className={"information-list"}>
                                <li className={"horizontal-list"}>
                                    <div className={"highlightText"}>
                                        <span>{e.scaffolding[0].Quantity.expected}</span>
                                    </div>
                                    <div className={"highlightText-caption"}>
                                        <span>Expected</span>
                                    </div>
                                </li>
                                <li className={"horizontal-list"}>
                                    <div className={"highlightText"}>
                                        <span>{e.period.endDate}</span>
                                    </div>
                                    <div className={"highlightText-caption"}>
                                        <span>Return date</span>
                                    </div>
                                </li>
                            </ul>
                            <div className={"scaffolding-btn"}>
                                <Link className={"btn"} to={"/project/" + e.projectID}>Mer infromasjon</Link>
                            </div>
                        </div>
                        <hr/>

                    </div>
                )
                }
            )
        )
    }
}


/**
 * Function will display a Modal, of with information of the projects that has the selected scaffolding types.
 *
 * @param props is type of scaffolding.
 * @returns {JSX.Element}
 */
export default function InfoModal(props) {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);
    //https://codesandbox.io/s/react-week-date-view-forked-ruxjr9?file=/src/App.js:857-868
    //todo gjør om variablenavn
    return (
        <div>
            <Button className="nextButton" onClick={handleShow}>
                Vis detaljer
            </Button>

            <Modal show={show}
                   onHide={handleClose}
                   centered
                   dialogClassName="modal-dialog modal-xl"
            >
                <Modal.Header closeButton>
                    <Modal.Title>
                        {props.type}
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    {
                        ScaffoldingInProject(props.type, props.project)
                    }
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleClose}>
                        Close
                    </Button>
                </Modal.Footer>
            </Modal>
        </div>
    );
}
