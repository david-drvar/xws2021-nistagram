import React, { useState, useEffect } from 'react';
import { Modal } from 'react-bootstrap';

import Post from './../Post/Post'
import PostPreviewThumbnail from './../Post/PostPreviewThumbnail';
import PostPreviewModal from './../Post/PostPreviewModal';

import "./../../style/PostPreviewGrid.css"

const PostPreviewGrid = (props) => {
    const { shouldReload, isAd } = props;

    const [localPosts, setLocalPosts] = useState([...props.posts])
    const [selectedPost, setSelectedPost] = useState({});
    const [showModal, setShowModal] = useState(false);

    const openPost = (id) => {
        !isAd && setSelectedPost(localPosts.filter(post => post.id === id)[0]);
        isAd && setSelectedPost(localPosts.filter(post => post.post.id === id)[0]);
        setShowModal(true);
    }

    useEffect(() => {
    }, [])

    const updateLocalPostsList = (changedPost) => {
        const tempLocalPosts = [...localPosts];
        for(let i = 0; i < localPosts.length; i++){
            if(localPosts[i].id === changedPost.id){
                tempLocalPosts[i] = {...changedPost}
                break
            }
        }

        setLocalPosts([...tempLocalPosts])
    }

    return (
        (localPosts &&
            <div class="postPreviewGrid__Wrapper">
                {localPosts.map(post => <PostPreviewThumbnail post={isAd ? post.post : post} openPost={openPost} />)}
                {showModal &&
                    <PostPreviewModal 
                    post={selectedPost} 
                    setPosts={updateLocalPostsList}
                    isAd={isAd}
                    shouldReload={shouldReload}
                    showModal={showModal}
                    setShowModal={setShowModal}
                    />
                }
            </div>
        )
    )
}

export default PostPreviewGrid;