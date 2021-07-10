import React, {useState} from "react";
import {Alert, Col, Container, FormControl, Row} from "react-bootstrap";
import axios from "axios";
import PasswordStrengthBar from "react-password-strength-bar";
import userService from "../../services/user.service";
import {useDispatch, useSelector} from "react-redux";
import ProfileInfo from "./ProfileInfo";
import toastService from "../../services/toast.service";

function ChangePassword() {
    const [user, setUser] = useState();
    const[passwords,setPasswords]=useState({oldPass:'',newPass:'',repPass:''})
    const[submitted,setSubmitted]=useState(false)
    const[oldErr,setOldErr]=useState('');
    const[newErr,setNewErr]=useState('');
    const[repErr,setRepErr]=useState('');
    const [passwordStrength, setPasswordStrength] = useState("");
    const dispatch = useDispatch()
    const store = useSelector(state => state);

    async function changePass() {
        const response = await userService.changePassword({
            id: store.user.id,
            oldPassword: passwords.oldPass,
            newPassword: passwords.newPass,
            repeatedPassword: passwords.repPass,
            jwt: store.user.jwt,

        })

        if (response.status === 200) {
            setOldErr('');
            toastService.show("success", "Successfully updated!");
        } else {
            toastService.show("error", "Could not update image!");

            setOldErr('Please enter valid old password!');
        }
    }



    function handleChange(event) {
        setPasswords({
            ...passwords,
            [event.target.name]: event.target.value,
        });
        validationErrorMessage(event);
    }
    function validationErrorMessage(event) {
        const { name, value } = event.target;
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
    function isValidRepeatedPassword (value)  {
        var pera= passwords.newPass === passwords.repPass;
        return pera;
    }

    function  activateUpdateMode(event){
        setSubmitted(true);
        validatePasswords();

        if(newErr=='' && repErr==''){
            changePass();
        }
    }
    function checkPassword (password) {
        console.log("Checking")
        if(/^(?=.*[\d])(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*])[\w!@#$%^&*]{8,}$/.test(password)){
            setPasswordStrength(password);
            return false;
            //else if(blacklistedPasswords.includes(password)){
            //   setPasswordStrength(password);
            //   return false;
        } else {
            setPasswordStrength("");
            return true;
        }
    }

    function validatePasswords(){
        if(passwords.newPass=="" && oldErr==""  ){
            setNewErr('Please enter new password.')
        }else if(checkPassword(passwords.newPass)){
            setNewErr('Password must contains at least 8 characters (lowercase letter, capital letter, number and special character) or not be a common password!')

        }else{
            setNewErr('')
        }

        if(passwords.newPass!=passwords.repPass && oldErr==""){
            setRepErr('This password must match the previous.')
        }else{
            setRepErr('')

        }
    }
    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
            <div>
                <Row className="m-2">
                    <FormControl name="oldPass" type="password" placeholder="Please enter old password"  value={passwords.oldPass} onChange={handleChange}/>
                    {submitted &&  <label className="text-danger">{oldErr}</label>}
                </Row>
                <Row className="m-2">
                    <FormControl name="newPass" type="password" placeholder="Enter new Password" value={passwords.newPass} onChange={handleChange}/>
                    <PasswordStrengthBar password={passwordStrength} />

                    {submitted &&  <label className="text-danger">{newErr}</label>}
                </Row>
                <Row className="m-2">
                    <FormControl name="repPass" type="password" placeholder="Repeat new Password" value={passwords.repPass} onChange={handleChange}/>
                    {submitted &&  <label className="text-danger">{repErr}</label>}
                </Row>
                <button  style={{marginRight:'100px', float:'right'}} type="button" className="btn btn-outline-danger" onClick={activateUpdateMode}>Edit password</button>
            </div>
        </div>
        </div>
    );
}
export default ChangePassword;