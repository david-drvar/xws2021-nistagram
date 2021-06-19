import React, { useState, useEffect } from "react";
import Stories from "react-insta-stories";
import { Modal } from "react-bootstrap";
import Slider from './../Post/Slider';
import renderer from './StoryRenderer';
import empty from './../../images/empty.png'

import storyService from './../../services/story.service'

import '../../style/Story.css';
import '../../style/ProfileIcon.css';

const Highlight = (props) => {
    const { highlight } = props;
    const [showModal, setModal] = useState(false);
    const [convertedStory, setConvertedStory] = useState([])
    const [coverPhoto, setCoverPhoto] = useState("")
    const [header, setHeader] = useState({})

    // Convert story with multiple media to multiple stories with single media, to comply with react-insta-stories
    useEffect(()=>{
        const convertedStories = [];
        highlight.stories.forEach(singleStory => {  
            convertedStories.push(...storyService.convertStory(singleStory))
        })
        setConvertedStory(convertedStories);

        const cover = highlight.stories && highlight.stories[0] && 
                      highlight.stories[0].media && highlight.stories[0].media[0] ? highlight.stories[0].media[0].content : ""
                      console.log(cover)
        setCoverPhoto(cover ? cover : empty)

        setHeader({
            username: highlight.name,
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
                        header={header}
                        width={600} 
                        height={800}/>
                </Modal>
            }
        </div>
    );
}

export default Highlight;