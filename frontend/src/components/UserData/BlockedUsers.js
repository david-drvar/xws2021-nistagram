import React, {useEffect, useState} from 'react';
import ProfileInfo from "./ProfileInfo";
import {useDispatch, useSelector} from "react-redux";
import followersService from "../../services/followers.service";
import userService from "../../services/user.service";
import ProfileForSug from "../HomePage/ProfileForSug";
import {Button} from "react-bootstrap";
import privacyService from "../../services/privacy.service";

function BlockedUsers() {
    const [blocked,setBlocked]=useState({})
    const [users,setUsers]=useState([]);
    const store = useSelector(state => state);

    useEffect(() => {
        getBlockedUsers();
    }, [users]);

    async function getBlockedUsers(){
        const response = await userService.getBlockedUsers({
            id: store.user.id,
            jwt: store.user.jwt,
        })
        if(response.status === 200) {
            setBlocked(response.data.id)
            getUsers(response.data.id)
        }else{
            console.log("Nije uspeo")
        }
    }

    async function getUserById(id) {
        const response = await userService.getUserById({
            id: id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            if(users.some(item=>item.id=response.data.id)){
                return;
            }
            setUsers(users=>[...users,response.data])
        } else {
            console.log("getuserbyid error")
        }
    }
    function getUsers(value)
    {
        value.map((id, i) =>
            getUserById(id)
        );
    }
    async function unblock(id) {
        const response = await privacyService.unBlockUser({
            UserId: store.user.id,
            BlockedUserId: id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            getBlockedUsers()
            setUsers([])
        } else {
            console.log("Nije uspeo")
        }
    }



    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                {users.map((user, i) =>
                    <div style={{display: "flex"}}>
                        <ProfileForSug user={user} username={user.username} image={user.profilePhoto} caption={user.biography} urlText="Follow" iconSize="big" captionSize="small" storyBorder={true}
                                       firstName={user.firstName} lastName={user.lastName} />
                        <Button style={{height:'45px', marginLeft:'20em', fontSize:'12px'}} variant="outline-danger"   onClick={() => unblock(user.id)}>Unblock</Button>{' '}
                    </div>
                ) }
            </div>
        </div>
    );
}

export default BlockedUsers;