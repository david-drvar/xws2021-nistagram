import Switch from "react-switch";
import React, {useEffect, useState} from "react";
import {Alert, Button} from "react-bootstrap";
import privacyService from "../../services/privacy.service";
import {useDispatch, useSelector} from "react-redux";
import ProfileInfo from "./ProfileInfo";
import toastService from "../../services/toast.service";

function  EditUserPrivacy(){
    const [checkedPrivacy,setCheckedPrivacy]=useState({});
    const [checkedDm,setCheckedDm]=useState({});
    const [checkedTag,setCheckedTag]=useState({});

    const [submitted,setSubmitted]=useState(false);
    const store = useSelector(state => state);

    useEffect(() => {
        getUserPrivacyInfo()
    }, []);

    async function getUserPrivacyInfo() {
        const response = await privacyService.getUserPrivacyInfo({
            id: store.user.id,
        })

        if (response.status === 200) {
            setCheckedPrivacy(!response.data.isProfilePublic)
            setCheckedDm(!response.data.isDmPublic)
            setCheckedTag(!response.data.isTagEnabled)
        } else {
            console.log("nebravo")
        }
    }

    function handlePrivacyChange() {
        setCheckedPrivacy(!checkedPrivacy)
    }
    function handleDmChange() {
        setCheckedDm(!checkedDm)
    }
    function handleTagChange() {
        setCheckedTag(!checkedTag)
    }
    async function editPrivacy() {
        const response = await privacyService.updateUserPrivacy({
            Id: store.user.id,
            isProfilePublic: !checkedPrivacy,
            isDmPublic: !checkedDm,
            isTagEnabled: !checkedTag,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            toastService.show("success", "Successfully updated!");

            setSubmitted(true)
            console.log("bravo")
        } else {
            console.log("nebravo")
        }
    }

    return(

        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                <div style={{marginTop:'15'}}>
                <tr>
                    <td>
                        <p style={{marginRight:'38px', fontWeight:'bold'}}>Profile privacy:</p>
                    </td>
                    <td >
                        <Switch  onChange={handlePrivacyChange} checked={checkedPrivacy}/>
                    </td>
                    <td>
                        {checkedPrivacy ? <p style={{marginLeft:'12px', color:'red'}} >private</p> :<p style={{marginLeft:'12px', color:'red'}} >public</p>}
                    </td>
                </tr>
                    <tr>
                        <td>
                            <p style={{marginRight:'28px', fontWeight:'bold'}}>Dm privacy:</p>
                        </td>
                        <td >
                            <Switch  onChange={handleDmChange} checked={checkedDm}/>
                        </td>
                        <td>
                            {checkedDm ? <p style={{marginLeft:'12px', color:'red'}} >private</p> :<p style={{marginLeft:'12px', color:'red'}} >public</p>}
                        </td>
                    </tr>
                    <tr>
                        <td>
                            <p style={{marginRight:'28px', fontWeight:'bold'}}>Tag privacy:</p>
                        </td>
                        <td >
                            <Switch  onChange={handleTagChange} checked={checkedTag}/>
                        </td>
                        <td>
                            {checkedTag ? <p style={{marginLeft:'12px', color:'red'}} >private</p> :<p style={{marginLeft:'12px', color:'red'}} >public</p>}
                        </td>
                    </tr>
                    <Button style={{float: "right",marginRight:'130px', marginTop:'15px'}} variant="secondary" onClick={editPrivacy}>Save</Button>
                </div>
        </div>
        </div>
    );
}export default EditUserPrivacy;