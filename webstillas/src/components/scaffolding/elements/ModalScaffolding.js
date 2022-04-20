import React, {useState} from "react";
import {Button, Modal} from "react-bootstrap";
import img from "../images/spirstillas_solideq_spir_klasse_5_stillas_135_1.jpg";
import {Link} from "react-router-dom";


function scaffoldingInProject(type, props){
    const projects = sessionStorage.getItem('allProjects')
    const jsonProjects = JSON.parse(projects)

   const result = jsonProjects.map((element) => {
         return {...element, scaffolding: element.scaffolding.filter((subElement) =>
                subElement.type.toLowerCase() === type.toLowerCase() && subElement.Quantity.expected !== 0)}
    })

    //todo mulig legge til sortering??
    const results = result.filter(element => Object.keys(element.scaffolding).length !== 0)
    return(
        results.map(e =>{
            return(
                <article key={e.projectID} className={"project-card-long"}>
                        <section className={"header"}>
                            <h3>{e.projectName.toUpperCase()}</h3>
                        </section>
                        <div className={"main-body-project-card"}>
                            <section  className={"information-highlights-cta"}>
                                <div  className={"information-highlights"}>
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
                                </div>
                            </section>
                            <div>
                                <section className={"image"}>
                                    <img className={"img"} src={img} alt={""}/>
                                </section>
                                <section className={"card-btn"}>
                                    <div className={"card-btns"}>
                                        <Link className={"btn"} to={"/project/" + e.projectID}>Mer Informasjon</Link>
                                    </div>
                                </section>
                            </div>
                        </div>
                    <hr/>

                    </article>

            )
        })
    )
}







export default function InfoModal(props) {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);
    //https://codesandbox.io/s/react-week-date-view-forked-ruxjr9?file=/src/App.js:857-868
    //todo gjør om variablenavn
    return (
        <>
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
                    {scaffoldingInProject(props.type, props)}
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleClose}>
                        Close
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}
