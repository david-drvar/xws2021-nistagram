import {Button, Dropdown, Modal} from "react-bootstrap";
import {BsThreeDotsVertical} from "react-icons/all";
import React, {useEffect, useState} from "react";
import followersService from "../../services/followers.service";
import privacyService from "../../services/privacy.service";
import {useSelector} from "react-redux";
import {useHistory} from "react-router-dom";
import Switch from "react-switch";

function BlockMuteAndNotifications(props){
    const { isApprovedRequest, isMuted, notifications } = props;
    const [muted, setIsMuted]=useState({})
    const [showBlockModal, setBlockModal] = useState(false);
    const [showMuteModal, setMuteModal] = useState(false);
    const [showNotificationsModal, setNotificationsModal] = useState(false);
    const [isMessageNotificationEnabled, setMessagesNotifications] = useState(notifications.isMessageNotificationEnabled);
    const [isPostNotificationEnabled, setPostNotifications] = useState(notifications.isPostNotificationEnabled);
    const [isStoryNotificationEnabled, setStoryNotifications] = useState(notifications.isStoryNotificationEnabled);
    const [isCommentNotificationEnabled, setCommentsNotifications] = useState(notifications.isCommentNotificationEnabled);
    const [update, setUpdate] = useState(false);

    const store = useSelector(state => state);
    const history = useHistory()

    useEffect(() => {
        setIsMuted(isMuted)
    },[isMuted]);

    useEffect(() => {
        setMessagesNotifications(isMessageNotificationEnabled)
        setPostNotifications(isPostNotificationEnabled)
        setStoryNotifications(isStoryNotificationEnabled)
        setCommentsNotifications(isCommentNotificationEnabled)
    },[isMessageNotificationEnabled,isPostNotificationEnabled,isStoryNotificationEnabled,isCommentNotificationEnabled]);

    function handleBlockModal(){
        setBlockModal(!showBlockModal)
    }
    function handleMuteModal(){
        setMuteModal(!showMuteModal)
    }
    function handleNotificationsModal(){
        setNotificationsModal(!showNotificationsModal)
    }
    function handleMessagesNotifications(){
        if(!update) setUpdate(true)
        setMessagesNotifications(!isMessageNotificationEnabled)
    }
    function handlePostNotifications(){
        if(!update) setUpdate(true)
        setPostNotifications(!isPostNotificationEnabled)
    }
    function handleCommentNotifications(){
        if(!update) setUpdate(true)
        setCommentsNotifications(!isCommentNotificationEnabled)
    }
    function handleStoryNotifications(){
        if(!update) setUpdate(true)
        setStoryNotifications(!isStoryNotificationEnabled)
    }

    async function blockUser() {
        const response = await privacyService.blockUser({
            UserId: store.user.id,
            BlockedUserId: store.followers.followerId,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
           console.log("USPEO")
            history.push({ pathname: '/' })
        } else {
            console.log("Nije uspeo")
        }
    }

    async function muteUser() {
        const response = await followersService.updateUserConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            isCloseFriends: false,
            isMuted:true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("JESTE mjut")
            setIsMuted(true)
            handleMuteModal()
        } else {
            console.log("NIJE mjut")
        }
    }
    async function unMuteUser() {
        const response = await followersService.updateUserConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            isCloseFriends: false,
            isMuted:false,
            isNotificationEnabled:true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("JESTE")
            setIsMuted(false)
            handleMuteModal()
        } else {
            console.log("NIJE ")
        }
    }

    async function menageNotifications() {
        const response = await followersService.updateUserConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            isCloseFriends: false,
            isMuted:false,
            isMessageNotificationEnabled: isMessageNotificationEnabled,
            isPostNotificationEnabled:isPostNotificationEnabled,
            isStoryNotificationEnabled:isStoryNotificationEnabled,
            isCommentNotificationEnabled:isCommentNotificationEnabled,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            handleNotificationsModal()
        } else {
            console.log("NIJE ")
        }
    }

    return(
        <div>
            <Dropdown>
                <Dropdown.Toggle variant="link" id="dropdown-basic">
                    <BsThreeDotsVertical/>
                </Dropdown.Toggle>
                <Dropdown.Menu>
                    <Dropdown.Item onClick={handleBlockModal}>Block</Dropdown.Item>
                    {isApprovedRequest && <Dropdown.Item onClick={handleMuteModal}>Mute/Unmute</Dropdown.Item>}
                    {isApprovedRequest && <Dropdown.Item onClick={handleNotificationsModal}>Menage notifications</Dropdown.Item>}
                </Dropdown.Menu>
            </Dropdown>

            <Modal show={showBlockModal} onHide={handleBlockModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Block user</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <p style={{fontWeight:'bold'}}>Are you sure you want to <p style={{color:'red',display: 'inline'}}>block </p>this user?</p>
                    <p style={{fontSize:'13px'}}>After blocking you will be redirected to home page.</p>
                    <div style={{display:'flex',float:'right'}}>
                        <Button variant="danger" style={{marginRight:'10px'}} onClick={blockUser}>Yes</Button>
                        <Button variant="info" onClick={handleBlockModal}>No</Button>
                    </div>
                </Modal.Body>
            </Modal>

            <Modal show={showMuteModal} onHide={handleMuteModal}>
                <Modal.Header closeButton>
                    <Modal.Title>{muted ? "Unmute" :"Mute"}</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    {!muted ?
                    <div>
                        <p style={{fontWeight:'bold',display: 'inline'}}>Are you sure you want to <p style={{color:'red',display: 'inline'}}>mute</p> this user?</p>
                        <div style={{display:'flex',float:'right'}}>
                            <Button variant="danger" style={{marginRight:'10px'}} onClick={muteUser}>Yes</Button>
                            <Button variant="info" onClick={handleMuteModal}>No</Button>
                        </div>
                    </div>
                    :
                    <div>
                        <div>
                            <p style={{fontWeight:'bold',display: 'inline'}}>Are you sure you want to <p style={{color:'red',display: 'inline'}}>unmute</p> this user?</p>
                            <div style={{display:'flex',float:'right'}}>
                                <Button variant="danger" style={{marginRight:'10px'}} onClick={unMuteUser}>Yes</Button>
                                <Button variant="info" onClick={handleMuteModal}>No</Button>
                            </div>
                        </div>
                    </div>}
                    <br/>
                </Modal.Body>
            </Modal>

            <Modal show={showNotificationsModal} onHide={handleNotificationsModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Menage notifications</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <tr>
                        <td> <p style={{marginRight:'38px', fontWeight:'bold'}}>Message notifications:</p>  </td>
                        <td > <Switch  onChange={handleMessagesNotifications} checked={isMessageNotificationEnabled}/></td>
                        <td> {isMessageNotificationEnabled ? <p style={{marginLeft:'12px', color:'green'}} >enabled</p> :<p style={{marginLeft:'12px', color:'grey'}} >not enabled</p>}</td>
                    </tr>
                    <tr>
                        <td> <p style={{marginRight:'38px', fontWeight:'bold'}}>Post notifications:</p>  </td>
                        <td > <Switch  onChange={handlePostNotifications} checked={isPostNotificationEnabled}/></td>
                        <td> {isPostNotificationEnabled ? <p style={{marginLeft:'12px', color:'green'}} >enabled</p> :<p style={{marginLeft:'12px', color:'grey'}} >not enabled</p>}</td>
                    </tr>
                    <tr>
                        <td> <p style={{marginRight:'38px', fontWeight:'bold'}}>Story notifications:</p>  </td>
                        <td > <Switch  onChange={handleStoryNotifications} checked={isStoryNotificationEnabled}/></td>
                        <td> {isStoryNotificationEnabled ? <p style={{marginLeft:'12px', color:'green'}} >enabled</p> :<p style={{marginLeft:'12px', color:'grey'}} >not enabled</p>}</td>
                    </tr>
                    <tr>
                        <td> <p style={{marginRight:'38px', fontWeight:'bold'}}>Comment notifications:</p>  </td>
                        <td > <Switch  onChange={handleCommentNotifications} checked={isCommentNotificationEnabled}/></td>
                        <td> {isCommentNotificationEnabled ? <p style={{marginLeft:'12px', color:'green'}} >enabled</p> :<p style={{marginLeft:'12px', color:'grey'}} >not enabled</p>}</td>
                    </tr>
                    {update &&
                    <div style={{display: 'flex', float: 'right'}}>
                        <Button variant="success" style={{marginRight: '10px'}} onClick={menageNotifications}>Update settings</Button>
                    </div>
                    }
                </Modal.Body>
            </Modal>
        </div>



    );

}export default  BlockMuteAndNotifications;