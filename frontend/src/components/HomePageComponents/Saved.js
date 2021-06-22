import React, { useState, useEffect } from 'react';
import { useSelector } from "react-redux"
import { ListGroup, Button } from 'react-bootstrap';

import PostPreviewGrid from './../Post/PostPreviewGrid'
import Navigation from "../HomePage/Navigation";
import Spinner from './../../helpers/spinner';

import favoritesService from './../../services/favorites.service' 
import collectionsService from './../../services/collections.service' 
import toastService from './../../services/toast.service' 

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
        favoritesService.getUserFavorites({
            userId: store.user.id,
            jwt: store.user.jwt
        }).then(res => {
            const allPosts = [];
            res.data.collections.forEach(collection => allPosts.push(...collection.posts) )
            res.data.unclassified.forEach(post => allPosts.push(post) )
            setPosts(allPosts)
            setPostsToRender(allPosts)
            const allCollections = [...res.data.collections]
            allCollections.push({
                id: null,
                name: "Unclassified",
                posts: [...res.data.unclassified]
            })
            setCollections(allCollections)
            setPostsLoading(false);
        }).catch(err => {
            console.log(err)
            toastService.show("error", "Could not retrieve your saved posts.")
        })
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
        const collectionRequest = {
            name: newCollectionName,
            userId: store.user.id,
            jwt: store.user.jwt
        }
        const response = await collectionsService.createCollection(collectionRequest)

        if(response.status === 200) {
            setNewCollectionName("");
            setIsNewCollectionButtonDisabled(true)
            setCollections([...collections, {
                id: response.data.id,
                name: collectionRequest.name,
                userId: collectionRequest.userId,
                posts: [],
            }])
            toastService.show("success", "Successfully created new collection called " + collectionRequest.name + ".");
        }
    }

    return (
        <div>
            <Navigation/>
            <div className="Saved__Wrapper">
                <div className="posts">
                    <div className="title">Saved Posts</div>
                    <div className="content">
                        { postsLoading ? <Spinner type="MutatingDots" /> : 
                            (postsToRender.length > 0 ? <PostPreviewGrid shouldReload={true} posts={postsToRender} /> : <div>No saved posts in here!</div>) }
                    </div>
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