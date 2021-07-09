import React, { useState, useEffect } from "react";
import "../../style/Posts.css";
import Post from "./Post";
import { useSelector } from 'react-redux';
import postService from './../../services/post.service'
import toastService from './../../services/toast.service'
import Spinner from './../../helpers/spinner';

const Posts = (props) => {
    const [loading, setLoading] = useState(true);
    const [posts, setPosts] = useState([]);
    const store = useSelector(state => state)

    // TODO Retrieve ads as well
    useEffect(()=>{
        postService.getHomepagePosts({ jwt: store.user.jwt })
            .then(response => {
                if(response.status === 200) {
                    console.log(response.data)
                    setPosts([...response.data.posts, ...response.data.ads])
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
                posts && posts.map((post) => <Post post={post} isAd={post.link ? true : false} /> )
            }
        </div>

    );
}

export default Posts;