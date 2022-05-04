import React from 'react'
import 'bootstrap/dist/css/bootstrap.min.css';
import './card.css'
import img from '../images/blog-item.jpg'
import {Link} from "react-router-dom";
import {PROJECTS_URL} from "../../../modelData/constantsFile";
import deleteModel from "../../../modelData/deleteProject";
import {useQueryClient} from "react-query";
import {IconButton} from "@material-ui/core";
import DeleteIcon from "@material-ui/icons/Delete";

function CardElement(props) {
    const queryClient = useQueryClient()
    const DeleteProject = async () => {
        if (window.confirm("Are you sure you want to delete " + props.name + "?" )){
            const deleteBody =
                [
                    {
                        id: props.id
                    }
                ]
            await deleteModel(PROJECTS_URL, (deleteBody)).catch(e => console.log(e)).then(e => console.log("success"))
            await queryClient.invalidateQueries("allProjects")
        }
    }

    return (
        <div className={"main"}>
            <article className={"card"}>
                <div className={"name-btn"}>
                    <section className={"header"}>
                        <h3>{props.name}</h3>
                    </section>
                    <IconButton className={"btn-delete"} onClick={DeleteProject}>
                        <DeleteIcon style={{ fontSize: 50 }} />
                    </IconButton>
                </div>

                <section className={"image-project"}>
                    <img src={img} alt={""}/>
                </section>
                <section className={"information-highlights-cta"}>
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
                </section>
                <section className={"contact-highlights-cta"}>
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
                </section>
                <section className={"card-btn"}>
                    <div className={"card-btns"}>
                        <Link className={"btn"} to={"/project/" + props.id}>Mer Informasjon</Link>
                    </div>
                </section>
            </article>
        </div>
    )
}

export default CardElement
