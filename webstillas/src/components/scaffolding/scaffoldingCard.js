import React from 'react'
import '../../Assets/Styles/scaffoldingCard.css'
import InfoModal from "./projectsModal";


/**
 * Will return an infromation card, where the user get information about how many of the specific
 * scaffolding the storage has, including total amount that is in projects.
 *
 * @param props type of scaffolding
 * @returns {JSX.Element}
 */
function CardElement(props){
    console.log(props.type)
    console.log(`../images/${props.type}.jpg`)
    return(
        <div className={"scaffoldingElement"}>
            <article className={"card"}>
                <section className={"header"}>
                    <h3>{props.type.toUpperCase()}</h3>
                </section>
                <section className={"image"}>
                    <img className={"img"} src={require(`../../Assets/Images/scaffoldingImages/${props.type.charAt(0).toUpperCase() + props.type.slice(1)}.jpg`)} alt={""}></img>
                </section>
                <section className={"information-highlights-cta"}>
                    <div className={"information-highlights"}>
                        <ul className={"information-list"}>
                            <li className={"horizontal-list"}>
                                <div className={"highlightText"}>
                                    <span>{props.total}</span>
                                </div>
                                <div className={"highlightText-caption"}>
                                    <span>Totalmengde</span>
                                </div>
                            </li>
                            <li className={"horizontal-list"}>
                                <div className={"highlightText"}>
                                    <span >{props.storage}</span>
                                </div>
                                <div className={"highlightText-caption"}>
                                    <span>Lager</span>
                                </div>
                            </li>
                        </ul>
                    </div>
                </section>
                <section className={"card-btn"}>
                    <div className={"card-btns"}>
                        <InfoModal type ={props.type}
                                    total = {props.total}
                                    storage = {props.storage}/>
                    </div>
                </section>
            </article>
            <>
            </>
        </div>
    )
}

export default CardElement
