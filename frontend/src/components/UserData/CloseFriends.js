import React, {useEffect, useState} from 'react';
import ProfileInfo from "./ProfileInfo";
import userService from "../../services/user.service";
import {userActions} from "../../store/actions/user.actions";
import followersService from "../../services/followers.service";
import {useSelector} from "react-redux";
import {Button} from "react-bootstrap";
import ProfileForSug from "../HomePage/ProfileForSug";

function CloseFriends() {
    const store = useSelector(state => state);
    const [closeFriends,setCloseFriends]=useState({})
    const [users,setUsers]=useState([]);


    useEffect(() => {
        getCloseFriends()
    },[users]);


    async function getCloseFriends(){
        const response = await followersService.getCloseFriends({
            id: store.user.id,
            jwt: store.user.jwt,
        })
        if(response.status === 200) {
            setCloseFriends(response.data.users)
            getUsers(response.data.users)
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
            getUserById(id.UserId)
        );
    }

    async function removeFromCloseFriends(id) {

        const response = await followersService.updateUserConnection({
            userId: store.user.id,
            followerId: id,
            isApprovedRequest: true,
            isCloseFriends: false,
            isMuted:false,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            getCloseFriends()
            setUsers([])
        } else {
            console.log("NIJE ")
        }
    }



    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                {users.map((user, i) =>
                    <div style={{display: "flex"}}>
                        <ProfileForSug user={user} username={user.username} caption={user.biography} urlText="Follow" iconSize="big" captionSize="small" storyBorder={true}
                                               firstName={user.firstName} lastName={user.lastName} image={user.profilePhoto}/>
                        <Button style={{height:'45px', marginLeft:'20em', fontSize:'12px'}} variant="outline-danger"   onClick={() => removeFromCloseFriends(user.id)}>Remove from close friends</Button>{' '}
                    </div>
                ) }
            </div>
        </div>
    );
}

export default CloseFriends;