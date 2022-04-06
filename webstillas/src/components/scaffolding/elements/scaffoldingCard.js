import React from 'react'
import img from "../images/spirstillas_solideq_spir_klasse_5_stillas_135_1.jpg";
import './scaffoldingCard.css'


function CardElement(props){
    return(
        <div className={"scaffoldingElement"}>
            <article className={"card"}>
                <section className={"header"}>
                    <h3>{props.type.toUpperCase()}</h3>
                </section>
                <section className={"image"}>
                    <img className={"img"} src={img} alt={""}/>
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
                        <button className={"btn"} type={"button"}>Mer informasjon</button>
                    </div>
                </section>
            </article>
        </div>
    )
}

export default CardElement
