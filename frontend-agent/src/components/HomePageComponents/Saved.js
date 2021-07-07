import React, { useState, useEffect } from 'react';
import { useSelector } from "react-redux"
import { ListGroup, Button } from 'react-bootstrap';
import Navigation from "../HomePage/Navigation";


import './../../style/Saved.css'

const Saved = () => {
    const [postsLoading, setPostsLoading] = useState(true);

    const [posts, setPosts] = useState([]);
    const [postsToRender, setPostsToRender] = useState([]);
    const [collections, setCollections] = useState([]);
    const [selectedCollection, setSelectedCollection] = useState({id: "1"});
    const [newCollectionName, setNewCollectionName] = useState("");
    const [isNewCollectionButtonDisabled, setIsNewCollectionButtonDisabled] = useState(true);

    const store = useSelector(state => state);

    useEffect(() => {
    }, []);

    useEffect(() => {
        if(selectedCollection.id === "1") {
            setPostsToRender(posts);
        }else{
            const newPosts = selectedCollection.posts.filter(collectionPost => {
                return posts.some(post => post.id === collectionPost.id)
            })
            setPostsToRender(newPosts)
        }
    }, [selectedCollection])

    const renderCollections = () => {
        return (
            <ListGroup variant="flush" className="collectionsList">
                <ListGroup.Item 
                    action    
                    active={selectedCollection.id === "1"} 
                    onClick={() => setSelectedCollection({id: "1"})}
                    className="collectionCard">
                        All posts
                </ListGroup.Item>
                {collections && collections.map((collection, key) => {
                    return (
                        <ListGroup.Item 
                            action    
                            active={selectedCollection.id === collection.id} 
                            onClick={() => setSelectedCollection(collection)}
                            className="collectionCard">
                                { collection.name } 
                        </ListGroup.Item>
                    )
                })}
            </ListGroup>
        )
    }

    useEffect(()=>{
        newCollectionName && setIsNewCollectionButtonDisabled(false)
    }, [newCollectionName])

    const handleNewCollectionName = (e) => {
        setNewCollectionName(e.target.value)
    }

    const createNewCollection = async () => {

    }

    return (
        <div>
            <Navigation/>
            <div className="Saved__Wrapper">
                <div className="posts">
                    <div className="title">Saved Posts</div>
                </div>
                <div className="collections">
                    <div className="list">
                        <div className="title">Your Collections</div>
                        { renderCollections() }
                        <div className="newCollection">
                            <p className="newCollectionLabel">New collection: </p>
                            <input 
                                className="newCollectionInput" 
                                placeholder="Enter collection name..." 
                                value={newCollectionName} 
                                onChange={handleNewCollectionName}
                            />
                            
                            <Button 
                                className="newCollectionButton" 
                                variant="outline-primary" 
                                disabled={isNewCollectionButtonDisabled} 
                                onClick={createNewCollection}>Create</Button>
                        </div>
                    </div> 
                </div>
            </div>
        </div>
    );
}

export default Saved;