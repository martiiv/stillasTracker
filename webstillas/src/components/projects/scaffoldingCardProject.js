import React from "react";


/**
 * Function that will display number of expected and registered scaffolding parts in the project.
 *
 * @param props wil return key, type, expected and registered, of that scaffolding
 * @returns {JSX.Element}
 */
function ScaffoldingProject(props){
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
                                    <span>{props.expected}</span>
                                </div>
                                <div className={"highlightText-caption"}>
                                    <span>Expected</span>
                                </div>
                            </li>
                            <li className={"horizontal-list"}>
                                    <div className={"highlightText"}>
                                        <span >{props.registered}</span>
                                    </div>
                                    <div className={"highlightText-caption"}>
                                        <span>Registered</span>
                                    </div>
                                </li>
                        </ul>
                    </div>
                </section>
            </article>
        </div>
    )
}

export default ScaffoldingProject
