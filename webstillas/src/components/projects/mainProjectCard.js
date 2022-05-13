import React from 'react'
import 'bootstrap/dist/css/bootstrap.min.css';
import '../../Assets/Styles/card.css'
import img from '../../Assets/Images/scaffoldingimg.jpg'
import {Link} from "react-router-dom";
import {PROJECTS_URL} from "../../Constants/apiURL";
import deleteModel from "../../Middleware/deleteProject";
import {useQueryClient} from "react-query";
import {IconButton} from "@material-ui/core";
import DeleteIcon from "@material-ui/icons/Delete";


/**
 * Function that will display an information card, with information of a project.
 *
 * @param props data sent from another view.
 * @returns {JSX.Element}
 */
function CardElement(props) {
    const queryClient = useQueryClient()

    /**
     * Function that will delete an existing project.
     *
     * @returns {Promise<void>}
     */
    const DeleteProject = async () => {
        if (window.confirm("Er du sikker på at du vil slette? " + props.name + "?")) {
            const deleteBody =
                [
                    {
                        id: props.id
                    }
                ]
            try {
                await deleteModel(PROJECTS_URL, (deleteBody)).then(() =>
                    window.alert("Prosjektet er slettet")
                )
                await queryClient.invalidateQueries("allProjects")
            } catch (e) {
                window.alert("Noe gikk galt! Prøv igjen senere")
        }
        }
    }


    return (
        <div className={"card"}>
            <div className={"name-btn"}>
                <section className={"header"}>
                    <h3>{props.name}</h3>
                </section>
                <IconButton onClick={DeleteProject}>
                    <DeleteIcon style={{fontSize: 50}}/>
                </IconButton>
            </div>
            <img className={"image-project"} src={img} alt={""}/>
            <div className={"information-highlights"}>
                <ul className={"information-list"}>
                    <li className={"horizontal-list"}>
                        <div className={"highlightText"}>
                            <span>{props.state}</span>
                        </div>
                        <div className={"highlightText-caption"}>
                            <span>Status</span>
                        </div>
                    </li>
                    <li className={"horizontal-list"}>
                        <div className={"highlightText"}>
                            <span>{props.rentPeriod}</span>
                        </div>
                        <div className={"highlightText-caption"}>
                            <span>Leieperiode</span>
                        </div>
                    </li>
                    <li className={"horizontal-list"}>
                        <div className={"highlightText"}>
                            <span>&nbsp;&nbsp; {props.size}</span>
                            <span>&#13217;</span>
                        </div>
                        <div className={"highlightText-caption"}>
                            <span>Størrelse</span>
                        </div>
                    </li>
                </ul>
            </div>
            <div className={"information-highlights"}>
                <ul className={"contact-list"}>
                    <li className={"horizontal-list-contact"}>
                        <span className={"left-contact-text"}>Kontakt person</span>
                        <span className={"right-contact-text"}>{props.contactPerson}</span>
                    </li>
                    <li className={"horizontal-list-contact"}>
                        <span className={"left-contact-text"}>Adresse</span>
                        <span
                            className={"right-contact-text"}>{props.address_Street}, {props.address_zip} {props.address_Municipality}</span>
                    </li>
                    <li className={"horizontal-list-contact"}>
                        <span className={"left-contact-text"}>Nummer</span>
                        <span className={"right-contact-text"}>{props.contactNumber}</span>
                    </li>
                </ul>
            </div>
            <div className={"card-btns"}>
                <Link className={"btn"} to={"/project/" + props.id}>Mer Informasjon</Link>
            </div>
        </div>
    )
}

export default CardElement
