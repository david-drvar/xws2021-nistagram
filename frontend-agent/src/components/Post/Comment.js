import React from 'react';
import './../../style/comment.css';

const Comment = (props) => {
    const { username, comment } = props;

    return (
        <div className="commentContainer">
            <div className="username">{username}</div>
            <div className="comment">{comment}</div>
        </div>
    )
}

export default Comment;