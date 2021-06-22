import React, {useEffect, useState} from "react";
import {Button} from "react-bootstrap";
import followersService from "../../services/followers.service";
import {useDispatch, useSelector} from "react-redux";
import Switch from "react-switch";
import toastService from "../../services/toast.service";
import {user} from "../../store/reducers/user.reducer";
import userService from "../../services/user.service";

function FollowAndUnfollow(props) {
    const {isCloseFriend, followers,publicProfile} = props;
    const [follows, setFollows] = useState(false);
    const [requestIsPending, setRequestIsPending] = useState(false);
    const [closeFriend, setCloseFriend] = useState({});

    const store = useSelector(state => state);
    useEffect(() => {
        getFollowersConnection()
        setCloseFriend(props.isCloseFriends)
    }, [props.isCloseFriends])

    useEffect(() => {
        getFollowersConnection()
        props.getFollowers(store.followers.followerId)
    }, [])

    useEffect(() => {
        setFollows(followers.some(item => item.UserId === store.user.id))
    }, [followers])


    async function getFollowersConnection() {
        const response = await followersService.getFollowersConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
        })

        if (response.status === 200) {
            setFollows(response.data.isApprovedRequest)
            setCloseFriend(response.data.isCloseFriends)
            setRequestIsPending(response.data.requestIsPending)
        } else {
            console.log("followings ne radi")
        }
    }

    async function follow() {
        const response = await followersService.createConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("ZAPRACENO")
            props.getFollowers(store.followers.followerId)
            props.funcIsCloseFriend(store.followers.followerId)
            getFollowersConnection()

        } else {
            console.log("NIJE ZAPRACENO")
        }
    }

    async function unfollow() {
        const response = await followersService.unfollow({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            props.getFollowers(store.followers.followerId)
            props.funcIsCloseFriend(store.followers.followerId)
            deleteNotification()
            getFollowersConnection()
        } else {
            console.log("NIJE otpratio")
        }
    }
    async function deleteNotification() {
        const response = await userService.deleteByTypeAndCreator({
            creatorId: store.user.id,
            type: "FollowPrivate",
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log(" nije obrisao")

        } else {
            console.log("obrisao")
        }
    }
    function handleCloseFriends() {
        setCloseFriend(!closeFriend)
        setCloseFriends()
    }
    async function setCloseFriends() {
        const response = await followersService.updateUserConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            isCloseFriends: !closeFriend,
            isMuted:false,
            isNotificationEnabled:true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("JESTE")
        } else {
            console.log("NIJE ")
        }
    }

    return (
        <div style={{marginTop: "10px"}}>
            {!follows  && !requestIsPending &&
                <Button variant="primary" style={{margin: "10px"}} onClick={follow}>Follow</Button>
            }
            
            {follows  && !requestIsPending &&
                <div>
                    <div className='row'>
                        <p style={{color: '#64f427'}}>Close friend: </p>
                        <Switch onChange={handleCloseFriends} checked={closeFriend}/>
                    </div>
                    <Button onClick={unfollow}>UnFollow</Button>
                </div>
            }

            {requestIsPending &&
                <div>
                 <p style={{marginLeft: '15px', marginRight: '3em', color: '#64f427'}}>Request is pending</p>
                <Button variant="outline-primary" style={{margin: "10px", marginRight: '78px'}} onClick={unfollow}>Remove follow request</Button>
                </div>
            }
        </div>
    );
}
export default FollowAndUnfollow;