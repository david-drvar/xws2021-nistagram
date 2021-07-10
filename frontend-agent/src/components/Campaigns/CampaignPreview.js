import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import { useParams, useHistory } from "react-router-dom";
import { Button, Modal } from 'react-bootstrap';
import moment from 'moment';
import Stories from "react-insta-stories";
import PostPreviewGrid from './../Post/PostPreviewGrid'

import renderer from './../StoryCompoent/StoryRenderer';
import Navigation from './../HomePage/Navigation';
import Spinner from './../../helpers/spinner';

import campaignsService from '../../services/nistagram api/campaigns.service';
import toastService from '../../services/toast.service';

import "./../../style/campaignsHome.css"
import "./../../style/campaignsPreview.css"
import CampaignUpdateModal from './CampaignUpdateModal';

const CampaignPreview = (props) => {
    const [showUpdateModal, setShowUpdateModal] = useState(false);
    const [campaign, setCampaign] = useState({});
    const [loading, setLoading] = useState(true);
    const { id } = useParams()
    const [confirmDeleteModal, setConfirmDeleteModal] = useState(false);

    const store = useSelector(state => state)
    const history = useHistory();

    useEffect(() => {
        (async function () {
            const response = await campaignsService.getCampaignById({ jwt: store.apiKey.jwt, id: id })
            if (response && response.status === 200) {
                setCampaign({ ...response.data })
                setLoading(false)
            } else {
                toastService.show("error", "Could not load your campaign.")
            }
        })()
    }, [])

    const displayDate = () => {
        let date = "Dates: " + moment(campaign.startDate).format("DD/MM/YY")
        !campaign.isOneTime ? date += ` - ${moment(campaign.endDate).format("DD/MM/YY")}` : date += ""
        date += `, being placed ${campaign.isOneTime ? "" : "every day"} from ${campaign.startTime < 10 ? "0" + campaign.startTime : campaign.startTime}h - ${campaign.endTime}h`
        return date;
    }

    const renderStories = (ads) => {
        const header = {
            heading: store.apiKey.username,
            profileImage: store.apiKey.photo
        };

        return ads.map(story => (
            <div className="archivedStory">
                <Stories
                    loop={true}
                    renderers={[renderer]}
                    stories={[story.post]}
                    defaultInterval={10000}
                    header={{ ...header, link: story.link }}
                    width={400}
                    height={550} />
            </div>
        ))
    }

    const renderPosts = (ads) => {
        return <PostPreviewGrid isAd={true} posts={ads} />
    }

    const deleteCampaign = async () => {
        const response = await campaignsService.deleteCampaign({ jwt: store.apiKey.jwt, id: campaign.id })
        if (response && response.status === 200) {
            toastService.show("success", "Successfuly deleted your campaign.")
            setTimeout(() => {
                history.push({
                    pathname: `/campaigns`
                })
            }, 2000)
        } else {
            toastService.show("error", "Could not delete your campaign.")
        }
    }

    const showConfirmDeleteModal = () => {
        return (<Modal show={confirmDeleteModal} onHide={() => setConfirmDeleteModal(!confirmDeleteModal)} style={{ 'height': 650 }} >
            <Modal.Header closeButton>
                <Modal.Title>Delete a Campaign</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                Are you sure you want to delete this campaign?
            </Modal.Body>
            <Modal.Footer>
                <Button variant="outline-danger" onClick={deleteCampaign}>Yes</Button>
                <Button variant="primary" onClick={() => setConfirmDeleteModal(false)}>No</Button>
            </Modal.Footer>
        </Modal>)
    }

    return (
        <div>
            <Navigation />
            <main className="CampaignsHome__Wrapper">
                {loading ? <Spinner /> :
                    (<div className="myCampaigns">
                            <div className="header">
                                <div className="title">
                                    {campaign.name}
                                    <div className="stats">
                                        <div>{displayDate()}</div>
                                        <div>{(campaign.isOneTime ? "One time " : "Long term ") + campaign.type + " campaign"}</div>
                                        <div>Category: {campaign.category.name ? campaign.category.name : ""}</div>
                                    </div>
                                </div>
                                <div className="buttons">
                                    <Button onClick={() => setShowUpdateModal(!showUpdateModal)} variant="outline-primary">Update</Button>
                                    <Button onClick={() => setConfirmDeleteModal(!confirmDeleteModal)} variant="outline-danger">Delete</Button>
                                </div>
                            </div>
                            <div className="ads">
                                {campaign.type === "Story" ? renderStories(campaign.ads) : renderPosts(campaign.ads)}
                            </div>
                        </div>
                    )}
                {showUpdateModal && <CampaignUpdateModal showModal={showUpdateModal} setShowModal={setShowUpdateModal} campaign={campaign} />}
                {confirmDeleteModal && showConfirmDeleteModal()}
            </main>
        </div>
    )
}

export default CampaignPreview;