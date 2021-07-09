import ProfileInfo from "./ProfileInfo";
import React, {useEffect, useState} from "react";
import {Button} from "react-bootstrap";
import userService from "../../services/user.service";
import {useSelector} from "react-redux";
import toastService from "../../services/toast.service";
import followersService from "../../services/followers.service";
import apiService from "../../services/api.service";

function APIKey(){
    const [apiKey, setApiKey] = useState("asd");
    const store = useSelector(state => state);


    useEffect(() => {
        getAPIKey();
    }, []);

    async function getAPIKey() {
        const response = await apiService.GetKeyByUserId({
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

    return(
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                   <input type={"text"} value={apiKey}/>
            </div>
        </div>
    );

}export  default  APIKey;