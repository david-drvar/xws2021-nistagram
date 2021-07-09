import React, { useEffect, useState } from 'react';
import { Button, Modal, ListGroup } from "react-bootstrap";
import { useSelector } from 'react-redux';
import { ReactComponent as Check } from './../../images/icons/check.svg'
import collectionsService from './../../services/nistagram api/collections.service'
import favoritesService from './../../services/nistagram api/favorites.service'
import toastService from './../../services/toast.service'
import './../../style/CollectionsModal.css'

const CollectionsModal = (props) => {
    const { 
        showModal, 
        setShowModal, 
        postId,
        collections, 
        setCollections, 
        savedInCollections, 
        setSavedInCollections,
        shouldReload // shouldReload is used for refreshing page when a post is saved within collections menu
    } = props;

    const [newCollectionName, setNewCollectionName] = useState("");
    const [isNewCollectionButtonDisabled, setIsNewCollectionButtonDisabled] = useState(true);
    const [selectedCollection, setSelectedCollection] = useState({});   // Collection to save a post in 
    const [isButtonRemove, setIsButtonRemove] = useState(false)  // Checking if button should be save or remove

    const store = useSelector(state => state);

    useEffect(()=>{
        setIsNewCollectionButtonDisabled(newCollectionName === "")
    }, [newCollectionName])

    useEffect(() => {
        setIsButtonRemove(savedInCollections.some(id => selectedCollection.id === id))
    }, [selectedCollection])

    const saveToCollection = async () => {
        const newFavorite = {
            userId: store.apiKey.id,
            postId: postId,
            collectionId: selectedCollection.id,
            jwt: store.apiKey.jwt,
        }
        
        const response = await favoritesService.createFavorite(newFavorite);
        
        const collectionName = selectedCollection.name ? selectedCollection.name : "unclassified posts" 
        if(response.status === 200){
            toastService.show("success", "Successfully saved this post to " + collectionName + ".");
            setSavedInCollections([...savedInCollections, selectedCollection.id])
            setSelectedCollection({});
            setShowModal(!showModal);
            shouldReload && window.location.reload()
        }else{
            toastService.show("error", "Could not save this post to " + collectionName + ".");
        }
    }

    const removeFromCollection = async () => {
        const removingFavorite = {
            userId: store.apiKey.id,
            postId: postId,
            collectionId: selectedCollection.id,
            jwt: store.apiKey.jwt,
        }
        
        const response = await favoritesService.removeFavorite(removingFavorite);
        
        const collectionName = selectedCollection.name ? selectedCollection.name : "unclassified posts" 
        if(response.status === 200){
            toastService.show("success", "Successfully removed this post from " + collectionName + ".");
            setSavedInCollections([...savedInCollections.filter(collection => collection.id !== selectedCollection.id)])
            setSelectedCollection({});
            setShowModal(!showModal);
            shouldReload && window.location.reload()
        }else{
            toastService.show("error", "Could not remove this post from " + collectionName + ".");
        }
    }

    const createNewCollection = async () => {
        const collectionRequest = {
            name: newCollectionName,
            userId: store.apiKey.id,
            jwt: store.apiKey.jwt
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

    const handleNewCollectionName = (e) => {
        setNewCollectionName(e.target.value)
    }

    return (
        <Modal className="saveModal" size="lg" show={showModal} onHide={setShowModal} animation={true} centered>
            <Modal.Header closeButton>
                <Modal.Title>Choose a collection</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <ListGroup variant="flush" className="collectionsList">
                    {collections.map((collection, key) => {
                        return (
                            <ListGroup.Item 
                                action    
                                active={selectedCollection.id === collection.id} 
                                onClick={() => setSelectedCollection(collection)}
                                className="collectionCard">
                                    { savedInCollections.some(id => collection.id === id) ? 
                                        <div className="savedInCollection">{collection.name} <Check className="savedCheck" /> </div> : 
                                        collection.name 
                                    } 
                            </ListGroup.Item>
                        )
                    })}
                </ListGroup>
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
            </Modal.Body>
            <Modal.Footer style={{'background':'#E0E0E0'}}>
                { !isButtonRemove ?
                    <Button variant="primary" onClick={saveToCollection}>Save</Button> : 
                    <Button variant="primary" onClick={removeFromCollection}>Remove</Button>
                }
            </Modal.Footer>
        </Modal>
    )
}

export default CollectionsModal;