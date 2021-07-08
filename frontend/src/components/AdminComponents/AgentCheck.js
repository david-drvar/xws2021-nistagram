import React, { useEffect, useState } from 'react';
import {Button, Modal, ListGroup, Table} from "react-bootstrap";
import { useSelector } from 'react-redux';
import { ReactComponent as Check } from './../../images/icons/check.svg'
import collectionsService from './../../services/collections.service'
import favoritesService from './../../services/favorites.service'
import toastService from './../../services/toast.service'
import './../../style/CollectionsModal.css'
import Navigation from "../HomePage/Navigation";
import complaintService from "../../services/complaint.service";
import userService from "../../services/user.service";
import followersService from "../../services/followers.service";

const AgentCheck = (props) => {
    const[users,setUsers]=useState( [])
    const store = useSelector(state => state);


    useEffect(() => {
        getAllPendingRequests()
    }, []);

    async function getAllPendingRequests() {
        const response = await userService.getAllPendingRequests(
            {
                     jwt: store.user.jwt
                 })

        if (response.status === 200) {
           setUsers(response.data.registrationRequests)
            console.log(response.data.registrationRequests)
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }

    async function changeRequestStatus(user, status) {
        console.log(user)
        const response = await userService.agentUpdateRequest({
            id:user.id,
            userId : user.userId,
            status : status,
            jwt : store.user.jwt
        });
        if (response.status === 200) {
            toastService.show("success", "Registration request status changed successfully")
            getAllPendingRequests();
        }
        else
            toastService.show("error", "Something went wrong. Try again")

    }

    return (
        <div>
            <Navigation/>
            <div style={{marginTop:'5%',marginLeft:'10%', marginRight:'20%', marginBottom:'20%'}}>
                <h3 style={{borderBottom:'1px solid black'}}>Registration requests</h3>
                    <Table striped bordered hover>
                        <thead>
                        <tr>
                            <th>#</th>
                            <th>First name</th>
                            <th>Last name</th>
                            <th>Username</th>
                            <th>Email</th>
                            <th>Status</th>
                            <th>Website</th>
                            <th></th>
                        </tr>
                        </thead>
                        <tbody>
                        {users.map((user,index) => {
                            return (
                                <tr>
                                    <td>{index+1}</td>
                                    <td>{user.firstName}</td>
                                    <td>{user.lastName}</td>
                                    <td>{user.username}</td>
                                    <td>{user.email}</td>
                                    <td>{user.status}</td>
                                    <td>
                                        <a className="website" target="_blank" rel="noreferrer"
                                             href={user.website.includes('https://') ? user.website : `https://${user.website}`}>
                                        {user.website}
                                    </a>
                                    </td>
                                    <td>
                                        <Button variant={"success"} onClick={() => changeRequestStatus(user,'Accepted')}>Accept</Button>
                                        <span  style={{marginLeft: '5%'}}/>
                                        <Button variant={"danger"} onClick={() => changeRequestStatus(user,'Refused')}>Refuse</Button>

                                    </td>
                                </tr>
                            )
                        })}
                        </tbody>
                    </Table>
                </div>
        </div>

    )
}

export default AgentCheck;