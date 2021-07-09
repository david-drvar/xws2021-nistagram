import React, {useEffect, useState} from 'react';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";
import userService from "../../services/user.service";
import UserAutocomplete from "../Post/UserAutocomplete";
import {Button, Dropdown, DropdownButton, InputGroup, Modal} from "react-bootstrap";

import "./../../style/chat.css"
import toastService from "../../services/toast.service";
import chatService from "../../services/chat.service";
import postService from "../../services/post.service";
import PostPreviewModal from "../Post/PostPreviewModal";
import PostPreviewThumbnail from "../Post/PostPreviewThumbnail";
import storyService from "../../services/story.service";
import Story from "../StoryCompoent/Story";


function Chats() {
    const [users, setUsers] = useState([]);
    const dispatch = useDispatch();
    const store = useSelector(state => state);
    const [selectedUser, setSelectedUser] = useState({});
    const [messageText, setMessageText] = useState("");
    const [chatRoom, setChatRoom] = useState("");
    const [conn, setConn] = useState({});
    const [messages, setMessages] = useState([]);
    const [renderMessages, setRenderMessages] = useState([]);
    const [messageCategory, setMessageCategory] = useState("Message category");
    const [showModalPost, setShowModalPost] = useState(false);
    const [showModalStory, setShowModalStory] = useState(false);
    const [selectedPost, setSelectedPost] = useState({});
    const [showModal, setShowModal] = useState(false);
    const [modalImage, setModalImage] = useState({});

    useEffect(() => {
        getAllUsers();
    }, []);

    const openPost = async (postId) => { //samo postId dobijes
        const response = await postService.getPostById({id : postId, jwt: store.user.jwt});
        if (response.status === 200) {
            setSelectedPost(response.data)
            setShowModalPost(true);
        }
    }

    useEffect(() => {
        if (conn !== {}) {
            // conn.onmessage()
            // gledaj da budu javan i da su follow
            conn.onmessage = function(event) {
                let eventData = JSON.parse(event.data);
                if (eventData.ContentType === "Image") { //mora zbog slike za pravilno seenovanje, jer slika nema Id ako se ne getuje
                    getMessagesForChatRoom(chatRoom.Id);
                }
                else
                    updateChatBox(event);
            };
        }
    }, [conn, renderMessages]);

    async function getAllUsers() {
        const response = await userService.getAllUsers({ jwt: store.user.jwt });
        await setUsers(response.data.users);
    }

    async function updateChatBox(event) {
        console.log("event")
        console.log(JSON.parse(event.data))

        let temp = messages;
        temp.push(JSON.parse(event.data));
        await setMessages(temp);
        await updateRenderMessages(temp);

        // const response = await chatService.GetMessagesForChatRoom({
        //     roomId : chatRoom.Id,
        //     jwt : store.user.jwt
        // });
        // if (response.status === 200) {
        //     let temp = response.data;
        //     temp.push(event)
        //     await setMessages(temp)
        //     await updateRenderMessages(temp);
        // }
        // else
        //     toastService.show("error", "Something went wrong. Try again")
    }

    async function getMessagesForChatRoom(roomId) {
        const response = await chatService.GetMessagesForChatRoom({
            roomId : roomId,
            jwt : store.user.jwt
        });
        if (response.status === 200) {
            console.log("messages");
            console.log(response.data);
            toastService.show("success", "Chat room messages retrieved successfully")
            await setMessages(response.data);
            await updateRenderMessages(response.data);
        }
        else
            toastService.show("error", "Something went wrong. Try again")
    }

    async function updateRenderMessages(messages) {
        await setRenderMessages([]);
        let tempList = [];
        for (const message of messages) {
            let tempMessage = message;
            if (message.ContentType === "Post") {
                const response = await postService.getPostById({id : message.Content, jwt: store.user.jwt});
                if (response.status === 200)
                    tempMessage.Content = response.data;
                else
                    tempMessage.Content = "Cannot view this post";
            }
            else if (message.ContentType === "Story") {
                const response = await storyService.getStoryById({id : message.Content, jwt: store.user.jwt});
                if (response.status === 200) {
                    let story = [response.data];
                    tempMessage.Content =  {
                        stories: story
                    };
                }
                else
                    tempMessage.Content = "Cannot view this story";
            }
            tempList.push(tempMessage)
        }
        await setRenderMessages(tempList)
    }

    async function startChat() {
        await setMessages([]);
        await setRenderMessages([]);
        await setConn({});
        await setChatRoom({});

        const response = await chatService.StartConversation({
            person1: store.user.id,
            person2: selectedUser.id,
            jwt : store.user.jwt
        });
        if (response.status === 200) {
            toastService.show("success", "Chat room retrieved successfully")
            console.log("chat room")
            console.log(response.data)
            await setChatRoom(response.data)
            await setConn(new WebSocket("ws://localhost:8003" + "/ws/" + response.data.Id));
            await getMessagesForChatRoom(response.data.Id)
        }
        else
            toastService.show("error", "Something went wrong. Try again")
    }

    function sendMessage() {
        if (messageCategory === "Message category") {
            toastService.show("error", "Cannot send empty message");
            return;
        }
        conn.send(JSON.stringify({SenderId : store.user.id, ReceiverId : selectedUser.id, RoomId : chatRoom.Id, Content : messageText, ContentType : messageCategory}));
        setMessageCategory("Message category");
        setMessageText("");
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

    function handleModal(message) {
        console.log(message)
        setModalImage(message.Content)
        setShowModal(!showModal)
        seenPhoto(message)
    }

    async function seenPhoto(data) {
        //ako sam ja poslao i otvorio onda ne treba to da se desi
        console.log(data);
        if (data.SenderId === store.user.id)
            return;

        const response = await chatService.seenPhoto({
            Id: data.Id,
            jwt: store.user.jwt
        });
        if (response.status === 200) {
            getMessagesForChatRoom(chatRoom.Id);

        } else
            toastService.show("error", "Something went wrong. Try again")
    }

    function closeModal() {
        setShowModal(!showModal)
    }


    return (
        <div style={{marginTop:'5%', marginLeft : "5%"}}>
            <Navigation/>
            <h1>Chat</h1>
            <UserAutocomplete setSelectedUser={setSelectedUser} suggestions={users} />
            <Button style={{marginLeft : "1%"}} onClick={startChat}>Start chat</Button>

            <br/><br/><br/>

            <div style={{overflow: "scroll", height:"600px"}}>
                {renderMessages.map((message) => {
                    return (
                        <div>
                            {message.SenderId === store.user.id && <div className="containerChat">
                                {message.ContentType === "String" && <p>{message.Content}</p>}

                                {message.ContentType === "Image" && <Button disabled={message.SenderId === store.user.id ? false : message.IsMediaOpened}
                                    style={{marginLeft: '5px', marginTop: '22px', height: '32px', fontSize: '15px'}}
                                    variant="success"  onClick={() => handleModal(message)}>Photo </Button>
                                }

                                {message.ContentType === "Post" && message.Content !== "Cannot view this post" &&
                                    <PostPreviewThumbnail post={message.Content} openPost={openPost} />
                                }
                                {message.ContentType === "Post" && message.Content === "Cannot view this post" &&
                                <p style={{color : "red"}}>{message.Content}</p>
                                }


                                {message.ContentType === "Story" && message.Content !== "Cannot view this story" &&
                                <Story story={message.Content} iconSize={"xxl"} hideUsername={true} />
                                }
                                {message.ContentType === "Story" && message.Content === "Cannot view this story" &&
                                <p style={{color : "red"}}>{message.Content}</p>
                                }

                                <span style={{marginLeft: "20px"}} className="time-rightChat">{message.DateCreated}</span>
                            </div>
                            }




                            {message.ReceiverId === store.user.id && <div className="containerChat darkerChat">

                                {message.ContentType === "String" && <p>{message.Content}</p>}

                                {message.ContentType === "Image" && <Button disabled={message.SenderId === store.user.id ? false : message.IsMediaOpened}
                                                                            style={{marginLeft: '5px', marginTop: '22px', height: '32px', fontSize: '15px'}}
                                                                            variant="success"  onClick={() => handleModal(message)}>Photo </Button>
                                }

                                {message.ContentType === "Post" && message.Content !== "Cannot view this post" &&
                                <PostPreviewThumbnail post={message.Content} openPost={openPost} />
                                }
                                {message.ContentType === "Post" && message.Content === "Cannot view this post" &&
                                <p style={{color : "red"}}>{message.Content}</p>
                                }


                                {message.ContentType === "Story" && message.Content !== "Cannot view this story" &&
                                <Story story={message.Content} iconSize={"xxl"} hideUsername={true} />
                                }
                                {message.ContentType === "Story" && message.Content === "Cannot view this story" &&
                                <p style={{color : "red"}}>{message.Content}</p>
                                }

                                <span style={{marginLeft: "20px"}} className="time-leftChat">{message.DateCreated}</span>
                            </div>
                            }


                        </div>
                    )
                })}
            </div>

            <DropdownButton onSelect={(event) => handleMessageCategoryChange(event) } as={InputGroup.Append}  variant="outline-secondary" title={messageCategory} id="input-group-dropdown-2" >
                <Dropdown.Item eventKey={"String"} >String</Dropdown.Item>
                <Dropdown.Item eventKey={"Image"} >Image</Dropdown.Item>
            </DropdownButton>

            {(messageCategory === "String") &&
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

            { showModalPost &&
            <PostPreviewModal
                post={selectedPost}
                postUser={{ id: selectedPost.userId }}
                showModal={showModalPost}
                setShowModal={setShowModalPost}
            /> }


            <Modal show={showModal} onHide={closeModal}>
            <Modal.Header closeButton>
            <Modal.Title>Photo</Modal.Title>
            </Modal.Header>
            <Modal.Body>
            <img src={modalImage}  style={{width:'400px', height: '400px'}}/>
            </Modal.Body>
            </Modal>
        </div>
    );
}


export default Chats;