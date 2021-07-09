import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import { ListGroup, Button, Modal } from 'react-bootstrap';
import { useHistory } from "react-router-dom";
import moment from 'moment';

import Navigation from './../HomePage/Navigation';
import Spinner from './../../helpers/spinner';

import campaignsService from '../../services/campaigns.service';
import toastService from '../../services/toast.service';

import "./../../style/campaignsHome.css"

const CampaignsHome = (props) => {
    const [campaigns, setCampaigns] = useState([]);
    const [loading, setLoading] = useState(true);

    const history = useHistory();
    const store = useSelector(state => state)

    useEffect(() => {
        (async function() {
            const response = await campaignsService.getAgentsCampaigns({ jwt: store.user.jwt })
            if (response && response.status === 200){
                setCampaigns([...response.data.campaigns])
                setLoading(false)
            }else{
                toastService.show("error", "Could not load your campaigns.")
            }
        })()
    }, [])

    const previewCampaign = (id) => {
        history.push({
            pathname: `/campaigns/preview/${id}`
        })
    }

    const displayDate = (campaign) => {
        let date = moment(campaign.startDate).format("DD/MM/YY")
        !campaign.isOneTime ? date += ` - ${moment(campaign.endDate).format("DD/MM/YY")}` : date += ""
        date += `, being placed ${campaign.isOneTime ? "" : "every day"} from ${campaign.startTime < 10 ? "0" + campaign.startTime : campaign.startTime}h - ${campaign.endTime}h`
        return date;
    }

    const createCampaign = () => {
        history.push({
            pathname: "/campaigns/create"
        })
    }

    return (
        <div>
            <Navigation/>
            <main className="CampaignsHome__Wrapper">
                <div className="myCampaigns">
                    <div className="homeHeader">
                        <div className="title">My Campaigns</div>
                        <div className="homeButtons">
                            <Button variant="primary" onClick={() => createCampaign()}>Create New Campaign</Button>
                        </div>
                    </div>
                    { loading ? <Spinner/> : 
                    (<ListGroup className="list" >
                        { campaigns.map(campaign => {
                            return (
                                <ListGroup.Item 
                                    action
                                    onClick={() => previewCampaign(campaign.id)}
                                    className="campaignCard">
                                        <div className="name">{ campaign.name }</div>
                                        <div className="info">
                                            <div>{ displayDate(campaign) }</div>
                                            <div>{ (campaign.isOneTime ? "One time " : "Long term ") + campaign.type + " campaign" } </div>
                                            <div>Category: {campaign.category.name ? campaign.category.name : ""}</div>
                                        </div>
                                </ListGroup.Item>
                            )
                        }) }
                    </ListGroup>)}
                </div>
            </main>
        </div>
    )
}

export default CampaignsHome;