import React, {useEffect, useState} from 'react';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";
import Notification from "./Notification";
import userService from "../../services/user.service";


function Notifications(props) {
    const store = useSelector(state => state);
    const[notifications,setNotifications]=useState([])

    useEffect(() => {
        setNotifications(props.location.state.notifications)
        readNotifications()
    }, []);

    async function getUserNotifications() {
        const response = await userService.getUserNotifications({
            id: store.user.id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            setNotifications(response.data.notifications)
        } else {
            console.log("NIJE nasao notifikacije")
        }
    }

    async function readNotifications() {
        const response = await userService.readNotifications({
            id: store.user.id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("procitao")
        } else {
            console.log("NIJE procitao")
        }
    }
    return (
        <div style={{marginTop:'6%'}}>
            <Navigation/>
            <div style={{marginLeft:'20%', marginRight:'20%'}}>
                 <h3 style={{borderBottom:'1px solid black'}}>Notifications</h3>
                <div style={{marginTop:'4%'}}>
                    {notifications.reverse().map((notification, i) =>
                       <Notification getUserNotifications={getUserNotifications} id={notification.id} creatorId={notification.creatorId} userId={notification.userId} text={notification.text} type={notification.type} createdAt={notification.createdAt} contentId={notification.contentId}/>
                    ) }
                </div>
            </div>
        </div>
    );
}

export default Notifications;