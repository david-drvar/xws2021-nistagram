import React, {useEffect, useState} from 'react';
import { useDispatch, useSelector } from 'react-redux';
import PasswordStrengthBar from 'react-password-strength-bar';
import {Alert, Button, FormControl} from "react-bootstrap";
import userService from "../services/user.service";

const RegistrationPage = () => {
    // Declare a new state variable, which we'll call "count"
    const [password, setPassword] = useState("");
    const [passwordStrength, setPasswordStrength] = useState("");
    const [birthDate, setBirthDate] = useState("");
    const [birthDateErr, setBirthDateErr] = useState("Enter birthdate");
    const [sex, setSex] = useState("");
    const [sexErr, setSexErr] = useState("Select sex");
    const [biography, setBiography] = useState("");
    const [username, setUsername] = useState("");
    const [usernameErr, setUsernameErr] = useState("Enter username");
    const [phoneNumber, setPhoneNumber] = useState("");
    const [phoneNumberErr, setPhoneNumberErr] = useState("Enter phone number in pattern +38160123456");
    const [email, setEmail] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [rePassword, setRePassword] = useState("");
    const [emailErr, setEmailErr] = useState("Enter email");
    const [passwordErr, setPasswordErr] = useState("Enter password");
    const [firstNameErr, setFirstNameErr] = useState("Enter first name");
    const [lastNameErr, setLastNameErr] = useState ("Enter last name");
    const [rePasswordErr, setRePasswordErr] = useState("Repeat password");
    const [submitted, setSubmitted] = useState(false);
    const [successfullyReg, setSuccessfullyReg] = useState(false);
    const [disabled, setDisabled] = useState(false);
    const [errorMessage, setErrorMessage] = useState(false);
    const [blacklistedPasswords, setBlacklistedPasswords] = useState([]);
    const [website, setWebsite] = useState("");
    const [profilePhoto, setProfilePhoto] = useState("");


    useEffect(() => {
        setBirthDateErr( birthDate !== "" ? '' : 'Enter birthdate')
        setSexErr( sex !== "" ? '' : 'Select sex')
        setUsernameErr( isUsernameValid(username) ? '' : 'Enter username')
        setPhoneNumberErr( isPhoneNumberValid(phoneNumber) ? '' : 'Enter phone number')
        setRePasswordErr( isValidRepeatedPassword(rePassword) ? '' : 'This password must match the previous!')
        setPasswordErr(checkPassword(password) ? 'Password must contains at least 8 characters (lowercase letter, capital letter, number and special character) or not be a common password!' : '')
        setEmailErr(isValidEmail(email) && email.length > 1 ? '' : 'Email is not valid!')
        setLastNameErr(checkNameAndSurname(lastName) ? '' : 'EnterLastName')
        setFirstNameErr(checkNameAndSurname(firstName) ? '' : 'EnterFirstName')
    }, [birthDate,sex,username,phoneNumber,rePassword,password,email,lastName,firstName])

    const handleInputChange = (event) => {
        const target = event.target;
        switch (target.name) {
            case "firstName" :
                setFirstName(target.value);
                break;
            case "lastName" :
                setLastName(target.value);
                break;
            case "email" :
                setEmail(target.value);
                break;
            case "password" :
                setPassword(target.value);
                break;
            case "rePassword" :
                setRePassword(target.value);
                break;
            case "birthDate" :
                setBirthDate(target.value);
                break;
            case "sex" :
                setSex(target.value);
                break;
            case "phoneNumber" :
                setPhoneNumber(target.value);
                break;
            case "website" :
                setWebsite(target.value);
                break;
            case "username" :
                setUsername(target.value);
                break;
            case "biography" :
                setBiography(target.value);
                break;
        }
        validationErrorMessage(event);
    }

    function validationErrorMessage(event) {
        const { name } = event.target;

        switch (name) {
            case 'firstName':
                setFirstNameErr(checkNameAndSurname(firstName) ? '' : 'EnterFirstName')
                break;
            case 'lastName':
                setLastNameErr(checkNameAndSurname(lastName) ? '' : 'EnterLastName')
                break;
            case 'email':
                setEmailErr(isValidEmail(email) && email.length > 1 ? '' : 'Email is not valid!')
                break;
            case 'password':
                setPasswordErr(checkPassword(password) ? 'Password must contains at least 8 characters (lowercase letter, capital letter, number and special character) or not be a common password!' : '')
                break;
            case 'rePassword':
                setRePasswordErr( isValidRepeatedPassword(rePassword) ? '' : 'This password must match the previous!')
                break;
            case 'phoneNumber':
                setPhoneNumberErr( isPhoneNumberValid(phoneNumber) ? '' : 'Enter phone number')
                break;
            case 'username':
                setUsernameErr( isUsernameValid(username) ? '' : 'Enter username')
                break;
            case 'sex':
                setSexErr( sex !== "" ? '' : 'Select sex')
                break;
            case 'birthDate':
                setBirthDateErr( birthDate !== "" ? '' : 'Enter birthdate')
                break;
            default:
                /*this.setState({
                    validForm: true
                })*/
                break;
        }

    }

    function isUsernameValid(value) {
        return /^[a-z0-9_.]+$/.test(value);
    }

    function isPhoneNumberValid(value) {
        return /^[+]?[0-9]{8,12}$/.test(value);
    }

    function checkNameAndSurname(value) {
        return /^[a-zA-Z ,.'-]+$/.test(value);

    }

    function checkPassword (password) {
        console.log("Checking")
        if(/^(?=.*[\d])(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*])[\w!@#$%^&*]{8,}$/.test(password)){
            setPasswordStrength(password);
            return false;
        }else if(blacklistedPasswords.includes(password)){
            setPasswordStrength(password);
            return false;
        } else {
            setPasswordStrength("");
            return true;
        }
    }

    function isValidEmail (value) {
        return !(value && !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,64}$/i.test(value));
    }
    function isValidRepeatedPassword (value)  {
        return password === rePassword;
    }

    async function submitForm (event) {
        setSubmitted(true);

        event.preventDefault();
        const errors = ['email', 'password', 'firstName', 'rePassword', 'lastName', 'username', 'birthDate', 'sex', 'phoneNumber'];
        if (validateForm(errors)) {
            await sendParams()
        } else {
            console.log('Invalid Form')
        }
    }

    function validateForm(errors) {
        let valid = true;
        for(const Error of errors) {
            validationErrorMessage(createTarget(Error));
        }
        //todo promeniti!
        if(emailErr !== "" || passwordErr !== "" || firstNameErr !== "" ||
            lastNameErr !== "" || rePasswordErr !== "" || usernameErr !== "" || phoneNumberErr !== "" ||
            sexErr !== "" || birthDateErr !== "")
            return !valid;
        return valid;
    }

    function createTarget (error) {
        return {target : {value : error, name : error}}
    }

    async function sendParams() {
        //setBirthDate(new Date(birthDate));
        const jsonDate = birthDate + 'T' + '01:30:15.01Z';

        const response = await userService.createUser({
            id:'1',
            firstName: firstName,
            lastName: lastName,
            email: email,
            username: username,
            password: password,
            role: 'Basic',
            birthdate: jsonDate,
            profilePhoto: profilePhoto,
            phoneNumber: phoneNumber,
            sex: 'MAN',
            isActive: true,
            biography: biography,
            website: website,
        })
        if (response.status === 200) {
                console.log(response.data)
                setErrorMessage(false);
                setSuccessfullyReg(true);
                setDisabled(!disabled);
        } else {
            setErrorMessage(true);
            console.log("NE RADI")
        }
    }

    function handleChangeImage(evt) {
        console.log("Uploading");
        var self = this;
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function(upload) {
            setProfilePhoto(upload.target.result)
        };
        reader.readAsDataURL(file);
    }

    return (
        <div  className="App">
            {/*<h2 id="createCertifiacate"> Create certificate </h2>*/}
            <div className="row">
                <label className="col-sm-2 col-form-label">*Name</label>
                <div className="col-sm-5 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""} type="text" value={firstName} name="firstName" onChange={(e) =>
                        handleInputChange(e) } className="form-control" placeholder="First Name"/>
                    {submitted && firstNameErr.length > 0 && <span className="text-danger">{firstNameErr}</span>}

                </div>
                <div className="col-sm-5 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}  type="text" value={lastName} name="lastName" onChange={(e) => handleInputChange(e) } className="form-control" placeholder="Last Name"/>
                    {submitted && lastNameErr.length > 0 && <span className="text-danger">{lastNameErr}</span>}
                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">*Birth date</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""} min="1900-01-02" max="2009-01-01"  type="date" value={birthDate} name="birthDate" onChange={(e) => handleInputChange(e) } className="form-control" id="birthDate" />
                    {submitted && birthDateErr.length > 0 && <span className="text-danger">{birthDateErr}</span>}
                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">*Sex</label>
                <div className="col-sm-6 mb-2">
                    <select onChange={(e) => handleInputChange(e)} name={"sex"} value={sex}>
                        <option disabled={true} value="">Select sex</option>
                        <option value="MALE">Male</option>
                        <option value="FEMALE">Female</option>
                        <option value="OTHER">Other</option>
                    </select>
                    {submitted && sexErr.length > 0 && <span className="text-danger">{sexErr}</span>}
                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">*Phone number</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="text" value={phoneNumber} name="phoneNumber" onChange={(e) => handleInputChange(e) } className="form-control" id="phoneNumber" placeholder="+38160123456" />
                    {submitted && phoneNumberErr.length > 0 && <span className="text-danger">{phoneNumberErr}</span>}

                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">Website</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="text" value={website} name="website" onChange={(e) => handleInputChange(e) } className="form-control" id="website" placeholder="www.example.com" />
                    {/*{submitted && website.length > 0 && <span className="text-danger">{phoneNumberErr}</span>}*/}

                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">*Username</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="text" value={username} name="username" onChange={(e) => handleInputChange(e) } className="form-control" id="username" />
                    {submitted && usernameErr.length > 0 && <span className="text-danger">{usernameErr}</span>}

                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">Biography</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="text" value={biography} name="biography" onChange={(e) => handleInputChange(e) } className="form-control" id="biography" />
                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">*Email</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="email" value={email} name="email" onChange={(e) => handleInputChange(e) } className="form-control" id="email" placeholder="example@gmail.com" />
                    {submitted && emailErr.length > 0 && <span className="text-danger">{emailErr}</span>}

                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label className="col-sm-2 col-form-label">*Password</label>
                <div className="col-sm-6 mb-2">
                    <FormControl disabled = {(disabled)? "disabled" : ""}  name="password" type="password" placeholder="Password"  value={password} onChange={(e) => handleInputChange(e) }/>
                    {submitted && passwordErr.length > 0 &&  <span className="text-danger">{passwordErr}</span>}
                    <PasswordStrengthBar password={passwordStrength} />
                </div>
                <div className="col-sm-4">
                </div>
            </div>

            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">*Repeat password</label>
                <div className="col-sm-6 mb-2">
                    <FormControl  disabled = {(disabled)? "disabled" : ""}  name="rePassword" type="password" placeholder="Repeat new Password" value={rePassword} onChange={(e) => handleInputChange(e) }/>
                    {submitted && rePasswordErr.length > 0 &&  <span className="text-danger">{rePasswordErr}</span>}

                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">*Profile photo</label>
                <div className="col-sm-6 mb-2">
                    {/*<input type="file" onChange={(e) => setProfilePhoto(e.target.files[0])} />*/}
                    <input type="file" name="file"
                           className="upload-file"
                           id="file"
                           onChange={handleChangeImage}
                           formEncType="multipart/form-data"
                           required />
                </div>
                <div className="col-sm-4">
                </div>
            </div>

            {
                successfullyReg ?
                    <Alert variant='success' show={true}  style={({textAlignVertical: "center", textAlign: "center"})}>
                        Successfully registered please login.
                    </Alert>
                    :
                    <div className="row" style={{marginTop: '1rem'}}>
                        <div className="col-sm-5 mb-2">
                        </div>
                        <div className="col-sm-4">
                            <Button variant="success" onClick={submitForm}>Confirm</Button>
                        </div>
                    </div>
            }

            {
                errorMessage &&
                <Alert variant='danger' show={true}  style={({textAlignVertical: "center", textAlign: "center"})}>
                    The e-mail address must be unique! Please try again
                </Alert>
            }
        </div>
    );
}

export default RegistrationPage;