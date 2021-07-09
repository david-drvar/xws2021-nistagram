import React, {useEffect, useState} from 'react';
import "./../../style/PostMenu.css"
import { ReactComponent as Inbox } from './../../images/icons/inbox.svg'
import { ReactComponent as Comments } from './../../images/icons/comments.svg'
import { ReactComponent as Bookmark } from './../../images/icons/bookmark.svg'
import { ReactComponent as BookmarkFilled } from './../../images/icons/bookmark-filled.svg'

import { ReactComponent as Heart } from './../../images/icons/heart.svg'
import { ReactComponent as HeartFilled } from './../../images/icons/heart-filled.svg'
import { ReactComponent as BrokenHeart } from './../../images/icons/broken-heart.svg'
import { ReactComponent as BrokenHeartFilled } from './../../images/icons/broken-heart-filled.svg'
import {useSelector} from "react-redux";
import {Button, Modal} from "react-bootstrap";
import userService from "../../services/nistagram api/user.service";
import chatService from "../../services/nistagram api/chat.service";
import toastService from "../../services/toast.service";
import UserAutocomplete from "./UserAutocomplete";

const PostMenu = (props) => {
    const store = useSelector(state => state)
    const [users, setUsers] = useState([]);
    const [selectedUser, setSelectedUser] = useState({});
    const [chatRoom, setChatRoom] = useState("");
    const [conn, setConn] = useState({});

    const [showModal, setShowModal] = useState(false);
    const { isLiked, isDisliked, likeClicked, dislikeClicked, commentClicked, saveClicked, isSaved, postId } = props;

    useEffect(() => {
        getAllUsers();
    }, []);

    async function getAllUsers() {
        const response = await userService.getAllUsers({ jwt: store.apiKey.jwt });
        await setUsers(response.data.users);
    }

    async function startChat() {
        const response = await chatService.StartConversation({
            person1: store.apiKey.id,
            person2: selectedUser.id,
            jwt : store.apiKey.jwt
        });
        if (response.status === 200) {
            toastService.show("success", "Chat room retrieved successfully")
            await setChatRoom(response.data)
            await setConn(new WebSocket("ws://localhost:8003" + "/ws/" + response.data.Id));
        }
        else
            toastService.show("error", "Something went wrong. Try again")
    }

    async function sendMessage() {
        await setConn(new WebSocket("ws://localhost:8003" + "/ws/" + selectedUser.id + "_" + store.apiKey.id))
        console.log(conn)
        conn.send(JSON.stringify({SenderId : store.apiKey.id, ReceiverId : selectedUser.id, RoomId : chatRoom.Id, Content : postId, ContentType : "Post"}));
    }


    return (
        <div>
        <div className="postMenu">
            <div className="interactions">
                { isLiked ? 
                    <HeartFilled onClick={likeClicked} fill="red" className="icon" /> : 
                    <Heart onClick={likeClicked} className="icon" /> 
                }
                { isDisliked ? 
                    <BrokenHeartFilled onClick={dislikeClicked} fill="red" className="icon" /> : 
                    <BrokenHeart onClick={dislikeClicked} className="icon" /> 
                }
                <Comments onClick={commentClicked} className="icon" />
                <div onClick={(e) => setShowModal(true)}><Inbox className="icon" /></div>
            </div>
            { isSaved ? 
                <BookmarkFilled onClick={saveClicked} fill="black" className="icon" /> : 
                <Bookmark onClick={saveClicked} className="icon" />
            }



        </div>
            <Modal show={showModal} onHide={setShowModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Send this post</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <UserAutocomplete setSelectedUser={setSelectedUser} suggestions={users} />
                    <div style={{display:'flex',float:'right'}}>
                        <Button variant="info" style={{marginRight:'10px'}} onClick={(e) => sendMessage()} >Send</Button>
                        <Button variant="secondary" onClick={(e) => setShowModal(false)} >Cancel</Button>
                    </div>
                </Modal.Body>
            </Modal>
        </div>
    )
}

export default PostMenu;