import {Alert, Button, FormControl, Table} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import axios from "axios";
import userService from "../../services/user.service";
import {useDispatch, useSelector} from "react-redux";
import ProfileInfo from "./ProfileInfo";
import toastService from "../../services/toast.service";

function EditProfile () {
    const [user, setUser] = useState({});
    const [edit, setEdit] = useState(false)
    const [firstNameErr, setFirstNameErr] = useState('');
    const [lastNameErr, setLastNameErr] = useState('');
    const [emailErr, setEmailErr] = useState('');
    const [usernameErr, setUsernameErr] = useState('');
    const [birthDateErr, setBirthDateErr] = useState('');
    const [sexErr, setSexErr] = useState('');
    const [phoneNumErr, setPhoneErr] = useState('');
    const [submitted, setSubmitted] = useState(false);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        getUserInfo();
    }, []);

    async function getUserInfo() {
        const response = await userService.getUserById({
            id: store.user.id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUser(response.data)
        } else {
            console.log("getuser error")
        }
    }

    async function sendParams() {
        console.log(user)
        const jsonDate = user.birthdate + 'T' + '01:30:15.01Z';

        const response = await userService.editProfile({
            id: user.id,
            firstName: user.firstName,
            lastName: user.lastName,
            email: user.email,
            phoneNumber: user.phoneNumber,
            username: user.username,
            profilePhoto: 'idk',
            birthdate: jsonDate,
            sex: user.sex,
            website: user.website,
            biography: user.biography,
            jwt: store.user.jwt,
            role : store.user.role
        })

        if (response.status === 200) {
            toastService.show("success", "Successfully updated!");
        } else {
            toastService.show("error", "Something went wrong.Please try again!");
        }

    }

    async function handleInputChange(event) {
        setUser({
            ...user,
            [event.target.name]: event.target.value,
        });
        validationErrorMessage(event);
    }

    function validationErrorMessage(event) {
        const {name, value} = event.target;
        switch (name) {
            case 'firstName':
                setFirstNameErr((checkNameAndSurname(user.firstName) && value.length > 1) ? '' : 'EnterFirstName');
                break;
            case 'lastName':
                setLastNameErr((checkNameAndSurname(user.lastName) && value.length > 1) ? '' : 'EnterLastName');
                break;
            case 'email':
                setEmailErr(isValidEmail(user.email) && value.length > 1 ? '' : 'Email is not valid!');
                break;
            case 'phoneNumber':
                setPhoneErr((isPhoneNumberValid(user.phoneNumber) && value.length > 1) ? '' : 'Enter phone number');
                break;
            case 'username':
                setUsernameErr((isUsernameValid(user.username) && value.length > 1) ? '' : 'Enter username');
                break;
            case 'sex':
                setSexErr(user.sex !== "" ? '' : 'Select sex')
                break;
            case 'birthDate':
                setBirthDateErr(user.birthdate !== "" ? '' : 'Enter birthdate')
                break;
            default:
                break;
        }
    }

    function isUsernameValid(value) {
        return /^[a-z0-9_.]+$/.test(value);
    }

    function isPhoneNumberValid(value) {
        return /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/.test(value);
    }

    function checkNameAndSurname(value) {
        return /^[a-zA-Z ,.'-]+$/.test(value);
    }

    function isValidEmail(value) {
        return !(value && !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,64}$/i.test(value));
    }

    async function submitForm(event) {
        setSubmitted(true);

        if (firstNameErr === "" && lastNameErr === "" && emailErr === "" && usernameErr === "" && phoneNumErr === "" && sexErr === "" && birthDateErr === "") {
            await sendParams()
        } else {
            console.log('Invalid Form')
        }
    }


    function activateUpdateMode() {
        setEdit(true);
    }

    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                <div>

                    <table className="table">

                    <tbody>
                        <tr>
                            <td>Username</td>
                            {edit
                                ?
                                <div>
                                    <FormControl name="username" className="mt-2 mb-1" value={user.username}
                                                 onChange={handleInputChange}/>
                                    {submitted && usernameErr.length > 0 &&
                                    <span className="text-danger">{usernameErr}</span>}

                                </div>
                                : <td>{user.username}</td>
                            }
                        </tr>
                        <tr>
                            <td>First Name</td>
                            {edit
                                ?
                                <div>
                                    <FormControl name="firstName" className="mt-2 mb-1" value={user.firstName}
                                                 onChange={handleInputChange}/>
                                    {submitted && firstNameErr.length > 0 &&
                                    <span className="text-danger">{firstNameErr}</span>}

                                </div>
                                : <td>{user.firstName}</td>
                            }
                        </tr>
                        <tr>
                            <td>Last Name</td>
                            {edit
                                ?
                                <div>
                                    <FormControl name="lastName" className="mt-2 mb-2" value={user.lastName}
                                                 onChange={handleInputChange}/>
                                    {submitted && lastNameErr.length > 0 &&
                                    <span className="text-danger">{lastNameErr}</span>}

                                </div>
                                : <td>{user.lastName}</td>
                            }
                        </tr>
                        <tr>
                            <td>Email</td>
                            <td>{user.email}</td>
                        </tr>
                        <tr>
                            <td>Birth date</td>
                            {edit
                                ?

                                <div className="col-sm-6 mb-2">
                                    <input defaultChecked={user.birthdate} min="1900-01-02" max="2009-01-01" type="date"
                                           value={user.birthdate} name="birthdate"
                                           onChange={(e) => handleInputChange(e)} className="form-control"
                                           id="birthdate"/>
                                    {submitted && birthDateErr.length > 0 &&
                                    <span className="text-danger">{birthDateErr}</span>}
                                </div>
                                : <td>{user.birthdate}</td>
                            }
                        </tr>
                        <tr>
                            <td>Phone Number</td>
                            {edit
                                ?
                                <div>
                                    <FormControl name="phoneNumber" className="mt-2 mb-2" value={user.phoneNumber}
                                                 onChange={handleInputChange}/>
                                    {submitted && phoneNumErr.length > 0 &&
                                    <span className="text-danger">{phoneNumErr}</span>}
                                </div>

                                : <td>{user.phoneNumber}</td>
                            }
                        </tr>
                        <tr>

                            <td>Sex</td>
                            {edit
                                ?
                                <div><select onChange={(e) => handleInputChange(e)} name={"sex"} value={user.sex}>
                                    <option disabled={true} value="">Select sex</option>
                                    <option value="MALE">Male</option>
                                    <option value="FEMALE">Female</option>
                                    <option value="OTHER">Other</option>
                                </select>
                                    {submitted && sexErr.length > 0 && <span className="text-danger">{sexErr}</span>}
                                </div>
                                : <td>{user.sex}</td>
                            }
                        </tr>
                        <tr style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>
                            <td>Biography</td>
                            {edit
                                ? <FormControl name="biography" className="mt-2 mb-2" value={user.biography}
                                               onChange={handleInputChange}/>
                                : <td>{user.biography}</td>
                            }
                        </tr>
                        <tr style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>
                            <td>Web site</td>
                            {edit
                                ? <FormControl name="website" className="mt-2 mb-2" value={user.website}
                                               onChange={handleInputChange}/>
                                : <td>{user.website}</td>
                            }
                        </tr>
                        {edit &&
                        <tr>
                            <button style={{margin: '10px'}} type="button" className="btn btn-primary btn-sm"
                                    onClick={submitForm}>Save
                            </button>
                            <button style={{marginLeft: '10px'}} type="button" className="btn btn-primary btn-sm"
                                    onClick={() => setEdit(false)}>Cancel
                            </button>
                        </tr>
                        }

                        </tbody>
                    </table>
                    {!edit &&
                    <button style={{marginRight: '100px', float: 'right'}} type="button"
                            className="btn btn-outline-danger" onClick={activateUpdateMode}>Edit profile</button>
                    }
                </div>
            </div>
        </div>
    );
};export default EditProfile;