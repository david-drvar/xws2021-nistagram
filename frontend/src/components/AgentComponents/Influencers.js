import React, {useEffect, useState} from "react";
import Navigation from "../HomePage/Navigation";
import userService from "../../services/user.service";
import toastService from "../../services/toast.service";
import {useSelector} from "react-redux";
import ProfileForSug from "../HomePage/ProfileForSug";
import {Button, Dropdown, FormControl, Modal} from "react-bootstrap";
import followersService from "../../services/followers.service";
import {ReactComponent as Plus} from "../../images/icons/plus.svg";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import campaignsService from "../../services/campaigns.service";

const Influencers = () => {
    const [influencers, setInfluencers] = useState([])
    const [campaigns, setCampaigns] = useState([])
    const [campaign, setCampaign] = useState([])
    const [renderInfluencers, setRenderInfluencers] = useState([])
    const store = useSelector(state => state);
    const [showModal, setShowModal] = useState(false);
    const [dateTime, setDateTime] = useState(new Date());
    const [startDate, setStartDate] = useState(new Date());
    const [modalUser, setModalUser] = useState({});

    useEffect(() => {
        getInfluencers()
        getCampaigns()
    }, []);



    async function getInfluencers() {
        const response = await userService.getInfluencers({
            //   id: store.user.id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            setInfluencers(response.data.users)
            checkConnection(response.data.users)
        } else {
            toastService.show("error", "Something went wrong.Please try again!");
        }
    }

    async function getCampaigns() {
        const response = await campaignsService.getAgentsCampaigns({
               id: store.user.id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log(response.data.campaigns)
            setCampaigns(response.data.campaigns)
        } else {
            toastService.show("error", "Something went wrong.Please try again!");
        }
    }


     function checkConnection(users){
         users.map((user, i) => {
            if(user.isProfilePublic==false){
                GetFollowersConnection(user)
            }else{
                let temp={id:user.id,username:user.username, firstname:user.firstName, lastname:user.lastName,profilePhoto:user.profilePhoto,isApprovedRequest:true,requestIsPending:false}
                if(renderInfluencers.some(item => item.id === temp.id)){
                    return;
                }
                setRenderInfluencers(renderInfluencers=>[...renderInfluencers, temp])
            }
        });
    }

    async function GetFollowersConnection(value) {
        const response = await followersService.getFollowersConnection({
            userId: store.user.id,
            followerId: value.id,
        })
        if (response.status === 200) {
            let temp={id:value.id,username:value.username,firstname:value.firstName, lastname:value.lastName,profilePhoto:value.profilePhoto, isApprovedRequest:response.data.isApprovedRequest,requestIsPending:response.data.requestIsPending }
            if(renderInfluencers.some(item => item.id === temp.id)){
                    return;
                    
            }
            setRenderInfluencers(renderInfluencers=>[...renderInfluencers, temp])
        } else {
            console.log("followings ne radi")
        }
    }

    function handleModal(user) {
        setModalUser(user)
        setShowModal(!showModal)

    }

    function closeModal() {
        setShowModal(!showModal)
    }

    function setCampaignForInfluencer(campaign) {
        setCampaign(campaign)
    }

    async function createCampaignRequest(modalUser) {
        const response = await campaignsService.createCampaignRequest({
            agentId: store.user.id,
            influencerId:modalUser.id,
            campaignId:campaign.id,
            status:'Pending',
            postAt:dateTime,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            toastService.show("success", "Successfully created!!");

        } else {
            toastService.show("error", "Something went wrong.Please try again!");
        }
    }

    return (
        <div style={{marginTop: '6%'}}>
            <Navigation/>
            <div style={{marginLeft: '20%', marginRight: '20%'}}>
                <h3 style={{borderBottom: '1px solid black'}}>Influencers</h3>
                <div style={{marginTop: '4%'}}>
                    {renderInfluencers.map((user, i) =>
                        <div style={{display: 'flex', borderBottom: '1px solid #dbe0de', marginTop: '5px'}}>
                            <ProfileForSug user={user} username={user.username} firstName={user.firstName}
                                           lastName={user.lastName} urlText="Follow" iconSize="big" captionSize="small" caption="influencer"
                                           image={user.profilePhoto} storyBorder={true}/>
                            {(!user.isProfilePublic && !user.isApprovedRequest && ! user.requestIsPending) &&
                                <div>
                                    <p style={{fontSize: '0.75em', paddingLeft: '250px',paddingBottom: '0.2em', paddingTop: '1.5em', color: 'red'}}>
                                        This account is private. Follow for more info.</p>
                                </div>
                            }
                            { user.requestIsPending &&
                                <div>
                                    <p style={{fontSize: '0.75em', paddingLeft: '250px',paddingBottom: '0.2em', paddingTop: '1.5em', color: 'green'}}>
                                        Follow request is pending</p>
                                </div>
                            }
                            { (user.isProfilePublic || (!user.isProfilePublic && user.isApprovedRequest && !user.requestIsPending)) &&
                                <div style={{paddingLeft: '250px'}}>
                                    <Button
                                        style={{marginLeft: '5px', marginTop: '22px', height: '32px', fontSize: '15px'}}
                                        variant="success"  onClick={() => handleModal(user)}>Hire for compaign </Button>

                                </div>
                            }
                        </div>
                    )}

                </div>
            </div>

            <Modal show={showModal} onHide={closeModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Hire for campaign</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{background:'#e4e6e5'}}>
                    <div>
                        <div style={{borderBottom: '1px solid #dbe0de'}}>
                            <h6>Influencers info:</h6>
                            <tr>
                                <td>Username:</td>
                                <td ><p style={{marginLeft:'15px'}}>{modalUser.username}</p></td>
                            </tr>

                            <tr>
                                <td>First name:</td>
                                <td ><p style={{marginLeft:'15px'}}>{modalUser.firstname}</p></td>
                            </tr>
                            <tr>
                                <td>Last name:</td>
                                <td ><p style={{marginLeft:'15px'}}>{modalUser.lastname}</p></td>
                            </tr>
                        </div>
                        <div style={{background:'#b6beba'}}>
                      <tr>
                          <td>Campaign:</td>
                          <Dropdown>
                              <Dropdown.Toggle variant="link" id="dropdown-basic">
                                  <Plus className="icon" style={{maxWidth:'20px', marginLeft:'20px'}}/>
                              </Dropdown.Toggle>
                                  {campaigns.map((campaign, i) =>
                                      <Dropdown.Menu>
                                      <Dropdown.Item onClick={() => setCampaignForInfluencer(campaign)}>{campaign.name}</Dropdown.Item>
                                      </Dropdown.Menu>

                                  )}

                          </Dropdown>
                      </tr>
                        <tr>
                            <td >  <p style={{marginRight:'10px'}}>Date and time posted at: </p></td>
                            <DatePicker
                                onChange={date => setDateTime(date)}

                                selected={dateTime}
                                onChange={date => setDateTime(date)}
                                showTimeSelect
                                timeFormat="HH:mm"
                                timeIntervals={15}
                                timeCaption="time"
                                dateFormat="MMMM d, yyyy h:mm aa"
                                startDate={startDate}
                                minDate={startDate}

                            />

                        </tr>
                        </div>
                        <tr>
                            <Button
                                style={{marginLeft: '5px', marginTop: '22px', height: '32px', fontSize: '15px'}}
                                variant="success"  onClick={() => createCampaignRequest(modalUser)}>Create campaign request</Button>
                        </tr>
                    </div>

                </Modal.Body>
            </Modal>
        </div>
    );
}
export  default  Influencers;
