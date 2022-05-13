import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Form, Alert } from "react-bootstrap";
import { Button } from "react-bootstrap";
import { useUserAuth } from "../Config/UserAuthContext";
import {PROJECT_URL, SIGNUP} from "../Constants/webURL";
import "../Assets/Styles/firebaselogin.css"


/**
 * Function to display login site.
 * @returns {JSX.Element}
 * @constructor
 */
const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const { logIn } = useUserAuth();
  const navigate = useNavigate();

  /**
   * Function to submit the users request to log in. On success navigate to project site.
   * @param e forms submit
   * @returns {Promise<void>}
   */
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    try {
      await logIn(email, password);
      navigate(PROJECT_URL);
    } catch (err) {
      setError("Feil brukernavn eller passord. \nVennligst prøv igjen");
    }
  };


  return (
    <div className={"card loginpage"}>
      <div className="box">
        <h2 className="mb-3">Stillas Login</h2>
        {error && <Alert variant="danger">{error}</Alert>}
        <Form onSubmit={handleSubmit}>
          <Form.Group className="mb-3" controlId="formBasicEmail">
            <Form.Control
              type="email"
              placeholder="Email address"
              onChange={(e) => setEmail(e.target.value)}
            />
          </Form.Group>

          <Form.Group className="mb-3" controlId="formBasicPassword">
            <Form.Control
              type="password"
              placeholder="Password"
              onChange={(e) => setPassword(e.target.value)}
            />
          </Form.Group>

          <div className="loginbtn">
            <Button variant="primary" type="Submit">
              Logg inn
            </Button>
          </div>
        </Form>
        <hr />
      </div>
      <div className="logintxt">
        Har du ikke en bruker? <Link to={SIGNUP}>Registrer</Link>
      </div>
    </div>
  );
};

export default Login;
