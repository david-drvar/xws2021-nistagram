import {Button} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import userService from "../../services/agent.service";
import {useSelector} from "react-redux";
import toastService from "../../services/toast.service";
import moment from "moment";
import "../../style/notification.css";

function Notification(props) {
    const {id, creatorId, userId, text, type, createdAt, contentId} = props;
    const [user, setUser] = useState({});
    const [privateFollow, setPrivateFollow] = useState(false);
    const [contentType, setContentType] = useState(false);
    const store = useSelector(state => state);
    const [hoursAgo, setHoursAgo] = useState(0)
    const [daysAgo, setDaysAgo] = useState(0);
    const [minutesAgo, setMinutesAgo] = useState(0)
    const [post, setPost] = useState({})
    const [showModal, setShowModal] = useState(false);

    useEffect(() => {
        setDate()
        getUserById()
        checkType()
    }, []);

    async function getUserById(id) {
        const response = await userService.getUserById({
            id: creatorId,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUser(response.data)
        } else {
            console.log("getuserbyid error")
        }
    }

    function checkType() {
        if (type === "FollowPrivate") {
            setPrivateFollow(true)
        }
        if (type === "Like" || type === "Dislike" || type === "Comment") {
            setContentType(true)
        }
    }

    async function acceptRequest() {

    }

    function handleReject() {
        unfollow()
        deleteNotification()
    }

    async function unfollow() {


    }

    async function deleteNotification() {
        const response = await userService.deleteNotification({
            id: id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            props.getUserNotifications()
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }

    function setDate() {
        const currentTime = moment(new Date())

        if (Math.floor(moment.duration(currentTime.diff(createdAt)).asHours()) < 24) {
            if (Math.floor(moment.duration(currentTime.diff(createdAt)).asHours()) <= 0) {
                setMinutesAgo(Math.floor(moment.duration(currentTime.diff(createdAt)).asMinutes()))
            } else {
                setHoursAgo(Math.floor(moment.duration(currentTime.diff(createdAt)).asHours()))
            }
        } else {
            setDaysAgo(Math.floor(moment.duration(currentTime.diff(createdAt)).asDays()))
        }
    }

    async function getPostById() {

    }

    return (
        <div style={{display: "flex", marginLeft: '10%'}}>
            {contentType ?
                <font face="Comic Sans MS" size="3" style={{marginRight: '5em', fontWeight: 'bold', color:'black'}}>
                    <Button variant="link" style={{ color:'black' }} onClick={getPostById}>{text}</Button>
                </font>
                :
                <font face="Comic Sans MS" size="3" style={{marginRight: '5em', fontWeight: 'bold'}}>{text}</font>
            }
            <div className="timePosted">
                {daysAgo < 1 ? (
                    hoursAgo < 1 ? <p style={{fontSize: '11px'}}>{minutesAgo} minutes ago</p> :
                        <p> {hoursAgo} hours ago</p>
                ) : <p style={{fontSize: '11px'}}>{daysAgo} days ago</p>
                }
            </div>
            {privateFollow &&
            <div style={{display: "flex", marginLeft: '85px'}}>
                <Button style={{height: '27px', fontSize: '12px'}} variant="success"
                        onClick={() => acceptRequest()}>Accept</Button>
                <Button style={{marginLeft: '5px', height: '27px', fontSize: '12px'}} variant="secondary"
                        onClick={() => handleReject()}>Reject</Button>
            </div>
            }

        </div>
    );
}
export default Notification;