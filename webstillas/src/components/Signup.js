import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Form, Alert } from "react-bootstrap";
import { Button } from "react-bootstrap";
import { useUserAuth } from "../context/UserAuthContext";
import postModel from "../modelData/postModel";
import {formatDate, formatDateToString} from "./projects/projects";

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
  const { signUp } = useUserAuth();

  let navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    try {
      signUp(email, password).then(newUser => {
        console.log(newUser.user.uid)
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
        console.log(JSON.stringify(user))
        postModel("user", user)
            .then(() => navigate("/"))
            .catch(e => console.log(e))
      })
    } catch (err) {
      setError(err.message);
    }
  };





  console.log(birthDay)

  return (
    <>
      <div className="p-4 box">
        <h2 className="mb-3">Firebase Auth Signup</h2>
        {error && <Alert variant="danger">{error}</Alert>}
        <Form onSubmit={handleSubmit}>
          <Form.Group className="mb-3" controlId="firstName">
            <Form.Control
                type="text"
                placeholder="First name"
                onChange={(e) => setFirstName(e.target.value)}
            />
          </Form.Group>
          <Form.Group className="mb-3" controlId="lastName">
            <Form.Control
                type="text"
                placeholder="Last name"
                onChange={(e) => setLastName(e.target.value)}
            />
          </Form.Group>

          <Form.Group className="mb-3" controlId="lastName">
            <Form.Control
                type="number"
                placeholder="phone"
                onChange={(e) => setPhone(Number(e.target.value))}
            />
          </Form.Group>

          <Form.Group className="mb-3" controlId="admin">
            <Form.Select onChange={(e) => setAdmin(Boolean(e.target.value))}>
              <option value={"false"}>False</option>
              <option value={"true"}>True</option>
            </Form.Select>
          </Form.Group>

          <Form.Group className="mb-3" controlId="role">
            <Form.Select onChange={(e) => setRole(e.target.value)}>
              <option value={"admin"}>Administrator</option>
              <option value={"installer"}>Installatør</option>
              <option value={"storage"}>Lagerarbeider</option>

            </Form.Select>
          </Form.Group>

          <label htmlFor="startDate">Birthday</label>
          <input type="date" onChange={(event) => setBirthDay(formatDateToString(event.target.value))}/>

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
          <div className="d-grid gap-2">
            <Button variant="primary" type="Submit">
              Sign up
            </Button>
          </div>
        </Form>
      </div>
      <div className="p-4 box mt-3 text-center">
        Already have an account? <Link to="/">Log In</Link>
      </div>
    </>
  );
};

export default Signup;
