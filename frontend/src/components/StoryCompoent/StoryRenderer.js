import React, { useState, useEffect } from 'react';
import Slider from './../Post/Slider';
import { WithHeader } from 'react-insta-stories';
import moment from 'moment';


const StoryRenderer = ({ story, action, isPaused, config }) => {
    const { header } = config;

    const [localStory, setLocalStory] = useState(story);
    const [caption, setCaption] = useState("");

    useEffect(()=>{
        let timeCreated = "";

        header.setStoryId && header.setStoryId(story.id)

        if(story.createdAt){
            const currentTime = moment(new Date())
            const difference = moment.duration(currentTime.diff(story.createdAt))

            if(difference.asHours() < 1){
                difference.asMinutes() < 1 ? 
                    timeCreated = Math.floor(difference.asSeconds()) + "s ago @ " + story.location : 
                    timeCreated = Math.floor(difference.asMinutes()) + "m ago @ " + story.location
            }else{
                difference.asHours() > 24 ? 
                    timeCreated = Math.floor(difference.asDays()) + "d ago @ " + story.location :
                    timeCreated = Math.floor(difference.asHours()) + "h ago @ " + story.location
            }
        }

        const newHeader = {
            heading: header.heading ? header.heading : "username",
        }
        if(timeCreated) newHeader["subheading"] = timeCreated;
        if(header.profileImage) newHeader["profileImage"] = header.profileImage;
        if(story.isCloseFriends) newHeader.heading += " - Close Friends"

        setLocalStory({
            ...story,
            header: newHeader
        })

        let hashtags = "";
        story.hashtags && story.hashtags.forEach(hashtag => hashtags += ` #${hashtag.text}`)
        const caption = story.description + hashtags
        setCaption(caption)

        action('play'); // Doesn't auto start if there are multiple stories
    }, [story])

    return (
        <WithHeader story={localStory}>
            <Slider 
                showStoryCaption={true}
                storyCaption={caption}
                showTags={true} 
                media={story.media}
                />
        </WithHeader>
    )
}

const Tester = (story) => {
    return { condition: true, priority: 10 }
}

export default {
    renderer: StoryRenderer,
    tester: Tester,
}