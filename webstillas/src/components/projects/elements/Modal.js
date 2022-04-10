import React, {useState} from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import {Modal, Button, Dropdown} from 'react-bootstrap';
import Form from 'react-bootstrap/Form'
import img from "../../scaffolding/images/spirstillas_solideq_spir_klasse_5_stillas_135_1.jpg";

//https://ordinarycoders.com/blog/article/react-bootstrap-modal

function scaffoldingSelection(){
    const projects = sessionStorage.getItem('project')
    const jsonProjects = JSON.parse(projects)


    return(
        <div className={"scaffoldingElement"}>
            {jsonProjects.scaffolding.map(e => {
                return(
                        <article className={"card"}>
                            <section className={"header"}>
                                <h3>{e.type.toUpperCase()}</h3>
                            </section>
                <section className={"image"}>
                        <img className={"img"} src={img} alt={""}/>
                    </section>
                            <input type={"number"} className={"input-transfer-scaffolding"}/>
                    </article>
                        )
            }
            )}

        </div>
    )
}


export default function InfoModal() {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);
    const projects = sessionStorage.getItem('allProjects')
    const jsonProjects = JSON.parse(projects)
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
                    <div>
                        <span>Overfør til prosjekt:</span>
                        <Form.Select>
                            {jsonProjects.map(e =>{
                                return(
                                    <option>{e.projectID}</option>
                                )
                            })}
                        </Form.Select>
                    </div>
                    <div>
                        <span>Overfør fra prosjekt:</span>
                        <Form.Select>
                            {jsonProjects.map(e =>{
                                return(
                                    <option>{e.projectID}</option>
                                )
                            })}
                        </Form.Select>
                    </div>
                    {scaffoldingSelection()}
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="primary" onClick={handleClose}>
                        Save Changes
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}
