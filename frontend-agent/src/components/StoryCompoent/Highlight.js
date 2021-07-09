import React, { useState, useEffect } from "react";
import Stories from "react-insta-stories";
import { Modal } from "react-bootstrap";
import Slider from './../Post/Slider';
import renderer from './StoryRenderer';
import empty from './../../images/empty.png'

import storyService from './../../services/nistagram api/story.service'

import '../../style/Story.css';
import '../../style/ProfileIcon.css';

const Highlight = (props) => {
    const { highlight } = props;
    const [showModal, setModal] = useState(false);
    const [convertedStory, setConvertedStory] = useState([])
    const [coverPhoto, setCoverPhoto] = useState("")
    const [header, setHeader] = useState({})
    const [storyId, setStoryId] = useState(""); // temp fix for reports

    // Convert story with multiple media to multiple stories with single media, to comply with react-insta-stories
    useEffect(()=>{
        const convertedStories = [];
        highlight.stories.forEach(singleStory => {  
            convertedStories.push(...storyService.convertStory(singleStory))
        })
        setConvertedStory(convertedStories);
                      
        // Set the first photo to be a cover, otherwise set a default empty highlight photo
        setCoverPhoto(highlight.stories && highlight.stories[0] &&  highlight.stories[0].media && highlight.stories[0].media[0] ?
                      highlight.stories[0].media[0].content : empty)

        setHeader({
            heading: highlight.name
        })
    }, [])

    return (
        <div>
            <div className="story">
                <div className={true ? "storyBorder" : ""}>
                    <img className={`profileIcon big`} src={coverPhoto} alt="" onClick={() => setModal(!showModal)}/>
                </div>
                <span className="accountName">{highlight.name}</span>
            </div>
            
            { highlight.stories.length > 0 &&
                <Modal show={showModal} onHide={() => setModal(!showModal)}>
                    <Stories
                        onAllStoriesEnd={() => setModal(!showModal)}
                        renderers={[renderer]}
                        stories={convertedStory} 
                        defaultInterval={10000} 
                        header={{...header, setStoryId}}
                        width={500}
                        height={700}/>
                </Modal>
            }
        </div>
    );
}

export default Highlight;