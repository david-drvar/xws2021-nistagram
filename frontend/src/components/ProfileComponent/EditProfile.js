import {Alert, Button, FormControl, Table} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import axios from "axios";

const EditProfile = props => {
    const [user, setUser] = useState(props.user);
    const[edit,setEdit]=useState(false)
    const[success,setSuccess]=useState(false)
    const[firstNameErr,setFirstNameErr]=useState('');
    const[lastNameErr,setLastNameErr]=useState('');
    const[emailErr,setEmailErr]=useState('');
    const[usernameErr,setUsernameErr]=useState('');
    const[birthDateErr,setBirthDateErr]=useState('');
    const[sexErr,setSexErr]=useState('');
    const[phoneNumErr,setPhoneErr]=useState('');
    const [submitted, setSubmitted] = useState(false);

       console.log("EDITTTTTT");
       console.log(user);
    console.log(user.biography);
    console.log(user.website);


    function sendParams(){
        axios
            .post('http://localhost:8080/api/users/api/users/update_profile', {
                user:{
                    id: user.id,
                    firstName: user.firstName,
                    lastName: user.lastName,
                    email: user.email,
                    phoneNumber: user.phoneNumber,
                    username: user.username,
                    profilePhoto: 'idk',
                    sex: user.sex,
                    website:user.website,
                    biography:user.biography,
                }
            })
            .then(res => {
                console.log("RADI")
                props.updateUser()
                setSuccess(true);
                sessionStorage.setItem("username",user.username);

            }).catch(res => {
            console.log("NE RADIs")
        })
    }
    async function handleInputChange(event) {
        setUser({
            ...user,
            [event.target.name]: event.target.value,
        });
        validationErrorMessage(event);
    }

    function validationErrorMessage(event) {
        const { name, value } = event.target;
        switch (name) {
            case 'firstName':
                setFirstNameErr((checkNameAndSurname(user.firstName)&& value.length>1) ? '' : 'EnterFirstName');
                break;
            case 'lastName':
                setLastNameErr((checkNameAndSurname(user.lastName) && value.length>1)? '' : 'EnterLastName');
                break;
            case 'email':
                setEmailErr(isValidEmail(user.email) && value.length > 1 ? '' : 'Email is not valid!');
                break;
            case 'phoneNumber':
                setPhoneErr((isPhoneNumberValid(user.phoneNumber) && value.length > 1)? '' : 'Enter phone number');
                break;
            case 'username':
                setUsernameErr( (isUsernameValid(user.username)&& value.length>1) ? '' : 'Enter username');
                break;
            case 'sex':
                setSexErr( user.sex !== "" ? '' : 'Select sex')
                break;
            case 'birthDate':
                setBirthDateErr( user.birthdate !== "" ? '' : 'Enter birthdate')
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

    function isValidEmail (value) {
        return !(value && !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,64}$/i.test(value));
    }

    async function submitForm (event) {
        setSubmitted(true);

        if(firstNameErr === "" && lastNameErr === "" &&  emailErr === "" &&  usernameErr === "" &&  phoneNumErr === "" && sexErr==="" && birthDateErr==="" ){
            await sendParams()
        }else {
            console.log('Invalid Form')
        }
    }



    function  activateUpdateMode(){
        setEdit(true);
    }

        return(
            <div >
                {success ?
                    <Alert variant='success' show={true} style={({textAlignVertical: "center", textAlign: "center"})}>Successfully
                        updated!</Alert>
                    :
                    <div>

                        <Table style={{marginLeft: '850px',display: 'inline-block', float: 'right', alignContent: '', maxWidth: '900px'   }}>
                            <tbody>
                            <tr>
                                <td>Username</td>
                                <td>{user.username}</td>
                                {/* {edit
                                    ?
                                    <div>
                                        <FormControl name="username" className="mt-2 mb-1" value={user.username}  onChange={handleInputChange}/>
                                        {submitted && usernameErr.length > 0 && <span className="text-danger">{usernameErr}</span>}

                                    </div>
                                    : <td>{user.username}</td>
                               }*/}
                            </tr>
                            <tr>
                                <td>First Name</td>
                                {edit
                                    ?
                                    <div>
                                        <FormControl name="firstName" className="mt-2 mb-1" value={user.firstName}   onChange={handleInputChange}/>
                                        {submitted && firstNameErr.length > 0 && <span className="text-danger">{firstNameErr}</span>}

                                    </div>
                                    : <td>{user.firstName}</td>
                                }
                            </tr>
                            <tr>
                                <td>Last Name</td>
                                {edit
                                    ?
                                    <div>
                                        <FormControl name="lastName" className="mt-2 mb-2" value={user.lastName}onChange={handleInputChange}/>
                                        {submitted && lastNameErr.length > 0 && <span className="text-danger">{lastNameErr}</span>}

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
                                        <input defaultChecked={user.birthdate} min="1900-01-02" max="2009-01-01"  type="date" value={user.birthdate} name="birthdate" onChange={(e) => handleInputChange(e) } className="form-control" id="birthdate" />
                                        {submitted && birthDateErr.length > 0 && <span className="text-danger">{birthDateErr}</span>}
                                    </div>
                                    : <td>{user.birthdate}</td>
                                }
                            </tr>
                            <tr>
                                <td>Phone Number</td>
                                {edit
                                    ?
                                    <div>
                                        <FormControl name="phoneNumber" className="mt-2 mb-2" value={user.phoneNumber} onChange={handleInputChange}/>
                                        {submitted && phoneNumErr.length > 0 && <span className="text-danger">{phoneNumErr}</span>}
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
                            <tr>
                                <td>Biography </td>
                                {edit
                                    ? <FormControl name="biography" className="mt-2 mb-2" value={user.biography}  onChange={handleInputChange}/>
                                    : <td>{user.biography}</td>
                                }
                            </tr>
                            <tr>
                                <td>Web site</td>
                                {edit
                                    ? <FormControl name="website" className="mt-2 mb-2" value={user.website} onChange={handleInputChange}/>
                                    : <td>{user.website}</td>
                                }
                            </tr>
                            {edit &&
                            <tr>
                                <button style={{margin: '10px'}} type="button" className="btn btn-primary btn-sm" onClick={submitForm}>Save
                                </button>
                                <button style={{marginLeft: '10px'}} type="button" className="btn btn-primary btn-sm" onClick={() => setEdit(false)}>Cancel
                                </button>
                            </tr>
                            }

                            </tbody>
                        </Table>
                        {!edit &&
                        <button style={{marginRight: '100px', float: 'right'}} type="button" className="btn btn-outline-danger" onClick={activateUpdateMode}>Edit profile</button>
                        }
                    </div>
                }
            </div>
        );
};export default EditProfile;