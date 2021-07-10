import React from 'react';
import { ReactComponent as FilledComments } from './../../images/icons/comments-filled.svg'
import { ReactComponent as FilledHeart } from './../../images/icons/heart-filled.svg'
import { ReactComponent as FilledBrokenHeart } from './../../images/icons/broken-heart-filled.svg'

import "./../../style/PostPreviewThumbnail.css"

const PostPreviewThumbnail = (props) => {
    const { post, openPost } = props;

    return (
        <div className="postPreviewThumbnail__Wrapper" onClick={() => openPost(post)}>
            <img src={post.media[0].content} className="thumbnail" alt="" />
            <div className="hover-overlay hover-overlay--blur">
                <div className="single-statistic">
                    <span><FilledHeart width={25} fill="white" /></span><span>{post.likes.length}</span>
                </div>
                <div className="single-statistic">
                    <span><FilledBrokenHeart width={25} fill="white"/></span><span>{post.dislikes.length}</span>
                </div>
                <div className="single-statistic">
                    <span><FilledComments fill="white"/></span><span>{post.comments.length}</span>
                </div>
            </div>
        </div>
    )
}

export default PostPreviewThumbnail;