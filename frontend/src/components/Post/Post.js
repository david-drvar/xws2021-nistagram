import React, { useEffect, useState, useRef } from "react";
import { NavLink } from "react-router-dom";
import { useSelector } from 'react-redux';
import moment from 'moment';
import { v4 as uuidv4 } from 'uuid';
import "../../style/post.css";

import Slider from './Slider'
import Comment from "./Comment";
import PostHeader from './PostHeader';
import CollectionsModal from './CollectionsModal';
import {Button, Dropdown, Modal} from "react-bootstrap";
import { ReactComponent as CardButton } from './../../images/icons/cardButton.svg' 
import PostMenu from "./PostMenu";

import userService from './../../services/user.service';
import toastService from './../../services/toast.service';
import likeService from './../../services/like.service';
import commentService from './../../services/comment.service';
import favoritesService from './../../services/favorites.service';
import complaintService from "../../services/complaint.service";

function Post (props) {
    const { postUser, shouldReload } = props;
    const [post, setPost] = useState(props.post)
    const [user, setUser] = useState({});
    const [hoursAgo, setHoursAgo] = useState(0)
    const [daysAgo, setDaysAgo] = useState(0);
    const [minutesAgo, setMinutesAgo] = useState(0)

    const [showSaveModal, setShowSaveModal] = useState(false);
    const [isSaved, setIsSaved] = useState(false);
    const [collections, setCollections] = useState([]);
    const [savedInCollections, setSavedInCollections] = useState([]);

    const [likesText, setLikesText] = useState("");
    const [dislikesText, setDislikesText] = useState("");
    const [isLiked, setIsLiked] = useState(false);
    const [isDisliked, setIsDisliked] = useState(false);
    
    const [newComment, setNewComment] = useState("");
    const [isCommentDisabled, setIsCommentDisabled] = useState(true)
    const [showReportModal, setReportModal] = useState(false);

    const [reportCategory, setReportCategory] = useState("");
    const [reportCategoryErr, setReportCategoryErr] = useState("Select report category");

    const commentInputRef = useRef()

    const store = useSelector(state => state)

    useEffect(() => {
        getUserInfo()
    }, [])

    useEffect(() => {
        const currentTime = moment(new Date())
        const difference = moment.duration(currentTime.diff(post.createdAt))
        if(difference.asHours() < 24){
            difference.asHours() < 0 ? setMinutesAgo(Math.floor(difference.asMinutes())) : setHoursAgo(Math.floor(difference.asHours()))
        }else{
            setDaysAgo(Math.floor(difference.asDays()))
        }
        changeLikesText();
        changeDislikesText()
        if(store.user.jwt !== "") getUserCollections()
    }, [])

    useEffect(()=>{
        changeLikesText()
        changeDislikesText()

        const postLikes = [...post.likes];
        const postDislikes = [...post.dislikes];
        
        setIsLiked(postLikes.filter(like => like.userId === store.user.id).length === 1)
        setIsDisliked(postDislikes.filter(dislike => dislike.userId === store.user.id).length === 1)
    }, [post])

    useEffect(()=>{
        setIsCommentDisabled(newComment.trim() === "" ? true : false)
    }, [newComment])

    useEffect(() => {
        if(store.user.jwt !== "" && store.user.role !== "Admin") getUserCollections()
    }, [showSaveModal])

    const getUserCollections = () => {
        favoritesService.getUserFavoritesOptimized({
            userId: store.user.id,
            jwt: store.user.jwt,
        }).then(response => {
            // Collection for posts that do not have any designated collection
            const newCollections = [...response.data.collections, {
                id: null,
                name: "Unclassified",
                posts: [...response.data.unclassified],
            }];
            console.log(newCollections)
            setCollections(newCollections)

            // Check in which collection the post has been saved
            newCollections.forEach(collection => {
                const found = collection.posts.some(collectionPost => collectionPost.id === post.id)
                if(found) {
                    setIsSaved(true)
                    setSavedInCollections([...savedInCollections, collection.id])
                };
            })
        }).catch(err => {
            toastService.show("error", "Could not load your collections. Please try again.")
        })
    }

    const changeLikesText = () => {
        if(post.likes.length === 0) setLikesText("no one")
        if(post.likes.length === 1) setLikesText("1 person")
        if(post.likes.length >= 2) setLikesText(post.likes.length + " people")
    }

    const changeDislikesText = () => {
        if(post.dislikes.length === 0) setDislikesText("no one")
        if(post.dislikes.length === 1) setDislikesText("1 person")
        if(post.dislikes.length >= 2) setDislikesText(post.dislikes.length + " people")
    }

    const getUserInfo = async () => {
        if(postUser && postUser.id && postUser.id !== store.user.id){
            const response = await userService.getUserById({
                id: postUser.id,
                jwt: store.user.jwt
            })
            
            if(response.status === 200){
                setUser(response.data)
            }else{
                toastService.show("error", "Error retrieving info about user's creator.");
            }
        }else{
            setUser({
                id: store.user.id,
                username: store.user.username,
                profilePhoto: store.user.photo,
            })
        }
    }

    const handleLikeClick = async () => store.user.jwt && await _handleLikeDislikeClick(true)
    const handleDislikeClick = async () => store.user.jwt && await _handleLikeDislikeClick(false)

    const _handleLikeDislikeClick = async (isLike) => {
        const response = await likeService.addLike({
            userId: store.user.id,
            postId: post.id,
            isLike: isLike,
            jwt: store.user.jwt,
        })

        if(response.status === 200){
            const changedPost = {...post};
            const newItem = {
                userId: store.user.id,
                postId: changedPost.id,
                isLike: isLike
            }
            if(isLike) {
                // If like already existed, remove it from the likes list. If it didn't add it there and remove dislike if it exists.
                if(isLiked){
                    changedPost.likes = changedPost.likes.filter(like => like.userId !== store.user.id)
                }else{
                    changedPost.likes.push(newItem)
                    // Remove existing dislike
                    if(isDisliked) changedPost.dislikes = changedPost.dislikes.filter(dislike => dislike.userId !== store.user.id)
                }
            }else{
                if(isDisliked) {
                    changedPost.dislikes = changedPost.dislikes.filter(dislike => dislike.userId !== store.user.id)
                }else{
                    changedPost.dislikes.push(newItem)
                    // Remove existing like
                    if(isLiked) changedPost.likes = changedPost.likes.filter(like => like.userId !== store.user.id)
                }
            }
            setPost(changedPost);
        }else{
            toastService.show("error", "Could not " + (isLike ? "like" : "dislike") + " this post.")
        }
    }

    const handleCommentInputChange = (e) => {
        setNewComment(e.target.value)
    }

    const handleCommentClick = () => {
        commentInputRef.current.focus()
    }

    const postNewComment = async () => {
        if(!store.user.jwt) return
        const comment = {
            userId: store.user.id,
            postId: post.id,
            content: newComment,
            jwt: store.user.jwt
        }
        const response = await commentService.addComment(comment)
        
        if(response.status === 200){
            const changedPost = {...post};
            changedPost.comments.push({
                userId: comment.userId,
                postId: comment.postId,
                content: comment.content,
                username: store.user.username,
                id: uuidv4()
            })
            setPost(changedPost)
            setNewComment("");
        }else{
            toastService.show("error", "Could not comment this post.")
        }
    }

    const handleSaveClick = () => {
        store.user.jwt && setShowSaveModal(!showSaveModal);
    }

    const handleReportModal = () => {
        if(store.user.jwt !== ""){
            setReportCategory("");
            setReportCategoryErr("");
            setReportModal(!showReportModal)
        }
    }

    const handleReportCategoryChange = (event) => {
        if(store.user.jwt !== ""){
            setReportCategory(event.target.value);
            setReportCategoryErr("");
        }
    }

    const sendReport = async () => {
        if(store.user.jwt !== "") return;

        if (reportCategory === "") {
            console.log(reportCategory);
            setReportCategoryErr("Select report category");
            return;
        }
        const response = await complaintService.createComplaint({
            id: "",
            category: reportCategory,
            postId: post.id,
            status: "",
            isPost: true,
            userId: store.user.id,
            jwt: store.user.jwt
        });
        
        if(response.status === 200){
            toastService.show("success", "Post report successfully created.")
        } else{
            toastService.show("error", "Could not create report for this post.")
        }
        handleReportModal();
    }


    return(
        <div className="Post">
            <header>
                <PostHeader 
                    username={user.username} 
                    hideUsername={false}
                    caption={post.location}
                    captionSize="small"
                    image={user.profilePhoto}
                    iconSize="medium"
                    />
                    { store.user.jwt !== "" &&
                        <Dropdown>
                            <Dropdown.Toggle variant="link" id="dropdown-basic">
                                <CardButton className="cardButton" />
                            </Dropdown.Toggle>
                            <Dropdown.Menu>
                                <Dropdown.Item onClick={handleReportModal}>Report</Dropdown.Item>
                            </Dropdown.Menu>
                        </Dropdown>
                    }
            </header>

            <div className="slider">
                <Slider media={post.media} />
            </div>
            <PostMenu 
                isLiked={isLiked}
                isDisliked={isDisliked}
                likeClicked={handleLikeClick}
                dislikeClicked={handleDislikeClick}
                commentClicked={handleCommentClick}
                saveClicked={handleSaveClick}
                isSaved={isSaved}
                postId={post.id}
            />
            <div className="likes-dislikes">
                <PostHeader 
                    hideAccountName={true} 
                    image={user.profilePhoto} 
                    captionSize="large"
                    iconSize="small" />
                <span>Liked by {likesText} and disliked by {dislikesText}. </span>
            </div>
            <div className="Post-caption">
                <strong> <NavLink className="username" to={{pathname: `/profile/${user.username}`}}>{user.username}</NavLink> </strong> {post.description} 
                {post.hashtags.map(hashtag => <span className="hashtag"> #{hashtag.text}</span> )}
            </div>
            <div className="comments">
                {post.comments.length > 0 ? post.comments.map((comment) => {
                    return (
                        <Comment key={comment.id} username={comment.username} comment={comment.content} />
                    );
                }) : <p className="noComments">No comments yet...</p> }
            </div>
            <div className="timePosted">
                { daysAgo < 1 ? (
                    hoursAgo < 1 ? minutesAgo + " minutes ago" : hoursAgo + " hours ago" ) :  
                    daysAgo + " days ago"
                }
            </div>
            <div className="addComment">
                <input 
                    className="commentInput" 
                    placeholder="Add a comment..." 
                    value={newComment} 
                    onChange={handleCommentInputChange}
                    ref={commentInputRef}
                />
                <button className="postText" disabled={isCommentDisabled}  onClick={postNewComment}>Post</button>
            </div>

            <CollectionsModal 
                showModal={store.user.jwt !== "" && showSaveModal} 
                setShowModal={handleSaveClick}
                collections={collections}
                setCollections={setCollections}
                savedInCollections={savedInCollections}
                setSavedInCollections={setSavedInCollections} 
                postId={post.id} 
                isSaved={isSaved}
                setIsSaved={setIsSaved}   
                shouldReload={shouldReload} 
            />

            <Modal show={showReportModal} onHide={setReportModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Report this post?</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <div className="col-sm-6 mb-2">
                        <select onChange={(e) => handleReportCategoryChange(e)} name={"reportCategory"} value={reportCategory}>
                            <option disabled={true} value="">Select report category</option>
                            <option value="Gore">Gore</option>
                            <option value="Nudity">Nudity</option>
                            <option value="Violence">Violence</option>
                            <option value="Suicide">Suicide</option>
                            <option value="Fake News">Fake News</option>
                            <option value="Spam">Spam</option>
                            <option value="Hate Speech">Hate Speech</option>
                            <option value="Terrorism">Terrorism</option>
                            <option value="Harassment">Harassment</option>
                            <option value="Other">Other</option>
                        </select>
                        {reportCategoryErr.length > 0 && <span className="text-danger">{reportCategoryErr}</span>}
                    </div>
                    <p >Are you sure you want to report this post? </p>
                <div style={{display:'flex',float:'right'}}>
                    <Button variant="info" style={{marginRight:'10px'}} onClick={sendReport} >Confirm</Button>
                    <Button variant="secondary" onClick={handleReportModal} >Cancel</Button>
                </div>
                </Modal.Body>
            </Modal>

        </div>
    );
}
export default Post;