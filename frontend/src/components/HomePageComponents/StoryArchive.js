import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import { ListGroup, Button } from 'react-bootstrap';
import Stories from "react-insta-stories";
import renderer from './../StoryCompoent/StoryRenderer';
import HighlightModal from './HighlightModal';

import Navigation from "./../HomePage/Navigation";
import Spinner from './../../helpers/spinner';

import storyService from './../../services/story.service' 
import highlightsService from './../../services/highlights.service' 
import toastService from './../../services/toast.service' 

import './../../style/StoryArchive.css'

const StoryArchive = (props) => {
    const [storiesLoading, setStoriesLoading] = useState(true);
    const [highlightsLoading, setHighlightsLoading] = useState(true);
    const [showModal, setShowModal] = useState(false);

    const [selectedStory, setSelectedStory] = useState({});
    const [stories, setStories] = useState([]);
    const [storiesToRender, setStoriesToRender] = useState([]);
    const [highlights, setHighlights] = useState([]);
    const [selectedHighlight, setSelectedHighlight] = useState({id: "1"});
    const [newHighlightName, setNewHighlightName] = useState("");
    const [isNewHighlightButtonDisabled, setIsNewHighlightButtonDisabled] = useState(true);

    const store = useSelector(state => state);

    useEffect(() => {
        storyService.getMyStories({
            userId: store.user.id,
            jwt: store.user.jwt
        }).then(res => {
            const tempStories = []
            res.data.stories.forEach(story => {
                tempStories.push(...storyService.convertStory(story))
            })
            setStories(tempStories)
            setStoriesToRender(renderStories(tempStories))
            setStoriesLoading(false);
        }).catch(err => {
            console.log(err)
            toastService.show("error", "Could not retrieve your stories.")
        })

        highlightsService.getUserHighlights({
            userId: store.user.id,
            jwt: store.user.jwt
        }).then(res => {
            setHighlights([...res.data.highlights]);
            setHighlightsLoading(false);
        }).catch(err => {
            console.log(err)
            toastService.show("error", "Could not retrieve your highlights.")
        })
    }, [])

    useEffect(() => {
        if(selectedHighlight.id === "1") {
            setStoriesToRender(renderStories(stories));
        }else{
            const newStories = selectedHighlight.stories.filter(highlightStory => {
                return stories.some(story => story.id === highlightStory.id)
            })
            setStoriesToRender(renderStories(newStories))
        }
    }, [selectedHighlight])
    
    const openHighlightModal = (story) => {
        setSelectedStory(story);
        setShowModal(true);
    }

    const renderStories = (stories) => {
        const header = {
            heading: store.user.username,
            profileImage: store.user.photo
        };

        return stories.map(story => (
            <div className="archivedStory" onClick={() => openHighlightModal(story)}>
                <Stories
                    loop={true}
                    renderers={[renderer]}
                    stories={[story]} 
                    defaultInterval={10000} 
                    header={header}
                    width={400} 
                    height={550}/>
            </div>
        ))
    }

    const renderHighlights = () => {
        return (
            <ListGroup variant="flush" className="collectionsList">
                <ListGroup.Item 
                    action    
                    active={selectedHighlight.id === "1"} 
                    onClick={() => setSelectedHighlight({id: "1"})}
                    className="highlightCard">
                        All stories
                </ListGroup.Item>
                {highlights && highlights.map((highlight, key) => {
                    return (
                        <ListGroup.Item 
                            action    
                            active={selectedHighlight.id === highlight.id} 
                            onClick={() => setSelectedHighlight(highlight)}
                            className="highlightCard">
                                { highlight.name } 
                        </ListGroup.Item>
                    )
                })}
            </ListGroup>
        )
    }

    useEffect(()=>{
        newHighlightName && setIsNewHighlightButtonDisabled(false)
    }, [newHighlightName])

    const handleNewHighlightName = (e) => {
        setNewHighlightName(e.target.value)
    }

    const createNewHighlight = async () => {
        const highlightRequest = {
            name: newHighlightName,
            userId: store.user.id,
            jwt: store.user.jwt
        }
        const response = await highlightsService.createHighlight(highlightRequest)

        if(response.status === 200) {
            setNewHighlightName("");
            setIsNewHighlightButtonDisabled(true);
            setHighlights([...highlights, {
                id: response.data.id,
                name: highlightRequest.name,
                userId: highlightRequest.userId,
                stories: [],
            }])
            toastService.show("success", `You have created a new highlight called ${highlightRequest.name}!`)
        }else{
            toastService.show("error", `Could not create a new highlight`)
        }
    }

    return (
        <div>
            <Navigation/>
            <div className="StoryArchive__Wrapper">
                <div className="stories">
                    <div className="title">Stories Archive</div>
                    <div className="content">
                        { storiesLoading ? <Spinner type="MutatingDots" /> : 
                            (storiesToRender.length > 0 ? storiesToRender : <div>No stories in this highlight!</div>) }
                    </div>
                    <HighlightModal 
                        story={selectedStory}
                        highlights={highlights}
                        setHighlights={setHighlights}
                        showModal={showModal}
                        setShowModal={setShowModal}
                        setStoriesToRender={setStoriesToRender}
                        storiesToRender={storiesToRender}
                        removeFromRender={selectedHighlight.id !== "1"}
                    /> 
                </div>
                <div className="highlights">
                    { highlightsLoading ? <Spinner /> : 
                        <div className="list">
                            <div className="title">Your Highlights</div>
                            { renderHighlights() }
                            <div className="newHighlight">
                                <p className="newHighlightLabel">New highlight: </p>
                                <input 
                                    className="newHighlightInput" 
                                    placeholder="Enter highlight name..." 
                                    value={newHighlightName} 
                                    onChange={handleNewHighlightName}
                                />
                                
                                <Button 
                                    className="newHighlightButton" 
                                    variant="outline-primary" 
                                    disabled={isNewHighlightButtonDisabled} 
                                    onClick={createNewHighlight}>Create</Button>
                            </div>
                        </div> 
                    }
                </div>
            </div>
        </div>
    )
}

export default StoryArchive;