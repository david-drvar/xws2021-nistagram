import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Button, Modal } from "react-bootstrap";
import { useParams } from 'react-router-dom'
import { ReactComponent as VerificationSymbol } from "../../images/icons/verification-symbol.svg";

import FollowAndUnfollow from "./FollowAndUnfollow";
import Navigation from "../HomePage/Navigation";
import { userActions } from "../../store/actions/user.actions";
import FollowersAndFollowings from "./FollowersAndFollowings";
import BlockMuteAndNotifications from "./BlockMuteAndNotifications";
import Highlight from './../StoryCompoent/Highlight';
import PostPreviewGrid from './../Post/PostPreviewGrid';
import Spinner from './../../helpers/spinner';
import Story from './../StoryCompoent/Story';

import userService from "../../services/user.service";
import privacyService from "../../services/privacy.service";
import followersService from "../../services/followers.service";
import postService from './../../services/post.service';
import storyService from './../../services/story.service';
import highlightsService from './../../services/highlights.service';
import toastService from './../../services/toast.service';

import '../../style/Profile.css';


const Profile = () => {
    const {username} = useParams()

    const [loadingPosts, setLoadingPosts] = useState(true);
    const [loadingHighlights, setLoadingHighlights] = useState(true);

    const [user, setUser] = useState({});
    const [follow, setFollow] = useState({});
    const [publicProfile, setPublicProfile] = useState(false);

    const [showModalFollowers, setModalFollowers] = useState(false);
    const [showModalFollowings, setModalFollowings] = useState(false);
    const [followers, setFollowers] = useState([]);
    const [following, setFollowings] = useState([]);

    const [closeFriend, setCloseFriend] = useState(false);
    const [isApprovedRequest, setIsApprovedRequest] = useState(false);
    const [isMuted, setIsMuted] = useState(false);
    const [notifications, setNotifications] = useState({
        isMessageNotificationEnabled: false,
        isPostNotificationEnabled: false,
        isStoryNotificationEnabled: false,
        isCommentNotificationEnabled: false
    })

    const [posts, setPosts] = useState([]);
    const [stories, setStories] = useState([]);
    const [highlights, setHighlights] = useState([]);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        (async function () {
            const tempUser = await getUserByUsername(); // Since it doesn't get saved in time for other requests
            if (tempUser) {
                getUserPrivacy(tempUser.id);
                getFollowers(tempUser.id)
                getFollowing(tempUser.id)
                checkUser(tempUser.id);
                getUserPrivacy(tempUser.id);
                getFollowers(tempUser.id);
                getFollowing(tempUser.id);
                getPosts(tempUser.id);
                getStories(tempUser);
                getHighlights(tempUser.id);
            }
        })();
    }, [username]);

    const getPosts = async (userId) => {
        const response = await postService.getPostsForUser({
            jwt: store.user.jwt,
            userId: userId
        })

        if (response.status === 200) {
            setPosts([...response.data.posts])
            setLoadingPosts(false);
        } else {
            console.log(response);
            toastService.show("error", "Could not retrieve user's posts.")
        }
    }

    const getStories = async (user) => {
        const response = await storyService.getUsersStories({
            jwt: store.user.jwt,
            userId: user.id
        })
        
        if (response.status === 200){ 
            setStories({
                username: user.username,
                userPhoto: user.profilePhoto,
                stories: [...response.data.stories]
            })
        }
        else{
            console.log(response);
            toastService.show("error", "Could not retrieve user's stories.")
        }
    }

    const getHighlights = async (userId) => {
        const response = await highlightsService.getUserHighlights({
            jwt: store.user.jwt,
            userId: userId
        })

        console.log(response);

        if (response.status === 200) {
            setHighlights([...response.data.highlights])
            setLoadingHighlights(false);
        } else {
            console.log(response);
            toastService.show("error", "Could not retrieve user's highlights.")
        }
    }

    async function getUserByUsername() {
        const response = await userService.getUserByUsername({
            username: username,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUser(response.data)
            console.log(response.data)
            return response.data
        } else {
            console.log("getuserbyusername error")
        }
    }

    async function checkUser(value) {
        if (value === store.user.id) {
            setFollow(false)
        } else {
            setFollow(true)
            dispatch(userActions.followRequest({
                userId: store.user.id,
                followerId: value,
            }))
            isCloseFriend(value)
        }
    }

    async function isCloseFriend(value) {
        const response = await followersService.getFollowersConnection({
            userId: store.user.id,
            followerId: value,
        })
        if (response.status === 200) {
            setCloseFriend(response.data.isCloseFriends)
            setIsApprovedRequest(response.data.isApprovedRequest)
            setIsMuted(response.data.isMuted)
            setNotifications({
                ...notifications,
                isMessageNotificationEnabled: response.data.isMessageNotificationEnabled,
                isPostNotificationEnabled: response.data.isPostNotificationEnabled,
                isStoryNotificationEnabled: response.data.isStoryNotificationEnabled,
                isCommentNotificationEnabled: response.data.isCommentNotificationEnabled
            });

        } else {
            console.log("followings ne radi")
        }
    }

    async function getUserPrivacy(value) {
        const response = await privacyService.getUserPrivacy({
            userId: value,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setPublicProfile(response.data.response)
        } else {
            console.log("privacy ne radi")
        }
    }

    async function getFollowing(value) {
        const response = await followersService.getFollowing({
            userId: value,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setFollowings(response.data.users);
        } else {
            console.log("followings ne radi")
        }
    }

    async function getFollowers(value) {
        const response = await followersService.getFollowers({
            userId: value,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setFollowers(response.data.users);
        } else {
            console.log("followers ne radi")
        }
    }

    function handleModalFollowers() {
        if (publicProfile || isApprovedRequest || !follow)
            setModalFollowers(!showModalFollowers)
    }

    function handleModalFollowings() {
        if (publicProfile || isApprovedRequest || !follow)
            setModalFollowings(!showModalFollowings)
    }

    return (
        <div>
            <Navigation/>
            <div className="profileGrid">
                    <div className="profileHeader">
                        { stories.stories && stories.stories.length > 0 && 
                        (!follow || // Moj profil
                        (follow && publicProfile) || // Tudji javan 
                        (follow && !publicProfile  && isApprovedRequest)) ?
                            <Story story={stories} iconSize={"xxl"} hideUsername={true} /> :
                            <img style={{marginLeft: "-1em", paddingRight: "4px"}} alt="" src={user.profilePhoto}/>
                        }
                        <div className="info">
                            <div className="fullname">
                                {user.firstName} {user.lastName} 
                                {user.role === "Verified" && <span><VerificationSymbol style={{width: "20px", height: "20px", marginLeft: "10px", display: "inline-block"}} fill="#0095f6" /></span>}
                                {follow && <span className="blockMute">
                                    <BlockMuteAndNotifications 
                                        isApprovedRequest={isApprovedRequest} isMuted={isMuted} notifications={notifications} />
                                </span>
                            }
                        </div>
                        <div className="username">@{user.username}</div>
                        <div className="stats">
                            <div class="single-stat postsNum"><strong>{posts.length ? posts.length : 0}</strong> posts
                            </div>
                            <div class="single-stat">
                                <Button variant="link" style={{color: 'black'}} onClick={handleModalFollowers}>
                                    <strong>{followers.length}</strong> followers
                                </Button>
                            </div>
                            <div class="single-stat">
                                <Button variant="link" style={{color: 'black'}} onClick={handleModalFollowings}>
                                    <strong>{following.length}</strong> following
                                </Button>
                            </div>
                        </div>
                        {user.biography && <div>{user.biography}</div>}
                        {user.website &&
                        <a className="website" target="_blank" rel="noreferrer"
                           href={user.website.includes('https://') ? user.website : `https://${user.website}`}>
                            {user.website}
                        </a>}
                        {follow &&
                        <FollowAndUnfollow className="followUnfollow" user={user} isCloseFriends={closeFriend}
                                           funcIsCloseFriend={isCloseFriend}
                                           followers={followers}
                                           getFollowers={getFollowers}
                        />}
                    </div>
                </div>

                <div className="content">
                    {(!follow || // Moj profil 
                    (follow && publicProfile) || // Tudji javan 
                    (follow && !publicProfile && isApprovedRequest)) && // Tudji privatan koji ja pratim
                    (<div className="highlights">
                        {loadingHighlights ?
                            <div style={{position: "relative", left: "45%", marginTop: "50px"}}>
                                <Spinner type="MutatingDots" height="100" width="100"/>
                            </div> :
                            highlights.map(highlight => {
                                highlight["profileImage"] = highlight.stories.length > 0 ? highlight.stories[0].media[0].content : ""; // check
                                return <Highlight highlight={highlight}/>
                            })
                        }
                    </div>)
                    }

                    {(!follow || // Moj profil
                    (follow && publicProfile) || // Tudji javan
                    (follow && !publicProfile && isApprovedRequest)) && // Tudji privatan koji ja pratim
                        <div className="posts">
                            {loadingPosts ?
                                <div style={{position: "relative", left: "45%", marginTop: "50px"}}>
                                    <Spinner type="MutatingDots" height="100" width="100"/>
                                </div> :
                                <PostPreviewGrid posts={posts}/>
                            }
                        </div>
                    }

                    {(follow && !publicProfile && !isApprovedRequest) && // Tudji koji je privatan i nije approve-ovan
                    <div style={{borderTop: '1px solid black'}}>
                        <p style={{textAlign: 'center', marginTop: '6%', fontWeight: 'bold'}}> This Account is
                            Private </p>
                        <p style={{textAlign: 'center', marginTop: '2%'}}>Follow to see their photos and videos!</p>
                    </div>
                    }
                </div>

                <Modal show={showModalFollowers} onHide={handleModalFollowers}>
                    <Modal.Header closeButton>
                        <Modal.Title>Followers:</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <FollowersAndFollowings ids={followers} handleModal={handleModalFollowers}/>
                    </Modal.Body>
                </Modal>
                <Modal show={showModalFollowings} onHide={handleModalFollowings}>
                    <Modal.Header closeButton>
                        <Modal.Title>Following:</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <FollowersAndFollowings ids={following} handleModal={handleModalFollowings}/>
                    </Modal.Body>
                </Modal>

            </div>
        </div>
    );
}
export default Profile;
