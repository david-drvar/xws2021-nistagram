import React, { useEffect, useState } from 'react';
import { useSelector } from "react-redux";
import Navigation from "../HomePage/Navigation";
import { Button, Modal, Dropdown } from "react-bootstrap";
import UserAutocomplete from "../Post/UserAutocomplete";
import ProfileForAutocomplete from "../Post/ProfileForAutocomplete";
import AutocompleteHashtags from "../Post/AutocompleteHashtags";
import userService from "../../services/user.service";
import postService from "../../services/post.service";
import toastService from "../../services/toast.service";
import hashtagService from "../../services/hashtag.service";
import Switch from 'react-switch'
import storyService from '../../services/story.service';

function NewPost(props) {
    const [user, setUser] = useState({});

    const [description, setDescription] = useState('');
    const [location, setLocation] = useState('');
    const [image, setImage] = useState('');
    const [hashtagList, setHashtagList] = useState([]);
    const [hashtagListForPrint, setHashtagListForPrint] = useState([]);
    const [showModal, setShowModal] = useState(false);
    const [tagList, setTagList] = useState([]);
    const [tagListForPrint, setTagListForPrint] = useState([]);
    const [mediaList, setMediaList] = useState([]);
    const [postPrint, setPostPrint] = useState([]);
    const [imageName, setImageName] = useState("");
    const [allUsers, setAllUsers] = useState([]);
    const [allHashtags, setAllHashtags] = useState([]);
    const [closeFriends, setCloseFriends] = useState(false);
    const [isStory, setIsStory] = useState(false); 

    const store = useSelector(state => state);

    useEffect(() => {
        getUserInfo();
        getAllUsers();
        getAllHashtags();
    }, []);

    async function getUserInfo() {
        const response = await userService.getUserById({
            id: store.user.id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUser(response.data)
        } else {
            console.log("getuser error")
        }
    }

    async function getAllUsers() {
        const response = await userService.getAllUsers({ jwt: store.user.jwt });
        await setAllUsers(response.data.users);
    }

    async function getAllHashtags() {
        const response = await hashtagService.getAllHashtags({ jwt: store.user.jwt });
        setAllHashtags(response.data.hashtags)
    }

    function handleTagAutocompleteClick(tag) {
        if (tagListForPrint.some((someTag) => someTag.id === tag.id)) //prevents adding the same tag
            return;
        // setTagList([...tagList, {userId: tag.userId, username: tag.username, mediaId: "1"}]);
        setTagListForPrint([...tagListForPrint, tag]);
        setTagList([...tagList, { userId: tag.id, mediaId: "1", username: tag.username }]);
    }

    function handleHashtagAutocompleteClick(tag) {
        if (hashtagListForPrint.some((someTag) => someTag.id === tag.id)) //prevents adding the same tag
            return;
        // setTagList([...tagList, {userId: tag.userId, username: tag.username, mediaId: "1"}]);
        setHashtagListForPrint([...hashtagListForPrint, tag]);
        setHashtagList([...hashtagList, { id: tag.id, text: tag.text }]);
    }

    function handleHashtagAutocompleteNewSuggestion(newTag) {
        setHashtagListForPrint([...hashtagListForPrint, { text: newTag }]);
        setHashtagList([...hashtagList, { text: newTag }]);
    }

    const postDetails = async () => {
        if (mediaList.length === 0) {
            toastService.show("warning", "Please add media for post")
            return;
        }
        let date = new Date();
        let month = date.getMonth() + 1;
        if (month < 10) month = "0" + month;
        const jsonDate = date.getFullYear() + "-" + month + "-" + date.getDate() + "T01:30:15.01Z";

        const contentRequest = {
            id: "1",
            userId: user.id,
            isAd: false,
            type: isStory ? "Story" : "Post",
            description: description,
            location: location,
            createdAt: jsonDate,
            media: mediaList,
            hashtags: hashtagList,
            jwt: store.user.jwt
        };
        if(isStory) contentRequest["isCloseFriends"] = closeFriends;

        let response = {} 
        isStory ? response = await storyService.createStory(contentRequest) : 
                  response = await postService.createPost(contentRequest)

        if (response.status === 200)
            toastService.show("success", `New ${isStory ? "story" : "post"} successfully created!`);
        else
            toastService.show("error", "Something went wrong, please try again!");
    }

    function handleChangeImage(evt) {
        console.log("Uploading");
        setImageName(evt.target.files[0].name);
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function (upload) {
            setImage(upload.target.result)
        };
        reader.readAsDataURL(file);
    }

    function handleModal() {
        setShowModal(!showModal)
    }

    function closeModal() {
        setTagList([]);
        setTagListForPrint([]);
        setImage("");
        setShowModal(!showModal)
    }

    async function saveModal() {
        let tagListFilter = await tagList.filter(tag => tag.userId !== "");
        let media = {
            id: "1",
            type: "Image",
            postId: "1",
            content: image,
            orderNum: mediaList.length + 1,
            tags: tagListFilter
        };
        setMediaList([...mediaList, media]);
        setPostPrint([...postPrint, { filename: imageName, tags: tagListFilter }]);

        setTagList([]);
        setTagListForPrint([]);
        setImage("");
        closeModal();
    }

    return (
        <div className='home'>
            <Navigation user={user} />

            <div className="card input-filed"
                style={{ margin: "30px auto", maxWidth: "500px", padding: "20px", textAlign: "center", marginTop: "5%" }} >
                <Dropdown style={{marginBottom: "1em"}}>
                    <Dropdown.Toggle variant="link" id="dropdown-basic">
                        { isStory ? "Story" : "Post" }
                    </Dropdown.Toggle>

                    <Dropdown.Menu>
                        <Dropdown.Item onClick={() => setIsStory(false)}> Post </Dropdown.Item>
                        <Dropdown.Item onClick={() => setIsStory(true)}> Story </Dropdown.Item>
                    </Dropdown.Menu>
                </Dropdown>

                <input type="text" placeholder="description" value={description} onChange={(e) => setDescription(e.target.value)} />
                <br />
                <input type="text" placeholder="location" value={location} onChange={(e) => setLocation(e.target.value)} />
                <br/>
                { isStory && <div className='row'>
                    <p style={{ color: '#64f427' }}>Close friends: </p>
                    <Switch onChange={() => setCloseFriends(!closeFriends)} checked={closeFriends} />
                </div >}
                <br />
                <AutocompleteHashtags addToHashtaglist={handleHashtagAutocompleteClick}
                    suggestions={allHashtags} handleHashtagAutocompleteNewSuggestion={handleHashtagAutocompleteNewSuggestion}
                />
                <br/>
                <h3>Hashtags:</h3>
                <div>
                    <ul>
                        {hashtagListForPrint.map((hashtag, i) => {
                            return (
                                <li>
                                    {hashtag.text}
                                </li>
                            );
                        })}
                    </ul>
                    {/*<div style={{ marginTop: 20 }}>{JSON.stringify(hashtagList)}</div>*/}
                </div>
                <br /><br />
                <Button type={"outline-primary"} onClick={handleModal} style={{ maxWidth: "150px", textAlign: "center" }}
                >Add file</Button>
                {postPrint.map((x, i) => {
                    return (
                        <div className="box">
                            Filename: {x.filename} <br />
                            Tag number: {x.tags.length}
                            <br /><br />
                        </div>

                    );
                })}
                <br /><br />
                <Button type={"primary"} onClick={() => postDetails()}>Submit post</Button>

            </div>
            <Modal show={showModal} onHide={closeModal} style={{ 'height': 650 }} >
                <Modal.Header closeButton style={{ 'background': 'silver' }}>
                    <Modal.Title>Post files</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{ 'background': 'silver' }}>
                    <input type="file" name="file"
                        className="upload-file"
                        id="file"
                        onChange={handleChangeImage}
                        formEncType="multipart/form-data"
                        required />
                    <br /><br />
                    <UserAutocomplete addToTaglist={handleTagAutocompleteClick} suggestions={allUsers} />
                    <h3>Tags:</h3>
                    <div>
                        <ul>
                            {tagListForPrint.map((tag, i) => {
                                return (
                                    <li>
                                        <ProfileForAutocomplete username={tag.username} firstName={tag.firstName} lastName={tag.lastName}
                                            caption={tag.biography} urlText="Follow" iconSize="medium" captionSize="small" storyBorder={true} />
                                    </li>
                                );
                            })}
                        </ul>
                        <br /><br />
                        <Button type={"primary"} onClick={() => saveModal()}>Save</Button>
                    </div>
                </Modal.Body>
            </Modal>
        </div>
    );
}

export default NewPost;