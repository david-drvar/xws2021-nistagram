import moment from 'moment';
import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux'
import { Modal, Button, DropdownButton, Dropdown } from 'react-bootstrap'

import adsService from '../../services/ads.service';
import toastService from '../../services/toast.service';

import './../../style/campaignUpdateModal.css'
import campaignsService from '../../services/campaigns.service';

const CampaignUpdateModal = (props) => {
    const { showModal, setShowModal, campaign } = props;
    const [newStartDate, setNewStartDate] = useState(new Date())
    const [newEndDate, setNewEndDate] = useState(new Date())
    const [newName, setNewName] = useState("")
    const [newCategory, setNewCategory] = useState({})
    const [categories, setCategories] = useState([])

    const store = useSelector(state => state)

    useEffect(() => {
        setNewStartDate(campaign.startDate.split("T")[0])
        setNewEndDate(campaign.endDate.split("T")[0])
        setNewName(campaign.name ? campaign.name : "")
        setNewCategory(campaign.category ? campaign.category : {})

        getCategories()
    }, [])

    const getCategories = async () => {
        const response = await adsService.getAdCategories({ jwt: store.user.jwt })
        if (response.status === 200){
            setCategories([...response.data.categories.filter(category => category.id !== campaign.category.id)])
        }else{
            toastService.show("error", "Could retrieve campaign categories")
        }
    }

    const saveChanges = async () => {
        const updateData = {
            id: campaign.id,
            name: newName,
            startDate: newStartDate + "T00:02:00.010Z",
            endDate: newEndDate + "T00:02:00.010Z",
            category: newCategory,
        }

        const response = await campaignsService.updateCampaign({
            campaign: { ...updateData },
            jwt: store.user.jwt
        })
        if (response.status === 200){
            toastService.show("success", "Successfuly updated your campaign! Changes will take place in 24hrs.")
            setTimeout(() => {
                window.location.reload()
            }, 3000)
        }else{
            toastService.show("error", "Could not update your campaign.")
        }
    }
    
    return(
        <Modal className="saveModal" size="lg" show={showModal} onHide={() => setShowModal(!showModal)} animation={true} centered>
            <Modal.Header closeButton>
                <Modal.Title>Update a Campaign</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <div className="updateInput">
                    <label for="name">Name</label>
                    <input
                        type="text" name="name"
                        value={newName} 
                        onChange={(e) => setNewName(e.target.value) } 
                        className="form-control" id="name" />
                </div>
                <div className="updateInput">
                    <label for="startDate">Start Date</label>
                    <input 
                        min={moment(new Date()).add(1, 'd')} 
                        max={moment(new Date()).add(365, 'd')} 
                        type="date" name="startDate"
                        value={newStartDate} 
                        onChange={(e) => setNewStartDate(e.target.value) } 
                        className="form-control" id="startDate" />
                </div>
                {!campaign.isOneTime &&
                (<div className="updateInput">
                    <label for="endDate">End Date</label>
                    <input 
                        min={moment(new Date()).add(1, 'd')} 
                        max={moment(new Date()).add(365, 'd')} 
                        type="date" name="endDate"
                        value={newEndDate} 
                        onChange={(e) => setNewEndDate(e.target.value) } 
                        className="form-control" id="endDate" />
                </div>
                )}
                <div className="dropdown">
                    <label for="category">Category</label>
                    <DropdownButton id="dropdown-basic-button" variant="outline-primary" title={newCategory.name + " "} style={{width: "10em"}}>
                        {categories.map(category => <Dropdown.Item onClick={() => setNewCategory(category)} >{category.name}</Dropdown.Item>)}
                    </DropdownButton>
                </div>
                <p className="note">Please take note that any updates will be applied after 24h cooldown period.</p>
            </Modal.Body>
            <Modal.Footer style={{'background':'#E0E0E0'}}>
                <Button variant="primary" onClick={saveChanges}>Update</Button>
            </Modal.Footer>
        </Modal>
    );
}

export default CampaignUpdateModal;