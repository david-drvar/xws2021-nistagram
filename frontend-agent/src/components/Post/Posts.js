import React, { useState, useEffect } from "react";
import "../../style/Posts.css";
import Post from "./Post";
import { useSelector } from 'react-redux';
import postService from './../../services/nistagram api/post.service'
import toastService from './../../services/toast.service'
import Spinner from './../../helpers/spinner';

const Posts = (props) => {
    const [loading, setLoading] = useState(true);
    const [posts, setPosts] = useState([]);
    const store = useSelector(state => state)

    // TODO Retrieve ads as well
    useEffect(()=>{
        postService.getHomepagePosts({ jwt: store.apiKey.jwt })
            .then(response => {
                if(response.status === 200) {
                    setPosts(response.data.posts)
                    setLoading(false);
                }
            })
            .catch(err => {
                toastService.show("error", "Could not retrieve homepage posts.")
            })
    }, [])

    return(
        <div className={`Posts ${loading ? "loading" : ""}`}>
            { loading ? 
                <Spinner type="MutatingDots" height="100" width="100" /> : 
                posts && posts.map((post) => <Post post={post} postUser={{ id: post.userId }}/> )
            }
        </div>

    );
}

export default Posts;