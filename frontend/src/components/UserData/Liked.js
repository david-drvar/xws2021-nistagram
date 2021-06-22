import React, {useEffect, useState} from 'react';
import ProfileInfo from "./ProfileInfo";
import {useDispatch, useSelector} from "react-redux";
import likeService from './../../services/like.service'
import postService from "../../services/post.service";
import toastService from "../../services/toast.service";
import Post from "../Post/Post";
import PostPreviewGrid from "../Post/PostPreviewGrid";


function Liked() {
    const dispatch = useDispatch()
    const store = useSelector(state => state);
    const [posts, setPosts] = useState([]);

    useEffect(() => {
        getUserLikedOrDislikedPosts();
        console.log(store)
    }, []);

    function getUserLikedOrDislikedPosts() {
        likeService.getUserLikedOrDislikedPosts({ jwt: store.user.jwt , userId : store.user.id, isLike : true})
            .then(response => {
                if(response.status === 200) setPosts(response.data.posts)
            })
            .catch(err => {
                toastService.show("error", "Could not retrieve liked posts.")
            })
    }

    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '30%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                <PostPreviewGrid posts={posts} />

            </div>
        </div>
    );
}

export default Liked;