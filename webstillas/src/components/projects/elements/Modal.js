import React, {useState} from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import {Modal, Button, Dropdown} from 'react-bootstrap';
import Form from 'react-bootstrap/Form'
import img from "../../scaffolding/images/spirstillas_solideq_spir_klasse_5_stillas_135_1.jpg";

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




export default function InfoModal() {
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


    const handleroom = (e, id) => {
        let result = [...roomRent];
        result = result.map((x) => {
            if (x.type.toLowerCase() === id.toLowerCase()) {
                const inputvalue = (e.target.value)
                const intValue = parseInt(inputvalue, 10)
                x.quantity = intValue
                return x;
            } else return x;
        });
        setRoomRent(result)
    };

    //todo add a note to the user if the transaction was a success or a fail.
    const addScaffolding = (body) =>{
        const requestOptions = {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(body)
        };

        fetch('http://localhost:8080/stillastracking/v1/api/project/scaffolding', requestOptions)
            .then(response => response.json())
            .catch(error => console.log(error))
        handleClose()
    }


    const move = {
        "toProjectID": Number(ToProject),
        "fromProjectID": Number(FromProject),
        "scaffold": roomRent
    }



    console.log(move)


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
                            <Form.Select value={ToProject} onChange={(e) => setToProject(e.target.value)}>
                                {jsonProjects.map(e =>{
                                    return(
                                        <option value={e.projectID}>{e.projectID}</option>
                                    )
                                })}
                            </Form.Select>
                        </div>
                        <div>
                            <span>Overfør fra prosjekt:</span>
                            <Form.Select value={FromProject} onChange={(e) => setFromProject(e.target.value)}>
                                {jsonProjects.map(e =>{
                                    return(
                                        <option value={e.projectID}>{e.projectID}</option>
                                    )
                                })}
                            </Form.Select>
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
                                        <input type="number" key={"input" + e.type} onChange={(j) => handleroom(j, e.type)}/>
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
                    <Button variant="primary" onClick={() => addScaffolding(move)}>
                        Save Changes
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}
