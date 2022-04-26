import React, {useState} from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import {Button, Modal} from 'react-bootstrap';
import img from "../../scaffolding/images/spirstillas_solideq_spir_klasse_5_stillas_135_1.jpg";
import putModel from "../../../modelData/putData";
import {PROJECTS_URL, TRANSFER_SCAFFOLDING} from "../../../modelData/constantsFile";
import fetchModel from "../../../modelData/fetchData";
import AddData from "../../../modelData/addData";

//https://ordinarycoders.com/blog/article/react-bootstrap-modal
const scaffoldingMove =
    [
        {
            "type": "Bunnskrue",
            "quantity": 0
        },
        {
            "type": "Diagonalstang",
            "quantity": 0
        },
        {
            "type": "Enrørsbjelke",
            "quantity": 0
        },
        {
            "type": "Gelender",
            "quantity": 0
        },
        {
            "type": "Lengdebjelke",
            "quantity": 0
        },
        {
            "type": "Plank",
            "quantity": 0
        },
        {
            "type": "Rekkverksramme",
            "quantity": 0
        },
        {
            "type": "Spire",
            "quantity": 0
        },
        {
            "type": "Stillaslem",
            "quantity": 0
        },
        {
            "type": "Trapp",
            "quantity": 0
        }
        ]




export default function InfoModal(props) {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);
    const data = new AddData()
    //https://codesandbox.io/s/react-week-date-view-forked-ruxjr9?file=/src/App.js:857-868
    //todo gjør om variablenavn


    let projects
    if (sessionStorage.getItem('allProjects') !== null) {
         projects = sessionStorage.getItem('allProjects')
    }else {
         projects = fetchModel(PROJECTS_URL)
        sessionStorage.setItem('allProjects', projects)
    }



    const jsonProjects = JSON.parse(projects)
    const project = sessionStorage.getItem('project')
    let jsonProject = JSON.parse(project)
    const [roomRent, setRoomRent] = useState(scaffoldingMove);
    const [ToProject, setToProject] = useState("");
    const [FromProject, setFromProject] = useState("");


    //Todo change variable
    const handleroom = (e, id) => {
        let result = [...roomRent];
        result = result.map((x) => {
            if (x.type.toLowerCase() === id.toLowerCase()) {
                const inputvalue = (e.target.value)
                x.quantity = parseInt(inputvalue, 10)
                return x;
            } else return x;
        });
        setRoomRent(result)
    };

    //todo add a note to the user if the transaction was a success or a fail.
    const AddScaffolding = async (body, id) => {
        putModel(TRANSFER_SCAFFOLDING, JSON.stringify(body))
        handleClose()
    }


    const move = {
        "toProjectID": Number(ToProject),
        "fromProjectID": Number(FromProject),
        "scaffold": roomRent
    }


    const validFormat = ToProject !== FromProject
    return(
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
                    <Button variant="primary" disabled={!validFormat} onClick={AddScaffolding(move, props.id)}>
                        Save Changes
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}
