import React, { useState, useEffect } from 'react';
import { useSelector } from "react-redux"
import {ListGroup, Button, FormControl, Table} from 'react-bootstrap';
import Navigation from "../HomePage/Navigation";
import productService from "../../services/product.service";


import './../../style/Saved.css'
import ProfileInfo from "./ProfileInfo";
import toastService from "../../services/toast.service";
import {useHistory} from "react-router-dom";

const Orders = () => {
    const store = useSelector(state => state);
    const [orders,setOrders] = useState([]);
    const history = useHistory()

    useEffect(() => {
        if (store.user.role === "Basic")
            getOrdersByUser();
        else if (store.user.role === "Agent")
            getOrdersByAgent();
    }, [orders]);

    async function getOrdersByUser() {
        const response = await productService.getOrdersByUser({
            id : store.user.id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setOrders(response.data.orders)
            console.log(response.data)
        } else {
            console.log(response);
            toastService.show("error", "Could not retrieve orders")
        }
    }

    async function getOrdersByAgent() {
        const response = await productService.getOrdersByAgent({
            id : store.user.id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setOrders(response.data.orders)
            console.log(response.data)
        } else {
            console.log(response);
            toastService.show("error", "Could not retrieve orders")
        }
    }

    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                <div>

                    {store.user.role === "Basic" && <Table striped bordered hover>
                        <thead>
                        <tr>
                            <th>#</th>
                            <th>Product name</th>
                            <th>Quantity</th>
                            <th>Total price</th>
                            <th>Date created</th>
                            <th>Agent username</th>
                        </tr>
                        </thead>
                        <tbody>
                        {orders.map((order,index) => {
                            return (
                                <tr>
                                    <td>{index+1}</td>
                                    <td>{order.productName}</td>
                                    <td>{order.quantity}</td>
                                    <td>{order.totalPrice}</td>
                                    <td>{order.dateCreated}</td>
                                    <td onClick={() => history.push({ pathname: '/profile/' + order.username })}>{order.username}</td>
                                </tr>
                            )
                        })}
                        </tbody>
                    </Table>}


                    {store.user.role === "Agent" && <Table striped bordered hover>
                        <thead>
                        <tr>
                            <th>#</th>
                            <th>Product name</th>
                            <th>Quantity</th>
                            <th>Total price</th>
                            <th>Date created</th>
                            <th>User username</th>
                        </tr>
                        </thead>
                        <tbody>
                        {orders.map((order,index) => {
                            return (
                                <tr>
                                    <td>{index+1}</td>
                                    <td>{order.productName}</td>
                                    <td>{order.quantity}</td>
                                    <td>{order.totalPrice}</td>
                                    <td>{order.dateCreated}</td>
                                    <td>{order.username}</td>
                                </tr>
                            )
                        })}
                        </tbody>
                    </Table>}
                </div>
            </div>
        </div>
    );
}

export default Orders;