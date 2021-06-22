import React, {useEffect, useState} from 'react';
import { useDispatch, useSelector } from 'react-redux';
import axios from "axios";
import Navigation from "../HomePage/Navigation";
import {Button, Table} from "react-bootstrap";
import {user} from "../../store/reducers/user.reducer";
import userService from "../../services/user.service";
import verificationRequestService from "../../services/verificationRequest.service";



function ViewMyVerificationRequests() {
    const [verificationRequests, setVerificationRequests] = useState([]);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        getVerificationRequestsByUser();
    }, [])

    async function getVerificationRequestsByUser() {
        const response = await verificationRequestService.getVerificationRequestsByUser({
            userId : store.user.id,
            jwt : store.user.jwt
        });
        setVerificationRequests(response.data.verificationRequests);
    }

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation/>
            <h1>My verification requests</h1>

            <Table striped bordered hover>
                <thead>
                <tr>
                    <th>#</th>
                    <th>Category</th>
                    <th>Status</th>
                    <th>Created at</th>
                    <th>Document photo</th>
                </tr>
                </thead>
                <tbody>
                {verificationRequests.map((verificationRequest,index) => {
                    return (
                        <tr>
                            <td>{index+1}</td>
                            <td>{verificationRequest.category}</td>
                            <td>{verificationRequest.status}</td>
                            <td>{verificationRequest.createdAt}</td>
                            <td>
                                <img  src={verificationRequest.documentPhoto} alt="document photo" style={{width:'200px', height: '200px'}}/>
                            </td>
                        </tr>
                    )
                })}
                {/*<tr>*/}
                {/*    <td>1</td>*/}
                {/*    <td>Mark</td>*/}
                {/*    <td>Otto</td>*/}
                {/*    <td>@mdo</td>*/}
                {/*</tr>*/}
                {/*<tr>*/}
                {/*    <td>2</td>*/}
                {/*    <td>Jacob</td>*/}
                {/*    <td>Thornton</td>*/}
                {/*    <td>@fat</td>*/}
                {/*</tr>*/}
                {/*<tr>*/}
                {/*    <td>3</td>*/}
                {/*    <td colSpan="2">Larry the Bird</td>*/}
                {/*    <td>@twitter</td>*/}
                {/*</tr>*/}
                </tbody>
            </Table>
        </div>
    );
}

export default ViewMyVerificationRequests;