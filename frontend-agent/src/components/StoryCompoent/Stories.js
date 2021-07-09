import React, { useState, useEffect } from 'react';
import ScrollMenu from "react-horizontal-scrolling-menu";
import Story from "./Story";
import '../../style/Stories.css';
import { useSelector } from 'react-redux';
import storyService from './../../services/story.service'
import toastService from './../../services/toast.service'
import Spinner from './../../helpers/spinner'; 

const Stories = () => {
    const [loading, setLoading] = useState(true);
    const [stories, setStories] = useState([]);
    const store = useSelector(state => state);

    // TODO Retrieve ads as well
    useEffect(()=>{
        storyService.getHomepageStories({ jwt: store.apiKey.jwt })
            .then(response => {
                console.log(response)
                if(response.status === 200) {
                    setStories(convertToComponentArray([...response.data.stories]));
                    setLoading(false);
                }
            })
            .catch(err => {
                console.log(err)
                toastService.show("error", "Could not retrieve homepage stories.")
            })
    }, [])
    
    const convertToComponentArray = (stories) => stories.map((story) => <Story fixMargins={true} story={story} /> )

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