import {useEffect, useState} from "react";
import {Alert, Button, Container, Form, Table} from "react-bootstrap";
import axios from "axios";
import PasswordStrengthBar from "react-password-strength-bar";

export function PasswordChange(props) {
        const [user, setUser] = useState();
        const [email, setEmail] = useState(props.email);
        const[password,setPassword]=useState('');
        const[rePassword,setRePassword]=useState('');
        const[passwordStrength,setPasswordStrength]=useState('');
        const[errPassword,setErrPassword]=useState('Enter password');
        const[errRePassword,setErrRePassword]=useState('Repeat password');
        const[blackListedPasswords,setBlackListedPasswords]=useState([]);
        const[success,setSuccess]=useState(false);
         const [submitted, setSubmitted] = useState(false);

         console.log(user)


    useEffect(() => {
        axios
            .get('http://localhost:8080/api/users/api/users/getByEmail/' + email)
            .then(res => {
                setUser(res.data)
            }).catch(res => {
        })
    },[]);

    function handleInputChange(event){
        setPassword(event.target.value);
        validationErrorMessage(event);
    }

    function handlePasswordChange(event){
        setRePassword(event.target.value);
        validationErrorMessage(event);
    }

    function validationErrorMessage(event){
        const {name, value} = event.target;
        switch (name) {
            case 'password':
                if(checkPassword(password)){
                    setErrPassword('Password must contains at least 8 characters (lowercase letter, capital letter, number and special character) or not be a common password!' );
                }else{
                    setErrPassword('');
                }
                break;
            case 'rePassword':
                if(isValidRepeatedPassword(rePassword)){
                    setErrRePassword('' );
                }else{
                    setErrRePassword('This password must match the previous!');
                }
                break;
            default:
                break;
        }
    }

    function isValidRepeatedPassword(value){
        if (password !== rePassword) {
            return false;
        } else {
            return true
        }
    }


    function checkPassword(password){
        console.log("Checking")
        if(/^(?=.*[\d])(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*])[\w!@#$%^&*]{8,}$/.test(password)){
            setPasswordStrength(password);
            return false;
        }else if(blackListedPasswords.includes(password)){
            setPasswordStrength(password);
            return false;
        } else {
            setPasswordStrength('');
            return true;
        }
    }

    function validateForm(errors){
        let valid = true;
        for (const Error of errors) {
            validationErrorMessage(createTarget(Error));
        }
        if (errPassword !== "" || errRePassword !== "") {
            return !valid;
        }
        return valid;
    }

    function createTarget(error) {
        return {target : {value : error, name : error}}
    }

    async function submitPassword(event){
        setSubmitted(true);
        event.preventDefault();
        const errors = ['password','rePassword'];

        if (validateForm(errors)) {
            await sendParams()
        } else {
            console.log('Invalid Form')
        }
    }

    async function sendParams(){
        console.log("FLLLLLLLLLLLLLLLLL")
        console.log(user)
        axios
            .post('http://localhost:8080/api/users/api/users/changeForgottenPass', {
                password: {
                    Id: user.id,
                    OldPassword: 'bla',
                    NewPassword: password,
                    RepeatedPassword: rePassword

                }
            })
            .then(res => {
               // setOldErr('');
                setSuccess(true);

            }).catch(res => {
            console.log("NE RADI")
            //setOldErr('Please enter valid old password!');


        })
    }

    return (
        <div>
            <tr>
                <td colSpan="2">
                    {!success ?
                        <p style={{textAlign: 'center', margin: 10}}> Change your password.<br/>Password must
                            contains at least 8 characters (lowercase letter, capital letter, number and special
                            character) or not be a common password! </p>
                        :
                        <p style={{textAlign: 'center', margin: 20}}> Successfully </p>
                    }
                </td>
                <td>
                    <a href={'/'} style={{'color': '#08B8A2',float : "right"}}> Back to login page?</a>

                </td>
            </tr>
            {success ?
                <tr>
                    <td>
                        <p style={{textAlign: 'center', margin: 20}}>Password Updated!<br/>
                            Your password has been changed successfully. <br/>
                            Use your new password to log in.</p>
                    </td>
                    <td>
                        <Button style={{width: "350px", display: 'block', margin: 'auto'}} variant="outline-warning"
                                href={'/'}>LOG IN</Button>

                    </td>
                </tr>
                :
                <div>
                    <tr>
                        <td style={{overflowY: "auto", width: "500px", textAlign: 'center'}}> Enter your password:
                        </td>
                        <td>
                            <Form.Control style={{width: "400px"}}  autoFocus type="password" name="password" onChange={handleInputChange} value={password}/>
                            {submitted &&
                            <span className="text-danger">{errPassword}</span>}
                            <PasswordStrengthBar password={passwordStrength}/>

                        </td>

                    </tr>

                    <tr>
                        <td style={{overflowY: "auto", width: "500px", textAlign: 'center'}}> Re-type password:</td>
                        <td>
                            <Form.Control style={{width: "400px"}} type="password" name="rePassword" placeholder="Repeat new Password"
                                          value={rePassword} onChange={handlePasswordChange}/>
                            {submitted &&
                            <span className="text-danger">{errRePassword}</span>}

                        </td>

                    </tr>
                    <tr>
                        <td colSpan="2">
                            <Button variant="info" style={{display: 'block', margin: 'auto', backgroundColor:"#08B8A2"}}
                                    onClick={submitPassword}> Submit </Button>
                        </td>

                    </tr>
                </div>

            }

        </div>
    )
}