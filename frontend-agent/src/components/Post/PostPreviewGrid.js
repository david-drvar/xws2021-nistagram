import React, { useState, useEffect } from 'react';
import { Modal } from 'react-bootstrap';

import PostPreviewThumbnail from './../Post/PostPreviewThumbnail';
import PostPreviewModal from './../Post/PostPreviewModal';

import "./../../style/PostPreviewGrid.css"

const PostPreviewGrid = (props) => {
    const { posts, shouldReload, isAd } = props;

    const [selectedPost, setSelectedPost] = useState({});
    const [showModal, setShowModal] = useState(false);

    const openPost = (id) => {
        !isAd && setSelectedPost(posts.filter(post => post.id === id)[0]);
        isAd && setSelectedPost(posts.filter(post => post.post.id === id)[0]);
        setShowModal(true);
    }

    useEffect(() => {
    }, [])

    return (
        ( posts && 
        <div class="postPreviewGrid__Wrapper">
            { posts.map(post => <PostPreviewThumbnail post={isAd ? post.post : post} openPost={openPost} /> ) }    
            { showModal && 
            <PostPreviewModal 
                post={selectedPost} 
                isAd={isAd}
                shouldReload={shouldReload}
                showModal={showModal}
                setShowModal={setShowModal}
            /> }
        </div>
        )
    )
}

export default PostPreviewGrid;