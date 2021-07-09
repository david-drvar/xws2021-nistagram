import ProfileInfo from "./ProfileInfo";
import React, {useEffect, useState} from "react";
import {Button} from "react-bootstrap";
import agentService from "../../services/agent.service";
import {useDispatch, useSelector} from "react-redux";
import toastService from "../../services/toast.service";
import {userActions} from "../../store/actions/user.actions";

function APIKey(){
    const [apiKey, setApiKey] = useState("");
    const store = useSelector(state => state);
    const dispatch = useDispatch();


    useEffect(() => {
        getAPIKey();
    }, []);

    async function getAPIKey() {
        const response = await agentService.GetKeyByUserId({
            id : store.user.id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            toastService.show("success", "Successfully retrieved!");
            setApiKey(response.data.token);
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }

    async function updateKey() {
        const response = await agentService.UpdateKey({
            id : store.user.id,
            token : apiKey,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            toastService.show("success", "Successfully updated!");
            await setApiKey(response.data.token);
            await dispatch(userActions.submitApiToken({
                token : apiKey
            }));
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }

    return(
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                   <input type={"text"} value={apiKey} onChange={(e) => setApiKey(e.target.value)}/>
                <Button variant={"outline-info"} onClick={updateKey}>Update</Button>
            </div>
        </div>
    );

}export  default  APIKey;