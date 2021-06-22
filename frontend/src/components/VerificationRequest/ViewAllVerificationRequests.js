import React, {useEffect, useState} from 'react';
import { useDispatch, useSelector } from 'react-redux';
import axios from "axios";
import Navigation from "../HomePage/Navigation";
import {Button, Table} from "react-bootstrap";
import {user} from "../../store/reducers/user.reducer";
import userService from "../../services/user.service";
import verificationRequestService from "../../services/verificationRequest.service";



function ViewAllVerificationRequests() {
    const [verificationRequests, setVerificationRequests] = useState([]);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        getAllVerificationRequests();
    }, [])

    async function getAllVerificationRequests() {
        const response = await verificationRequestService.getAllVerificationRequests({jwt : store.user.jwt})
        setVerificationRequests(response.data.verificationRequests);
    }

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation/>
            <h1>All verification requests</h1>

            <Table striped bordered hover>
                <thead>
                <tr>
                    <th>#</th>
                    <th>First name</th>
                    <th>Last name</th>
                    <th>Category</th>
                    <th>Status</th>
                    <th>Created at</th>
                    <th>Document photo</th>
                    <th>Profile</th>
                </tr>
                </thead>
                <tbody>
                {verificationRequests.map((verificationRequest,index) => {
                    return (
                        <tr>
                            <td>{index+1}</td>
                            <td>{verificationRequest.firstName}</td>
                            <td>{verificationRequest.lastName}</td>
                            <td>{verificationRequest.category}</td>
                            <td>{verificationRequest.status}</td>
                            <td>{verificationRequest.createdAt}</td>
                            <td>
                                <img  src={verificationRequest.documentPhoto} alt="document photo" style={{width:'200px', height: '200px'}}/>
                            </td>
                            <td>Profile</td>
                        </tr>
                    )
                })}
                </tbody>
            </Table>
        </div>
    );
}

export default ViewAllVerificationRequests;