import React, { useState, useEffect } from 'react';
import PostPreviewThumbnail from './../Post/PostPreviewThumbnail';
import PostPreviewModal from './../Post/PostPreviewModal';

import "./../../style/PostPreviewGrid.css"

const PostPreviewGrid = (props) => {
    const { shouldReload } = props;

    const [localPosts, setLocalPosts] = useState([])
    const [selectedPost, setSelectedPost] = useState({});
    const [showModal, setShowModal] = useState(false);

    useEffect(() => {
        setLocalPosts([...props.posts])
    }, [props.posts])

    const openPost = (selectPost) => {
        selectPost.link === "" && setSelectedPost(localPosts.filter(post => post.id === selectPost.id)[0]);
        selectPost.link !== "" && setSelectedPost(localPosts.filter(post => post.post.id === selectPost.id)[0]);
        setShowModal(true);
    }

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
                {localPosts.map(post => <PostPreviewThumbnail post={post.link ? post.post : post} openPost={openPost} />)}
                {showModal &&
                    <PostPreviewModal 
                    post={selectedPost} 
                    setPosts={updateLocalPostsList}
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