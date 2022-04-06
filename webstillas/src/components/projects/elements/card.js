import React from 'react'
import 'bootstrap/dist/css/bootstrap.min.css';
import './card.css'
import img from '../images/blog-item.jpg'


function CardElement(props){
    return(
        <div className={"main"}>
            <article className={"card"}>
                    <section className={"header"}>
                        <h3>Project Name</h3>
                    </section>
                    <section className={"image"}>
                        <img src={img} alt={""}/>
                    </section>
                    <section className={"information-highlights-cta"}>
                        <div className={"information-highlights"}>
                            <ul className={"information-list"}>
                                <li className={"horizontal-list"}>
                                    <div className={"highlightText"}>
                                        <span>NTNU Prosjekt</span>
                                    </div>
                                    <div className={"highlightText-caption"}>
                                        <span>Leier</span>
                                    </div>
                                </li>
                                <li className={"horizontal-list"}>
                                    <div className={"highlightText"}>
                                        <span >01.02.2020-02.02.2022</span>
                                    </div>
                                    <div className={"highlightText-caption"}>
                                        <span>Leieperiode</span>
                                    </div>
                                </li>
                                <li className={"horizontal-list"}>
                                    <div className={"highlightText"}>
                                        <span >&nbsp;&nbsp; 200</span>
                                        <span >&#13217;</span>
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
                                    <span className={"right-contact-text"}>Ola Nordmann</span>
                                </li>
                                <li className={"horizontal-list-contact"}>
                                    <span className={"left-contact-text"}>Adresse</span>
                                    <span className={"right-contact-text"}>TEKNOLOGIVEGEN 22, 2815 GJØVIK</span>
                                </li>
                                <li className={"horizontal-list-contact"}>
                                    <span className={"left-contact-text"}>Nummer</span>
                                    <span className={"right-contact-text"}>61135400</span>
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
