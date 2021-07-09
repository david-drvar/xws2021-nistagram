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
import { ReactComponent as Ad } from "../../images/icons/ad.svg";

import ProfileIcon from "../ProfileComponent/ProfileIcon";
import { NavLink, useHistory } from "react-router-dom";
import React, { useEffect, useState } from "react";
import { useSelector, useDispatch } from "react-redux";
import { Dropdown, Button } from "react-bootstrap";
import { userActions } from "../../store/actions/user.actions";

import userService from "../../services/agent.service";

function Menu() {
    const [username, setUsername]=useState('')
    const [notifications, setNotifications] = useState([])
    const [newIcon, setNewIcon] = useState(false)
    
    const store = useSelector(state => state);
    const history = useHistory(); 
    const dispatch = useDispatch();

    useEffect(() => {
        // store.user.jwt && getUsernameById();
        // getUserNotifications()
    }, []);

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

            {store.user.role === 'Agent' && store.user.jwt !== "" && (<NavLink to={{pathname: "/newproduct"}}> <Plus className="icon" />  </NavLink>) }
            <NavLink to={{pathname: "/info"}}> <Explore className="icon"/> </NavLink>

            {
                store.user.jwt !== "" && store.user.role === "Agent" &&
                (<NavLink to={"/campaigns"}>
                    <Ad className="icon" />
                </NavLink>)
            }
            <span style={{marginLeft: "20px"}}/>

            {store.user.role === 'Agent' && store.user.jwt !== "" && (<NavLink to={"/profile/" + store.user.username}>
                <ProfileIcon iconSize="medium"
                    image={store.user.photo ? store.user.photo : 'https://i.pravatar.cc/150?img=1'}/>
            </NavLink>
            )}



            <span style={{marginLeft: "20px"}}/>

            

            {store.user.jwt !== "" ? 
                <Button variant="outline-danger" onClick={logout} style={{width: "220px", display: "block"}}>Logout</Button> :
                <Button variant="primary" onClick={login} style={{width: "100px", marginLeft: "1em", display: "block"}}>Login</Button>
            }

        </div>
    );
}
export default Menu;