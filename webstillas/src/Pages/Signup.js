import React, {useState} from "react";
import {Link, useNavigate} from "react-router-dom";
import {Form, Alert} from "react-bootstrap";
import {Button} from "react-bootstrap";
import {useUserAuth} from "../Config/UserAuthContext";
import postModel from "../Middleware/postModel";
import {formatDateToString} from "./projects";
import {LOGIN} from "../Constants/webURL";
import {USER_POST_URL} from "../Constants/apiURL";
import "../Assets/Styles/firebaselogin.css"

/**
 * Function that will register a new user to the system.
 *
 * @returns {JSX.Element}
 * @constructor
 */
const Signup = () => {
    const [email, setEmail] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [role, setRole] = useState("");
    const [phone, setPhone] = useState(0);
    const [admin, setAdmin] = useState(false);
    const [birthDay, setBirthDay] = useState("");
    const [error, setError] = useState("");
    const [password, setPassword] = useState("");
    const {signUp} = useUserAuth();
    let navigate = useNavigate();

    /**
     * Will sign up the user, then add the user to the database.
     * On sucsess the user is navigated back to log in site
     *
     * @returns {Promise<void>}
     */
    const handleSubmit = async () => {
        try {
            signUp(email, password).then(newUser => {
                const user =
                    {
                        "employeeID": newUser.user.uid,
                        "name": {
                            "firstName": firstName,
                            "lastName": lastName
                        },
                        "role": role,
                        "phone": phone,
                        "email": email,
                        "admin": admin,
                        "dateOfBirth": birthDay
                    }
                    JSON.stringify(user)
                postModel(USER_POST_URL, user)
                    .then(() => navigate(LOGIN))

            })
        } catch (err) {
            console.log(err)
            setError(err.message);
        }
    };


    return (
        <div className={"card signup-box"}>
                <h2 className="mb-3">Stillas bruker registrering </h2>
                {error && <Alert variant="danger">{error}</Alert>}
                <Form onSubmit={handleSubmit}>
                    <Form.Group className="mb-3" controlId="firstName">
                        <Form.Control
                            type="text"
                            placeholder="Fornavn"
                            onChange={(e) => setFirstName(e.target.value)}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="lastName">
                        <Form.Control
                            type="text"
                            placeholder="Etternavn"
                            onChange={(e) => setLastName(e.target.value)}
                        />
                    </Form.Group>

                    <Form.Group className="mb-3" controlId="lastName">
                        <Form.Control
                            type="number"
                            placeholder="Telefonnummer"
                            onChange={(e) => setPhone(Number(e.target.value))}
                        />
                    </Form.Group>

                    <div className={"selectors"}>
                        <Form.Group className="admin-select" controlId="admin">
                            <Form.Select onChange={(e) => setAdmin(Boolean(e.target.value))}>
                                <option value={"false"}>False</option>
                                <option value={"true"}>True</option>
                            </Form.Select>
                        </Form.Group>

                        <Form.Group className="role-select" controlId="role">
                            <Form.Select onChange={(e) => setRole(e.target.value)}>
                                <option value={"admin"}>Administrator</option>
                                <option value={"installer"}>Installatør</option>
                                <option value={"storage"}>Lagerarbeider</option>

                            </Form.Select>
                        </Form.Group>
                    </div>


                    <div className={"date-picker-signup"}>
                        <label className={"date"} htmlFor="startDate">Fødselsdag</label>
                        <input  type="date" onChange={(event) => setBirthDay(formatDateToString(event.target.value))}/>
                    </div>

                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Control
                            type="email"
                            placeholder="Email"
                            onChange={(e) => setEmail(e.target.value)}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicPassword">
                        <Form.Control
                            type="password"
                            placeholder="Passord"
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </Form.Group>
                    <div className="signup-btn">
                        <Button variant="primary" type="Submit">
                            Registrer
                        </Button>

                    </div>
                </Form>
            <div className="signup-text ">
                Har du allerede en bruker? <Link to={LOGIN}>Logg inn</Link>
            </div>
        </div>
    );
};


export default Signup;
