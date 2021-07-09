import React, {useEffect, useState} from 'react';
import { useDispatch, useSelector } from 'react-redux'
import {Button, Form, Modal} from "react-bootstrap";
import ReCAPTCHA from "react-google-recaptcha";
import axios from "axios";
import { useHistory } from 'react-router-dom';
import GoogleLogin from 'react-google-login';
import { userActions } from '../store/actions/user.actions'

import agentService from '../services/agent.service'

import RegistrationPage from "./RegistrationPage";
import userService from "../services/nistagram api/user.service";
const TEST_SITE_KEY = "6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI";

const IndexPage = () => {
    const[details,setDetails]=useState({email:"", password:""});
    const[badCredentials,setCredentials]=useState(true);
    const[badEmail,setEmailErr]=useState(true);
    const[submitted,setSubmitted]=useState(false);
    const[reCaptcha,setCaptcha]=useState(0);
    const[logInDisabled,setLogInDisabled]=useState(false);
    const[showModal,setShowModal]=useState(false);
    const history = useHistory()

    const dispatch = useDispatch();
    const store = useSelector(state => state)

    async function sendParams(){
        const response = await agentService.login({
            email: details.email,
            password: details.password
        })

        //console.log(response)

        if(response.status === 200) {
            await dispatch(userActions.loginRequest({
                jwt: response.data.accessToken,
                id: response.data.userId,
                role: response.data.role,
                isSSO: response.data.isSSO,
                username: response.data.username,
                photo: response.data.photo,
            }));
            if (response.data.role === "Agent") {
                const responseToken = await agentService.GetKeyByUserId({
                    id : response.data.userId,
                    jwt: response.data.accessToken,
                });
                let responseUser = await userService.getUserById({id : store.apiKey.id, jwt : response.data.accessToken});
                await dispatch(userActions.submitApiToken({
                    token : responseToken.data.token,
                    role : responseUser.role,
                    username : responseUser.username,
                    photo : responseUser.profilePhoto
                }));
            }


            history.push({ pathname: '/' })
        }else{
            if (reCaptcha >= 2) {
                setCaptcha(reCaptcha+1);
                setLogInDisabled(true);
                setCredentials(false);
            } else {
                setCaptcha(reCaptcha+1);
                setCredentials(false);
            }
        }
    }

    function submitHandler(e){
        setSubmitted(true);
        e.preventDefault();
        if(checkEmail(details.email)){
            setEmailErr(true);
            sendParams();
        }else{
            setEmailErr(false)
        }
    }

    function handleChange(event) {
        setDetails({
            ...details,
            [event.target.name]: event.target.value,
        });

        if(event.target.name==="email" && submitted){
            if(checkEmail(event.target.value)==false) setEmailErr(false);
            else setEmailErr(true);
        }
    };

    function checkEmail(value){
        return !(value && !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,64}$/i.test(value));
    }

    function closeCaptcha(){
        setCaptcha(0);
        setLogInDisabled(false);
    }

    function handleModal(){
        setShowModal(!showModal)
    }

    function closeModal(){
        setShowModal(!showModal)
    }



    return (
        <div style={{ padding: '60px 0', margin: '0 auto', maxWidth: '320px' }}>
            <br />
            <Form onSubmit={submitHandler}>
                <Form.Group size="lg" controlId="email">
                    <Form.Label>Email</Form.Label>
                    <Form.Control autoFocus type="email" name="email" onChange={handleChange}  value={details.email} />
                    <p hidden={badEmail} style={{ color: "red" }}> Email is not valid!</p>

                </Form.Group>
                <Form.Group size="lg" controlId="password">
                    <Form.Label>Password</Form.Label>
                    <Form.Control  name="password" type="password" onChange={handleChange} value={details.password}/>
                </Form.Group>
                <p hidden={badCredentials} style={{ color: "red" }}> Invalid username or password!</p>

                {reCaptcha >= 3 &&
                <ReCAPTCHA
                    style={{ display: "inline-block" }}
                    theme="light"
                    ref={React.createRef()}
                    sitekey={TEST_SITE_KEY}
                    onChange={closeCaptcha}
                    // asyncScriptOnLoad={asyncScriptOnLoad}
                />
                }
                <Button disabled={logInDisabled} block size="lg" onClick={submitHandler}>Login </Button>
            </Form>
            <br />
            <div style={{ display: " table" }}>
                <p style={{ display: "table-cell" }}>Don't have account?</p>
                <a style={{ display: "table-cell" }} className="nav-link" style={{ 'color': '#00d8fe', 'fontWeight': 'bold' }} href='#' name="workHours" onClick={handleModal}>Register</a>
            </div>
            <div style={{ display: " table" }}>
                <a href={'/forgotten'} style={{'color': '#089A87',float : "right"}}> Forgot password?</a>
            </div>

            <Modal show={showModal} onHide={closeModal} style={{ 'height': 650 }} >
                <Modal.Header closeButton style={{ 'background': 'silver' }}>
                    <Modal.Title>Registration</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{ 'background': 'silver' }}>
                    <RegistrationPage/>
                </Modal.Body>
            </Modal>
        </div>
    );
}

export default IndexPage;