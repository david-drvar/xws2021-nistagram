import React, {useEffect, useState} from 'react';
import ProfileInfo from "./ProfileInfo";
import {useDispatch, useSelector} from "react-redux";
import likeService from './../../services/like.service'
import postService from "../../services/post.service";
import toastService from "../../services/toast.service";
import Post from "../Post/Post";
import PostPreviewGrid from "../Post/PostPreviewGrid";
import Spinner from "../../helpers/spinner";


function Liked() {
    const dispatch = useDispatch()
    const store = useSelector(state => state);
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        getUserLikedOrDislikedPosts();
    }, []);

    function getUserLikedOrDislikedPosts() {
        likeService.getUserLikedOrDislikedPosts({ jwt: store.user.jwt , userId : store.user.id, isLike : true})
            .then(response => {
                if(response.status === 200){
                    setPosts(response.data.posts)
                    setLoading(false);

                }
            })
            .catch(err => {
                toastService.show("error", "Could not retrieve liked posts.")
            })
    }

    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            {loading ?
                <div style={{marginRight: '30%', marginTop: '5%', display: 'flex', flexDirection: 'column'}}>
                <Spinner type="MutatingDots" height="100" width="100"/>
                    </div>:
                <div style={{marginRight: '30%', marginTop: '5%', display: 'flex', flexDirection: 'column'}}>
                    <PostPreviewGrid posts={posts}/>

                </div>
            }
        </div>
    );
}

export default Liked;