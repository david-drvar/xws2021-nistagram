import moment from 'moment';
import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux'
import { Modal, Button, DropdownButton, Dropdown } from 'react-bootstrap'

import Navigation from './../HomePage/Navigation';
import AutocompleteHashtags from "../Post/AutocompleteHashtags";
import UserAutocomplete from "../Post/UserAutocomplete";
import ProfileForAutocomplete from './../Post/ProfileForAutocomplete'

import adsService from '../../services/ads.service';
import toastService from '../../services/toast.service';

import './../../style/createCampaign.css'
import campaignsService from '../../services/campaigns.service';
import hashtagService from '../../services/hashtag.service';
import userService from '../../services/user.service';

const CreateCampaign = (props) => {
    const [newCampaign, setNewCampaign] = useState({
        name: "", isOneTime: false, startDate: new Date(), endDate: new Date(), agentId: "", category: { id: "", name: "" }, type: "", ads: []
    });
    const [newPost, setNewPost] = useState({
        type: "", description: "", location: "", media: [], hashtags: [], isAd: true, userId: "",
    })
    const [newAd, setNewAd] = useState({ link: "", post: {...newPost} })
    const [newMedia, setNewMedia] = useState({
        tags: [], type: "", content: "", orderNum: ""
    })
    const [categories, setCategories] = useState([]);
    const [allHashtags, setAllHashtags] = useState([]);
    const [allUsers, setAllUsers] = useState([]);
    const [showFilesModal, setShowFilesModal] = useState(false)

    const store = useSelector(state => state)

    useEffect(() => {
        getAllUsers()
        getCategories()
        getAllHashtags()
    }, [])

    const getCategories = async () => {
        const response = await adsService.getAdCategories({ jwt: store.user.jwt })
        if (response.status === 200) {
            setCategories([...response.data.categories])
            setNewCampaign({ ...newCampaign, type: "Post", category: response.data.categories[0] })
        } else {
            toastService.show("error", "Could retrieve campaign categories")
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

    function handleHashtagAutocompleteClick(tag) {
        if (newPost.hashtags.some((someTag) => someTag.id === tag.id)) return;
        setNewPost({ ...newPost, hashtags: [...newPost.hashtags, { text: tag }] });
    }

    function handleHashtagAutocompleteNewSuggestion(newTag) {
        setNewPost({ ...newPost, hashtags: [...newPost.hashtags, { text: newTag }] });
    }

    function handleChangeImage(evt) {
        console.log("Uploading");
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function (upload) {
            setNewMedia({...newMedia, content: upload.target.result})
        };
        reader.readAsDataURL(file);
    }

    function handleTagAutocompleteClick(tag) {
        if (newMedia.tags.some((someTag) => someTag.id === tag.id)) return;
        setNewMedia({...newMedia, tags: [...newMedia.tags, { userId: tag.id, mediaId: "1", username: tag.username }]});
    }

    async function saveMediaModal() {
        let media = {
            ...newMedia,
            type: "Image",
            orderNum: newPost.media.length + 1
        };
        setNewPost({...newPost, media: [...newPost.media, media]});
        setNewMedia({ tags: [], type: "", content: "", orderNum: "" })
        setShowFilesModal(false)
    }

    const saveAdToCampaign = () => {
        if(newPost.media.length === 0) return
        setNewCampaign({
                ...newCampaign, 
                ads: [
                    ...newCampaign.ads, 
                    {...newAd, post: {...newPost, userId: store.user.id, isAd: true, type: newCampaign.type}}
                ]
            })
        setNewPost({type: "", description: "", location: "", media: [], hashtags: [], tags: [], userId: "", isAd: true})
        setNewAd({ link: "" })
    }

    const updatePostType = (type) => {
        const ads = [...newCampaign.ads]
        ads.map(ad => {
            return {
                ...ad,
                post: {
                    ...ad.post,
                    type: type,
                }
            }
        })

        setNewCampaign({ ...newCampaign, type: type, ads: [...ads] })
    }

    const createCampaign = async () => {
        console.log(newCampaign)
        if(newCampaign.ads.length === 0) return;

        const response = await campaignsService.createCampaign({
            campaign: { 
                id: "",
                ...newCampaign, 
                agentId: store.user.id,
                startDate: newCampaign.startDate + "T02:00:00.00Z",
                endDate: (newCampaign.isOneTime ? newCampaign.startDate : newCampaign.endDate) + "T02:00:00.00Z",
            },
            jwt: store.user.jwt
        })
        if (response.status === 200) {
            toastService.show("success", "Successfuly create new campaign!")
            setTimeout(() => {
                window.location.reload()
            }, 3000)
        } else {
            toastService.show("error", "Could not create new campaign.")
        }
    }

    return (
        <div>
            <Navigation />
            <div className="CreateCampaign__Wrapper">
                <div className="title">
                    Create New Campaign
                </div>
                <div className="updateInput">
                    <label for="name">Name</label>
                    <input
                        type="text" name="name"
                        placeholder={"Enter campaign name..."}
                        value={newCampaign.name}
                        onChange={(e) => setNewCampaign({ ...newCampaign, name: e.target.value })}
                        className="form-control" id="name" />
                </div>
                <div className="dropdown">
                    <label for="category">Campaign Duration</label>
                    <DropdownButton id="dropdown-basic-button" variant="outline-primary" title={newCampaign.isOneTime ? "One time" : "Long term"} style={{ width: "10em" }}>
                        <Dropdown.Item onClick={() => setNewCampaign({...newCampaign, isOneTime: true})}>One time</Dropdown.Item>
                        <Dropdown.Item onClick={() => setNewCampaign({...newCampaign, isOneTime: false})}>Long term</Dropdown.Item>
                    </DropdownButton>
                </div>
                <div className="updateInput">
                    <label for="startDate">Start Date</label>
                    <input
                        min={moment().format("MM/DD/YYYY")}
                        max={moment(new Date()).add(365, 'd').toDate()}
                        type="date" name="startDate"
                        value={newCampaign.startDate}
                        onChange={(e) => setNewCampaign({ ...newCampaign, startDate: e.target.value })}
                        className="form-control" id="startDate" />
                </div>
                {!newCampaign.isOneTime && <div className="updateInput">
                    <label for="endDate">End Date</label>
                    <input
                        min={newCampaign.startDate}
                        max={moment(new Date()).add(365, 'd').toDate()}
                        type="date" name="endDate"
                        value={newCampaign.endDate}
                        onChange={(e) => setNewCampaign({ ...newCampaign, endDate: e.target.value })}
                        className="form-control" id="endDate" />
                </div>}
                <div className="dropdown">
                    <label for="category">Category</label>
                    <DropdownButton id="dropdown-basic-button" variant="outline-primary" title={newCampaign.category.name + " "} style={{ width: "10em" }}>
                        {categories.map(category => <Dropdown.Item onClick={() => setNewCampaign({ ...newCampaign, category: category })}>{category.name}</Dropdown.Item>)}
                    </DropdownButton>
                </div>
                <div className="dropdown">
                    <label for="category">Campaign Type</label>
                    <DropdownButton id="dropdown-basic-button" variant="outline-primary" title={newCampaign.type + " "} style={{ width: "10em" }}>
                        <Dropdown.Item onClick={() => updatePostType("Post")}>Post</Dropdown.Item>
                        <Dropdown.Item onClick={() => updatePostType("Story")}>Story</Dropdown.Item>
                    </DropdownButton>
                </div>
                <div className="updateInput">
                    <label>Currently {newCampaign.ads.length} ads.</label>
                    <Button 
                        disabled={newCampaign.ads.length === 0}
                        onClick={() => createCampaign()}>Create Campaign</Button>
                </div>
                <div className="createNewAd">
                    <div className="title">Create new {newCampaign.type} Ad </div>
                    <div className="updateInput">
                        <label for="link">Link</label>
                        <input
                            type="text" name="link"
                            placeholder={"Enter sponsored link..."}
                            value={newPost.link}
                            onChange={(e) => setNewAd({ ...newAd, link: e.target.value })}
                            className="form-control" id="link" />
                    </div>
                    <div className="updateInput">
                        <label for="description">Description</label>
                        <input
                            type="text" name="description"
                            placeholder={"Enter post description..."}
                            value={newPost.description}
                            onChange={(e) => setNewPost({ ...newPost, description: e.target.value })}
                            className="form-control" id="description" />
                    </div>
                    <div className="updateInput">
                        <label for="location">Location</label>
                        <input
                            type="text" name="location"
                            placeholder={"Enter post location..."}
                            value={newPost.location}
                            onChange={(e) => setNewPost({ ...newPost, location: e.target.value })}
                            className="form-control" id="location" />
                    </div>
                    <div className="updateInput">
                        <label for="location">Hashtags</label>
                        <AutocompleteHashtags
                            addToHashtaglist={handleHashtagAutocompleteClick}
                            handleHashtagAutocompleteNewSuggestion={handleHashtagAutocompleteNewSuggestion}
                            suggestions={allHashtags}
                        />
                    </div>
                    <ul className="updateInput">
                        {newPost.hashtags.map(hashtag => (
                            <li>{hashtag.text}</li>
                        ))}
                    </ul>
                    <div className="updateInput">
                        <label for="files">Add media to {newCampaign.type}</label>
                        <Button name="files" type="outline-primary" onClick={() => setShowFilesModal(!showFilesModal)}>Add</Button>
                        <Modal show={showFilesModal} onHide={() => setShowFilesModal(!showFilesModal)} style={{ 'height': 650 }} >
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
                                        {newMedia.tags.map((tag, i) => {
                                            return (
                                                <li>
                                                    <ProfileForAutocomplete username={tag.username} firstName={tag.firstName} lastName={tag.lastName}
                                                        caption={tag.biography} urlText="Follow" iconSize="medium" captionSize="small" storyBorder={true} />
                                                </li>
                                            );
                                        })}
                                    </ul>
                                    <br /><br />
                                    <Button type={"primary"} onClick={() => saveMediaModal()}>Save</Button>
                                </div>
                            </Modal.Body>
                        </Modal>
                    </div>
                    <div className="updateInput">
                        <label for="files">You have {newPost.media.length} media in this post.</label>
                        <Button onClick={() => saveAdToCampaign()}>Save {newCampaign.type} Ad </Button> 
                    </div>
                </div>
            </div>
        </div>
    );
}

export default CreateCampaign;