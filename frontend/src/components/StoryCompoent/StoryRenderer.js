import React, { useState, useEffect } from 'react';
import Slider from './../Post/Slider';
import { WithHeader, WithSeeMore } from 'react-insta-stories';
import adsService from './../../services/ads.service'
import moment from 'moment';


const StoryRenderer = ({ story, action, isPaused, config }) => {
    const { header } = config;

    const [localStory, setLocalStory] = useState(story);
    const [caption, setCaption] = useState("");

    const seeMoreFactory = (url) => {
        if(!url) return
        if (!url.startsWith('http')) {
            url = "http://" + url
        }
        return ({ close }) => {
            close()
            window.open(url, '_blank')
            adsService.incrementLinkClicks({jwt: "", adId: story.id})
            return null
        }
    }

    useEffect(() => {
        let timeCreated = "";

        header.setStoryId && header.setStoryId(story.id)

        if (story.createdAt) {
            const currentTime = moment(new Date())
            const difference = moment.duration(currentTime.diff(story.createdAt))

            if (difference.asHours() < 1) {
                difference.asMinutes() < 1 ?
                    timeCreated = Math.floor(difference.asSeconds()) + "s ago @ " + story.location :
                    timeCreated = Math.floor(difference.asMinutes()) + "m ago @ " + story.location
            } else {
                difference.asHours() > 24 ?
                    timeCreated = Math.floor(difference.asDays()) + "d ago @ " + story.location :
                    timeCreated = Math.floor(difference.asHours()) + "h ago @ " + story.location
            }
        }

        const sponsored = story.isAd ? " · Sponsored" : ""

        const newHeader = {
            heading: header.heading ? header.heading + sponsored : "username",
        }
        if (timeCreated) newHeader["subheading"] = timeCreated;
        if (header.profileImage) newHeader["profileImage"] = header.profileImage;
        if (story.isCloseFriends) newHeader.heading += " - Close Friends"

        let tempLocalStory = {
            ...story,
            header: newHeader
        }
        if (story.isAd && (header.link !== "" || story.link)) tempLocalStory["seeMore"] = seeMoreFactory(story.link)
        setLocalStory(tempLocalStory)

        let hashtags = "";
        story.hashtags && story.hashtags.forEach(hashtag => hashtags += ` #${hashtag.text}`)
        const caption = story.description + hashtags
        setCaption(caption)

        action('play'); // Doesn't auto start if there are multiple stories
    }, [story])

    const renderStory = () => {
        if (header.link !== "" || story.link !== "") {
            return (
                <WithSeeMore story={localStory} action={action}>
                    <WithHeader story={localStory}>
                        <Slider
                            showStoryCaption={true}
                            storyCaption={caption}
                            showTags={true}
                            media={story.media}
                        />
                    </WithHeader>
                </WithSeeMore>
            )
        } else {
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
    }

    return (
        renderStory()
    )
}

const Tester = (story) => {
    return { condition: true, priority: 10 }
}

export default {
    renderer: StoryRenderer,
    tester: Tester,
}