import React, {useState} from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import {Button, Modal, Spinner} from 'react-bootstrap';
import img from "../../scaffolding/images/spirstillas_solideq_spir_klasse_5_stillas_135_1.jpg";
import putModel from "../../../modelData/putData";
import {
    PROJECTS_URL_WITH_ID,
    PROJECTS_WITH_SCAFFOLDING_URL,
    TRANSFER_SCAFFOLDING,
    WITH_SCAFFOLDING_URL
} from "../../../modelData/constantsFile";
import {useQueryClient} from "react-query";
import "./Modal.css"
import {GetDummyData} from "../../../modelData/addData";
import {SpinnerDefault} from "../../Spinner";



//https://ordinarycoders.com/blog/article/react-bootstrap-modal

//JSON body that is used to send request.
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
            "type": "Spir",
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


/**
 * Function that will endable the user to transfer scaffolding from one location to another.
 *
 * @param props information of a given project
 * @returns {JSX.Element}
 */
export default function InfoModalFunc(props) {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);
    const queryClient = useQueryClient()

    const {isLoading, data: projects} = GetDummyData("allProjects", PROJECTS_WITH_SCAFFOLDING_URL)
    let project = queryClient.getQueryData(["project", props.id])
    let jsonProject = JSON.parse(project.text)

    const [scaffolding, setScaffolding] = useState(scaffoldingMove);
    const [ToProject, setToProject] = useState("");
    const [FromProject, setFromProject] = useState("");
    const [buttonPressed, setButtonPressed] = useState(false)

    /**
     * Function to set quantity of scaffolding types
     *
     * Code taken from https://codesandbox.io/s/react-week-date-view-forked-ruxjr9?file=/src/App.js:857-868
     *
     * @param e quantity the user has passed
     * @param id of the selected project.
     */
    const setQuantity = (e, id) => {
        let result = [...scaffolding];
        result = result.map((x) => {
            if (x.type.toLowerCase() === id.toLowerCase()) {
                const inputvalue = (e.target.value)
                x.quantity = parseInt(inputvalue, 10)
                return x;
            } else return x;
        });
        setScaffolding(result)
    };


    /**
     * Function that will execute request to transfer scaffolding parts.
     *
     * @returns {Promise<void>}
     */
    const AddScaffold = async () => {
        setButtonPressed(true)
        try {
            const adding = await putModel(TRANSFER_SCAFFOLDING, JSON.stringify(move))
            console.log(adding)
            await queryClient.resetQueries(["project", props.id])
        } catch (e) {
            if (e.text === "invalid body"){
                window.alert("500 Internal Server Error\nNoe gikk galt! Prøv igjen senere")
            }else {
                window.alert("Advarsel: Kan ikke overføre antall stillasdeler")
            }
        }
    }

    //JSON body that is sent with request
    const move = {
        "toProjectID": Number(ToProject),
        "fromProjectID": Number(FromProject),
        "scaffold": scaffolding
    }



    //Checks if the user did not set to project equal to from project.
    const validFormat = ToProject !== FromProject
    let jsonProjects
    if (!isLoading){
        jsonProjects = JSON.parse(projects.text)
    }
        return (
            <>
                {isLoading ? <Button className="nextButton" disabled>
                    <Spinner
                        as="span"
                        animation="grow"
                        size="sm"
                        role="status"
                        aria-hidden="true"
                    />
                    Loading
                </Button> :
                    <Button className="nextButton" onClick={handleShow}>
                    Overfør deler til Prosjekt
                </Button>}

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
                        <div className={"scaffoldingElement-modal"}>
                            <div className={"transfer-options"}>
                                <span>Overfør til prosjekt:</span>
                                <select
                                    className={"form-select"}
                                    value={ToProject}
                                    onChange={(e) => setToProject(e.target.value)}>
                                    <option  defaultValue="">Choose here</option>
                                    <option value={0}>Storage</option>
                                    {jsonProjects?.map(e => {
                                        return (
                                            <option key={e.projectID} value={e.projectID}>{e.projectName}</option>
                                        )
                                    })}
                                </select>
                            </div>
                            <div>
                                <span>Overfør fra prosjekt:</span>
                                <select
                                    className={"form-select"}
                                    value={FromProject}
                                    onChange={(e) => setFromProject(e.target.value)}>
                                    <option  defaultValue="">Choose here</option>
                                    <option value={0}>Storage</option>
                                    {jsonProjects?.map(e => {
                                        return (
                                            <option key={e.projectID} value={e.projectID}>{e.projectName}</option>
                                        )
                                    })}
                                </select>
                            </div>
                            {jsonProject[0].scaffolding.map(e => {
                                    return (
                                        <div key={e.type} className={"card"}>
                                            <section className={"header"}>
                                                <h3>{e.type.toUpperCase()}</h3>
                                            </section>
                                            <section className={"image"}>
                                                <img className={"img"}
                                                     src={require(`../../scaffolding/images/${e.type.charAt(0).toUpperCase() + e.type.slice(1)}.jpg`)}
                                                     alt={""}></img>
                                            </section>
                                            <input
                                                className={"form-control"}
                                                placeholder={"Enter quantity of scaffolding parts to transfer"}
                                                type="number"
                                                min={0}
                                                key={"input" + e.type}
                                                onChange={(j) => setQuantity(j, e.type)}/>
                                        </div>
                                    )
                                }
                            )}

                        </div>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant="secondary" onClick={handleClose}>
                            Close
                        </Button>


                        {buttonPressed ? <Button disabled>
                                <Spinner
                                    as="span"
                                    animation="grow"
                                    size="sm"
                                    role="status"
                                    aria-hidden="true"
                                />
                                Transferring...
                            </Button> :
                            <Button variant="primary" disabled={!validFormat} onClick={AddScaffold}>
                                Save Changes
                            </Button>}








                    </Modal.Footer>
                </Modal>
            </>
        );


}




