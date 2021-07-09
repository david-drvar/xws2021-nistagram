import React, { useState, useEffect } from 'react';
import ScrollMenu from "react-horizontal-scrolling-menu";
import Story from "./Story";
import '../../style/Stories.css';
import { useSelector } from 'react-redux';
import storyService from './../../services/story.service'
import userService from './../../services/user.service'
import toastService from './../../services/toast.service'
import Spinner from './../../helpers/spinner'; 
import { asyncForEach } from './../../util'

const Stories = () => {
    const [loading, setLoading] = useState(true);
    const [stories, setStories] = useState([]);
    const store = useSelector(state => state);

    useEffect(()=>{
        storyService.getHomepageStories({ jwt: store.user.jwt })
            .then(async response => {
                let allStoriesTemp = [...response.data.stories]
                console.log(response)

                await asyncForEach(response.data.ads, async singleAd => {
                    if(singleAd.ownerHasStories){
                        // Add them to owners ads
                        allStoriesTemp = [...allStoriesTemp.map(story => {
                            if(story.userId === singleAd.ad.post.userId){
                                return {
                                    ...story,
                                    stories: [...story.stories, singleAd.ad]
                                }
                            }else{
                                return story
                            }
                        })]
                    }else{
                        // Retrieve owner data and make standalone story
                        const response = await userService.getUserById({
                            id: singleAd.ad.post.userId,
                            jwt: store.user.jwt
                        })
                        
                        if(response.status === 200){
                            allStoriesTemp.push({
                                userId: singleAd.ad.post.userId,
                                username: response.data.username,
                                userPhoto: response.data.profilePhoto,
                                stories: [{...singleAd.ad}]
                            })
                        }
                    }
                })
                
                setStories(convertToComponentArray([...allStoriesTemp]));
                setLoading(false);
            })
            .catch(err => {
                console.log(err)
                toastService.show("error", "Could not retrieve homepage stories.")
            })
    }, [])
    
    const convertToComponentArray = (stories) => stories.map((story) => <Story fixMargins={true} story={story.link ? story.post : story} link={story.link ? story.link : ""} /> )

    return (
        <div className={`stories ${loading ? "loading" : ""}`}>
            { loading ? <Spinner /> :
            <ScrollMenu 
                alignCenter={false}
                data={stories}
                wheel={true}
                hideArrows={true}
                hideSingleArrow={true}
                />
            }
        </div>
    );
}

export default Stories;