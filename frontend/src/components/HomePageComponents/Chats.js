import React, {useEffect, useState} from 'react';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";
import userService from "../../services/user.service";
import UserAutocomplete from "../Post/UserAutocomplete";
import {Button, Dropdown, DropdownButton, InputGroup} from "react-bootstrap";

import "./../../style/chat.css"
import verificationRequestService from "../../services/verificationRequest.service";
import toastService from "../../services/toast.service";
import chatService from "../../services/chat.service";


function Chats() {
    const [users, setUsers] = useState([]);
    const dispatch = useDispatch();
    const store = useSelector(state => state);
    const [selectedUser, setSelectedUser] = useState({});
    const [messageText, setMessageText] = useState("");
    const [chatRoom, setChatRoom] = useState("");
    const [conn, setConn] = useState({});
    const [messages, setMessages] = useState([]);
    const [messageCategory, setMessageCategory] = useState("Message category");


    useEffect(() => {
        getAllUsers();
    }, []);

    useEffect(() => {
        //conn change
        if (conn !== {}) {
            // conn.onmessage()
            // gledaj da budu javan i da su follow

            conn.onmessage = function(event) {
                const message = JSON.parse(event.data)
                setMessages([...messages,message])
            };

        }
    }, [conn, messages]);

    async function getAllUsers() {
        const response = await userService.getAllUsers({ jwt: store.user.jwt });
        await setUsers(response.data.users);
    }

    async function getMessagesForChatRoom() {
        console.log("chatRoom");
        console.log(chatRoom);
        const response = await chatService.GetMessagesForChatRoom({
            roomId : chatRoom.Id,
            jwt : store.user.jwt
        });
        if (response.status === 200) {
            console.log("messages");
            console.log(response.data);
            toastService.show("success", "Chat room messages retrieved successfully")
            await setMessages(response.data)
        }
        else
            toastService.show("error", "Something went wrong. Try again")
    }

    async function startChat() {
        const response = await chatService.StartConversation({
            person1: store.user.id,
            person2: selectedUser.id,
            jwt : store.user.jwt
        });
        if (response.status === 200) {
            toastService.show("success", "Chat room retrieved successfully")
            await setChatRoom(response.data)
            await setConn(new WebSocket("ws://localhost:8003" + "/ws/" + response.data.Id));
            getMessagesForChatRoom()


        }
        else
            toastService.show("error", "Something went wrong. Try again")
    }

    function sendMessage() {
        setMessageCategory("Message category");
        setMessageText("");
        const temp = {senderId : store.user.id, receiverId : selectedUser.id, roomId : chatRoom.Id, content : messageText, contentType : "String"};
        console.log(temp);
        conn.send(JSON.stringify({SenderId : store.user.id, ReceiverId : selectedUser.id, RoomId : chatRoom.Id, Content : messageText, ContentType : messageCategory}));
    }


    function handleChangeImage(evt) {
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function (upload) {
            setMessageText(upload.target.result)
        };
        reader.readAsDataURL(file);
    }

    function handleMessageCategoryChange(event) {
        setMessageText("");
        setMessageCategory(event);
    }

    return (
        <div style={{marginTop:'5%', marginLeft : "5%"}}>
            <Navigation/>
            <h1>Chat</h1>
            <UserAutocomplete setSelectedUser={setSelectedUser} suggestions={users} />
            <Button style={{marginLeft : "1%"}} onClick={startChat}>Start chat</Button>

            <br/><br/><br/>

            <div style={{overflow: "scroll", height:"400px"}}>
                {messages.map((message) => {
                    return (
                        <div>
                            {message.SenderId === store.user.id && <div className="container">
                                <img src="" alt="Avatar"/>
                                {message.ContentType === "String" && <p>{message.Content}</p>}
                                {message.ContentType === "Image" && <img src={message.Content} alt="Avatar"/>}
                                {message.ContentType === "Link" && <p style={{color : "powderblue"}} onClick={(e) => alert("namesti link click")}>{message.Content}</p>}
                                <span style={{marginLeft: "20px"}} className="time-right">{message.DateCreated}</span>
                            </div>
                            }

                            {message.ReceiverId === store.user.id && <div className="container darker">
                                <img src="" alt="Avatar"/>
                                {message.ContentType === "String" && <p>{message.Content}</p>}
                                {message.ContentType === "Image" && <img src={message.Content} alt="Avatar"/>}
                                {message.ContentType === "Link" && <p style={{color : "powderblue"}} onClick={(e) => alert("namesti link click")}>{message.Content}</p>}
                                <span style={{marginLeft: "20px"}} className="time-right">{message.DateCreated}</span>
                            </div>}
                        </div>
                    )
                })}

                {/*<div className="container">*/}
                {/*    <img src="" alt="Avatar"/>*/}
                {/*        <p>Hello. How are you today?</p>*/}
                {/*        <span className="time-right">11:00</span>*/}
                {/*</div>*/}

                {/*<div className="container darker">*/}
                {/*    <img src="" alt="Avatar" className="right"/>*/}
                {/*        <p>Hey! I'm fine. Thanks for asking!</p>*/}
                {/*        <span className="time-left">11:01</span>*/}
                {/*</div>*/}
            </div>

            <DropdownButton onSelect={(event) => handleMessageCategoryChange(event) } as={InputGroup.Append}  variant="outline-secondary" title={messageCategory} id="input-group-dropdown-2" >
                <Dropdown.Item eventKey={"String"} >String</Dropdown.Item>
                <Dropdown.Item eventKey={"Image"} >Image</Dropdown.Item>
                <Dropdown.Item eventKey={"Link"} >Link</Dropdown.Item>
            </DropdownButton>

            {(messageCategory === "String" || messageCategory === "Link") &&
                <input type={"text"} value={messageText} onChange={(e) => setMessageText(e.target.value)}/>
            }
            {messageCategory === "Image" &&
            <input type="file" name="file"
                   className="upload-file"
                   id="file"
                   onChange={handleChangeImage}
                   formEncType="multipart/form-data"
                   required />            }
            <Button style={{marginLeft : "1%"}} onClick={sendMessage}>Send message</Button>

        </div>
    );
}


export default Chats;