import "../../style/menu.css";
import { ReactComponent as Home } from "../../images/icons/home.svg";
import { ReactComponent as Inbox } from "../../images/icons/inbox.svg";
import { ReactComponent as Notifications } from "../../images/icons/notifications.svg";
import { ReactComponent as NewNotifications } from "../../images/icons/newNotification.svg";
import { ReactComponent as Bookmark } from "../../images/icons/bookmark.svg";
import { ReactComponent as StoryArchive } from "../../images/icons/story-archive.svg";
import { ReactComponent as Plus } from "../../images/icons/plus.svg";
import { ReactComponent as Explore } from "../../images/icons/more.svg";
import { ReactComponent as VerificationSymbol } from "../../images/icons/verification-symbol.svg";

import ProfileIcon from "../ProfileComponent/ProfileIcon";
import { NavLink, useHistory } from "react-router-dom";
import React, { useEffect, useState } from "react";
import { useSelector, useDispatch } from "react-redux";
import { Dropdown, Button } from "react-bootstrap";
import { userActions } from "../../store/actions/user.actions";

import userService from "../../services/user.service";

function Menu() {
    const [username, setUsername]=useState('')
    const [notifications, setNotifications] = useState([])
    const [newIcon, setNewIcon] = useState(false)
    
    const store = useSelector(state => state);
    const history = useHistory(); 
    const dispatch = useDispatch();

    useEffect(() => {
        store.user.jwt && getUsernameById();
        getUserNotifications()
    }, []);

    async function getUserNotifications() {
        const response = await userService.getUserNotifications({
            id: store.user.id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            setNotifications(response.data.notifications)
            checkNotificationIcon(response.data.notifications)
        } else {
            console.log("NIJE nasao notifikacije")
        }
    }

    function checkNotificationIcon(value){
        if(value.some(item => item.isRead === false)){
            setNewIcon(true)
            console.log("TU SI")
        }
    }


    async function getUsernameById() {
        const response = await userService.getUsernameById({
            id: store.user.id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUsername(response.data.username)
        } else {
            console.log("getuserbyusername error")
        }
    }

    function verificationRedirect(text) {
        switch (text) {
            case 'submit-verification-request' :
                history.push({pathname: '/submit-verification-request'});
                break;
            case 'view-my-verification-request' :
                history.push({pathname: '/view-my-verification-request'})
                break;
            case 'view-pending-verification-request' :
                history.push({pathname: '/view-pending-verification-request'})
                break;
            case 'view-all-verification-request' :
                history.push({pathname: '/view-all-verification-request'})
                break;
            default:
                return;
        }
    }

    const logout = () => {
        dispatch(userActions.logoutRequest())
        history.push({ pathname: '/login' })
    }

    const login = () => {
        history.push({ pathname: '/login' })
    }

    return (
        <div className="menu">
            <NavLink to={{pathname: "/"}}>
                <Home className="icon"/>
            </NavLink>
            {store.user.role !== 'Admin' && store.user.jwt !== "" && (<NavLink to={{pathname: "/chats"}}> <Inbox className="icon"/> </NavLink>) }
            {store.user.role !== 'Admin' && store.user.jwt !== "" && 
                (<NavLink to={{pathname: "/notifications", state: {notifications: notifications}}}>
                    {newIcon ? <NewNotifications className="icon"/> : <Notifications className="icon"/>}
                </NavLink>) 
            }
            {store.user.role !== 'Admin' && store.user.jwt !== "" && (<NavLink to={{pathname: "/saved"}}>  <Bookmark className="icon"/> </NavLink>) }
            {store.user.role !== 'Admin' && store.user.jwt !== "" && (<NavLink to={{pathname: "/story-archive"}}> <StoryArchive className="icon" />  </NavLink>) }
            {store.user.role !== 'Admin' && store.user.jwt !== "" && (<NavLink to={{pathname: "/newpost"}}> <Plus className="icon" />  </NavLink>) }
            {store.user.role !== 'Admin' && store.user.jwt !== "" && (<NavLink to={{pathname: "/info"}}> <Explore className="icon"/> </NavLink>) }
            {store.user.jwt !== "" && (
                <Dropdown>
                    <Dropdown.Toggle variant="link" id="dropdown-basic">
                        <VerificationSymbol className="icon"/>
                    </Dropdown.Toggle>

                    <Dropdown.Menu>
                        {store.user.role !== 'Admin' && <Dropdown.Item onClick={() => verificationRedirect('submit-verification-request')}>Submit verification request</Dropdown.Item>}
                        {store.user.role !== 'Admin' && <Dropdown.Item onClick={() => verificationRedirect('view-my-verification-request')}>View my verification requests</Dropdown.Item>}
                        {store.user.role === 'Admin' && <Dropdown.Item onClick={() => verificationRedirect('view-pending-verification-request')}>View pending verification requests</Dropdown.Item>}
                        {store.user.role === 'Admin' && <Dropdown.Item onClick={() => verificationRedirect('view-all-verification-request')}>View all verification requests</Dropdown.Item>}
                    </Dropdown.Menu>
                </Dropdown>
            )}
            {store.user.role !== 'Admin' && store.user.jwt !== "" && (<NavLink to={"/profile/" + username}> 
                <ProfileIcon iconSize="medium"
                    image={store.user.photo ? store.user.photo : 'https://i.pravatar.cc/150?img=1'}/>
            </NavLink>
            )}

            {store.user.jwt !== "" ? 
                <Button variant="outline-danger" onClick={logout} style={{width: "220px", display: "block"}}>Logout</Button> :
                <Button variant="primary" onClick={login} style={{width: "100px", marginLeft: "1em", display: "block"}}>Login</Button>
            }

        </div>
    );
}
export default Menu;