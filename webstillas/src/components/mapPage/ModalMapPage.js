import React, {useState} from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import {Button, Modal} from 'react-bootstrap';
import img from "../../scaffolding/images/spirstillas_solideq_spir_klasse_5_stillas_135_1.jpg";


export default function MapModal() {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);
    //https://codesandbox.io/s/react-week-date-view-forked-ruxjr9?file=/src/App.js:857-868
    //todo gjør om variablenavn
    const projects = sessionStorage.getItem('allProjects')
    const jsonProjects = JSON.parse(projects)
    const project = sessionStorage.getItem('project')
    const jsonProject = JSON.parse(project)
    const [roomRent, setRoomRent] = useState(scaffoldingMove);
    const [ToProject, setToProject] = useState("");
    const [FromProject, setFromProject] = useState("");

    return (
        <>
            <Button className="nextButton" onClick={handleShow}>
                Overfør deler til Prosjekt
            </Button>

            <Modal show={show}
                   onHide={handleClose}
                   centered
                   backdrop="static"
                   dialogClassName="modal-dialog modal-xl"
            >
                <Modal.Header closeButton>
                    <Modal.Title>Stillas Overføring</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <div className={"scaffoldingElement"}>
                        <div>
                            <span>Overfør til prosjekt:</span>
                            <select value={ToProject} onChange={(e) => setToProject(e.target.value)}>
                                <option selected defaultValue="">Choose here</option>
                                {jsonProjects.map(e =>{
                                    return(
                                        <option value={e.projectID}>{e.projectID}</option>
                                    )
                                })}
                            </select>
                        </div>
                        <div>
                            <span>Overfør fra prosjekt:</span>
                            <select value={FromProject}
                                    onChange={(e) => setFromProject(e.target.value)}>
                                <option selected defaultValue="">Choose here</option>
                                {jsonProjects.map(e =>{
                                    return(
                                        <option value={e.projectID}>{e.projectID}</option>
                                    )
                                })}
                            </select>
                        </div>
                        {jsonProject.scaffolding.map(e => {
                                return(
                                    <article className={"card"}>
                                        <section className={"header"}>
                                            <h3>{e.type.toUpperCase()}</h3>
                                        </section>
                                        <section className={"image"}>
                                            <img className={"img"} src={img} alt={""}/>
                                        </section>
                                        <input type="number" min={0} key={"input" + e.type} onChange={(j) => handleroom(j, e.type)}/>
                                    </article>
                                )
                            }
                        )}

                    </div>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="primary" disabled={!validFormat} onClick={() => AddScaffolding(move)}>
                        Save Changes
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}
