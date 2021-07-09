import Navigation from "./Navigation";
import Feed from './Feed';
import React, { useEffect, useState } from "react";
import { Button, Modal } from "react-bootstrap";
import PasswordStrengthBar from "react-password-strength-bar";
import { useDispatch, useSelector } from 'react-redux';
import userService from "../../services/user.service";
import "../../style/home.css";
import Sidebar from "./Sidebar";

const Home = () => {
    const [showModal, setModal] = useState(false);
    const [submitted, setSubmitted] = useState(false);
    const [passwordStrength, setPasswordStrength] = useState('');
    const [passwords, setPasswords] = useState({ oldPassword: '', newPassword: '', repeatedPassword: '' });
    const [oldErr, setOldErr] = useState('');
    const [newErr, setNewErr] = useState('');
    const [repErr, setRepErr] = useState('');

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        // checkProfile();
    },[]);


    async function checkProfile() {
        const response = await userService.checkIsApproved({
            id: store.user.id,
            jwt: store.user.jwt
        })

        if (response.status === 200) {
            checkIsApproved(response.data.response)
        } else {
            console.log("NEJEJ")
        }
    }

    function  checkIsApproved(value){
        if(store.user.isSSO) return
        if(value === false){
            setModal(true)
        }else{
            setModal(false)
        }
    }

    async function approveAccount() {
        const response = await userService.approveAccount({
            id: store.user.id,
            oldPassword: passwords.oldPassword,
            newPassword: passwords.newPassword,
            repeatedPassword: passwords.repeatedPassword,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setOldErr('');
            setModal(false);
        } else {
            setOldErr('Please enter valid old password!');
        }
    }

    function  handleInputChange(event) {
        setPasswords({
            ...passwords,
            [event.target.name]: event.target.value,
        });
        validationErrorMessage(event)
    }

    function submitForm (event){
        setSubmitted(true)
        validatePasswords();

        if(newErr === '' &&  repErr === ''){
            approveAccount();
        }
    }

    function validatePasswords(){
        if(passwords.newPassword == "" && oldErr == ""  ){
            setNewErr('Please enter new password.')
        }else if(checkPassword(passwords.newPassword)){
            setNewErr('Password must contains at least 8 characters (lowercase letter, capital letter, number and special character) or not be a common password!')
        }else{
            setNewErr('')
        }

        if(passwords.newPassword !== passwords.repeatedPassword && oldErr === ""){
            setRepErr('This password must match the previous.')
        }else{
            setRepErr('')
        }
    }

    function validationErrorMessage(event) {
        const {name, value} = event.target;
        console.log(name)
        switch (name) {
            case 'newPass':
                setNewErr(checkPassword(passwords.newPass) ? 'Password must contains at least 8 characters (lowercase letter, capital letter, number and special character) or not be a common password!' : '')
                break;
            case 'repPass':
                setRepErr( isValidRepeatedPassword(passwords.repPass) ? '' : 'This password must match the previous!')
                break;
            default:
                break;
        }

    }

    function checkPassword (password){
        console.log("Checking")
        if(/^(?=.*[\d])(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*])[\w!@#$%^&*]{8,}$/.test(password)){
            setPasswordStrength(passwords.newPassword)
            return false;
        }else {
            setPasswordStrength("")
            return true;
        }
    }

    function isValidRepeatedPassword (){
        if(passwords.newPassword !== passwords.repeatedPassword) {
            return false;
        }else{
            return  true
        }
    }

    function handleModal() {
        setModal(!showModal)
    }

    function showModalDialog(){
        return (
            <Modal backdrop="static" show={showModal} onHide={handleModal}>
                <Modal.Header  style={{'background':'#E0E0E0'}}>
                    <Modal.Title>Verify your account:</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{'background':'#C0C0C0'}}>
                    <p> You have to change password when you log in for first time.</p> <br/>
                    <p> First password : </p> <input name="oldPassword" onChange={e=>handleInputChange(e)} value={passwords.oldPassword} type={"password"}/>
                    {submitted &&  <label className="text-danger">{oldErr}</label>}
                    <p> New password : </p> <input name="newPassword" onChange={e=>handleInputChange(e)} value={passwords.newPassword} type={"password"}/>
                    <PasswordStrengthBar password={passwordStrength}/>
                    {submitted &&  <label className="text-danger">{newErr}</label>}
                    <p> Repeat new password : </p> <input name="repeatedPassword" onChange={(e) => {handleInputChange(e)}} value={passwords.repeatedPassword} type={"password"}/>
                    {submitted &&  <label className="text-danger">{repErr}</label>}
                </Modal.Body>
                <Modal.Footer style={{'background':'#E0E0E0'}}>
                    <Button variant="secondary" onClick={submitForm}>
                        Send
                    </Button>
                </Modal.Footer>
            </Modal>
        )
    }

    return (
        <div className="App">
            <Navigation/>
            <main>
                <div className="container">
                <Feed />
                { /*showModalDialog()*/ }
                <Sidebar/>
                </div>
            </main>
        </div>
    );
}

export default Home;