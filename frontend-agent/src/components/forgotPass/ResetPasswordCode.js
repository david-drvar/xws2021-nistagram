import {useEffect, useState} from "react";
import {Alert, Button, Container, Form, Table} from "react-bootstrap";
import axios from "axios";

export function ResetPasswordCode(props) {

    const [resetCode, setResetCode] = useState('');
    const [email, setEmail] = useState(props.email);
    const [wrongEmail, setWrongEmail] = useState(false);
    const [errorResetCode, setErrorResetCode] = useState('');
    const [submitted, setSubmitted] = useState(false);
    const [user, setUser] = useState(false);

    useEffect(() => {
         axios
            .get('http://localhost:8080/api/users/api/users/getByEmail/' + email)
            .then(res => {
                setUser(res.data)
            }).catch(res => {
             setWrongEmail(true);
            })
    },[]);


    function handleInputChange(event){
        setResetCode(event.target.value);
    }

    function submitResetCode(){
        setSubmitted(true);
        validateResetCode();

    }

    function validateResetCode(){
        console.log(user)
        axios
            .post('http://localhost:8080/api/users/api/users/validateResetCode',{
                resetCode:resetCode,
                email: user.email
            })
            .then(res => {
                setErrorResetCode('');
                props.onChangeValue(user);

            }).catch(res => {
            setErrorResetCode('Please enter valid reset code!');

            console.log("BLAAAA")
        })
    }
    return (
        <div>
            {!wrongEmail ?
                <div>
                    <tr>
                        <td colSpan="2">
                            <p style={{textAlign: 'center', margin: 20}}>Please enter the reset code that was sent
                                to
                                your email address within 1 day. It will look something like "MFcRhYpDo1"</p>
                        </td>

                    </tr>

                    <tr>
                        <td> Enter your reset code:</td>
                        <td>
                            <Form.Control autoFocus type="text" name="resetCode"
                                          onChange={handleInputChange} value={resetCode}/>
                            {submitted && <span className="text-danger">{errorResetCode}</span>}

                        </td>
                        <td colSpan="2">
                            <Button variant="info"
                                    style={{display: 'block', margin: 'auto', backgroundColor: "#08B8A2"}}
                                    onClick={submitResetCode}> Confirm </Button>
                        </td>
                    </tr>

                </div>
                :
                <div>
                    <tr style={{overflowY: "auto", width: "500px", textAlign: 'center'}}>
                        <td colSpan="2">
                            <p style={{textAlign: 'center', margin: 20}}>We couldn't find {email} email
                                address please </p>
                        </td>
                        <td>
                            <Button style={{width: "350px", display: 'block', margin: 'auto'}}
                                    variant="outline-warning" href={'/forgotten'}>TRY AGAIN</Button>

                        </td>
                    </tr>
                </div>

            }
        </div>
    )
}