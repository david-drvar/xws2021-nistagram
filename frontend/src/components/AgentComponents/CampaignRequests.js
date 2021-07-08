import Navigation from "../HomePage/Navigation";
import {Button, Table} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import userService from "../../services/user.service";
import toastService from "../../services/toast.service";
import {useSelector} from "react-redux";
import ProfileForSug from "../HomePage/ProfileForSug";
import campaignsService from "../../services/campaigns.service";

const CampaignRequests = () => {
    const[requests,setRequests]=useState([])
    const[results,setResults]=useState([])
    const store = useSelector(state => state);

    useEffect(() => {
        getAllRequests()
    }, []);

    async function getAllRequests() {
        const response = await campaignsService.getCampaignRequests({
            agentId:store.user.id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            setRequests(response.data.campaignRequests)
            getUsers(response.data.campaignRequests)
        } else {
            toastService.show("error", "Something went wrong.Please try again!");
        }
    }

    function getUsers(requests) {
        requests.map((request, i) =>
            getUserById(request)
        );
    }

    async function getUserById(request) {
        const response = await userService.getUserById({
            id: request.influencerId,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            console.log(response.data)
            let temp={postAt:request.postAt,status:request.status, campaignId:request.campaignId,id:response.data.id,
                firstName:response.data.firstName,lastName:response.data.lastName,username:response.data.username,profilePhoto:response.data.profilePhoto}

            if(results.some(item => item.influencerId === temp.id)){
                if(results.some(item => item.campaignId === temp.campaignId)) {
                    return;
                }
            }

            setResults(results=>[...results,temp])

            console.log(results)

        } else {
            console.log("getuserbyid error")
        }
    }

    return(
        <div>
            <Navigation/>
            <div style={{marginTop:'5%',marginLeft:'10%', marginRight:'10%', marginBottom:'20%'}}>
                <h3 style={{borderBottom:'1px solid black'}}>Campaign requests</h3>
                <Table striped bordered hover style={{marginTop:'3%'}}>
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>Campaign</th>
                        <th>Influencer</th>
                        <th>Status</th>
                        <th>Post at</th>
                    </tr>
                    </thead>
                    <tbody>
                    {results.map((result,index) =>
                            <tr>
                                <td>{index+1}</td>
                                <td>cekaj jos</td>
                                <td> <ProfileForSug
                                                    username={result.username}
                                                    firstName={result.firstName}
                                                    lastName={result.lastName}
                                                    caption="profile"
                                                    urlText="Follow"
                                                    iconSize="big"
                                                    captionSize="small"
                                                    image={result.profilePhoto} storyBorder={true} /></td>
                                <td>{ result.status==="Pending" &&<p  style={{color:'blueviolet' }}>{result.status}</p> }
                                    { result.status==="Accepted" &&<p  style={{color:'yellowgreen' }}>{result.status}</p> }
                                    { result.status==="Rejected" &&<p  style={{color:'red' }}>{result.status}</p> }
                                </td>
                                <td>{result.postAt}</td>
                            </tr>
                    )}
                    </tbody>
                </Table>
            </div>
        </div>
    );
}
export default CampaignRequests;
