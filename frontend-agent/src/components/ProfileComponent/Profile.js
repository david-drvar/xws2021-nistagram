import React, { useEffect, useState } from "react";
// import { useDispatch, useSelector } from "react-redux";
// import { Button, Modal } from "react-bootstrap";
// import { useParams } from 'react-router-dom'
//
// import FollowAndUnfollow from "./FollowAndUnfollow";
// import Navigation from "../HomePage/Navigation";
// import { userActions } from "../../store/actions/user.actions";
// import FollowersAndFollowings from "./FollowersAndFollowings";
// import BlockMuteAndNotifications from "./BlockMuteAndNotifications";
// import Highlight from './../StoryCompoent/Highlight';
// import PostPreviewGrid from './../Post/PostPreviewGrid';
import Spinner from './../../helpers/spinner';
// import Story from './../StoryCompoent/Story';
//
import agentService from "../../services/agent.service";
import productService from "../../services/product.service";
// import privacyService from "../../services/privacy.service";
// import followersService from "../../services/followers.service";
// import postService from './../../services/post.service';
// import storyService from './../../services/story.service';
// import highlightsService from './../../services/highlights.service';
import toastService from './../../services/toast.service';

import '../../style/Profile.css';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";
import { useParams } from 'react-router-dom'
import ProductPreviewGrid from "../Product/ProductPreviewGrid";


const Profile = () => {
    const {username} = useParams()
    const [loadingPosts, setLoadingPosts] = useState(true);
    const [user, setUser] = useState({});
    const [posts, setPosts] = useState([]);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        (async function () {
            await getUserByUsername(); // Since it doesn't get saved in time for other requests
            await getProducts();

        })();
    }, [username, user]);

    const getProducts = async () => {
        const response = await productService.getProductsByAgent({
            id: user.id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setPosts([...response.data.products])
            setLoadingPosts(false);
        } else {
            console.log(response);
            toastService.show("error", "Could not retrieve user's posts.")
        }
    }

    async function getUserByUsername() {
        const response = await agentService.getUserByUsername({
            username: username,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUser(response.data)
            console.log(response.data)
            return response.data
        } else {
            console.log("getuserbyusername error")
        }
    }


    return (
        <div>
            <Navigation/>
            <div className="profileGrid">
                <div className="profileHeader">
                    <img style={{marginLeft: "-1em", paddingRight: "4px"}} alt="" src={user.profilePhoto}/>
                    <div className="info">
                        <div className="fullname">
                            {user.firstName} {user.lastName}
                        </div>
                        <div className="username">@{user.username}</div>

                    </div>
                </div>

                <div className="content">
                    <div className="posts">
                        {loadingPosts ?
                            <div style={{position: "relative", left: "45%", marginTop: "50px"}}>
                                <Spinner type="MutatingDots" height="100" width="100"/>
                            </div> :
                            <ProductPreviewGrid posts={posts}/>
                        }
                    </div>
                </div>
            </div>
        </div>
    );
}
export default Profile;
