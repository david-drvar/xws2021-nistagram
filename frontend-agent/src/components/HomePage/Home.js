import Navigation from "./Navigation";
import React, { useEffect, useState } from "react";
import { Button, Modal } from "react-bootstrap";
import PasswordStrengthBar from "react-password-strength-bar";
import { useDispatch, useSelector } from 'react-redux';
import userService from "../../services/agent.service";
import "../../style/home.css";
import productService from "../../services/product.service";
import toastService from "../../services/toast.service";
import Spinner from "../../helpers/spinner";
import ProductPreviewGrid from "../Product/ProductPreviewGrid";

const Home = () => {
    const [showModal, setModal] = useState(false);
    const [submitted, setSubmitted] = useState(false);
    const [passwordStrength, setPasswordStrength] = useState('');
    const [passwords, setPasswords] = useState({ oldPassword: '', newPassword: '', repeatedPassword: '' });
    const [oldErr, setOldErr] = useState('');
    const [newErr, setNewErr] = useState('');
    const [repErr, setRepErr] = useState('');
    const [products, setProducts] = useState([]);
    const [loadingPosts, setLoadingPosts] = useState(true);


    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        if (store.user.jwt !== "") {
            getAllProducts();
        }
    },[]);



    async function getAllProducts() {
        const response = await productService.getAllProducts({
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setProducts([...response.data.products])
            setLoadingPosts(false);
        } else {
            console.log(response);
            toastService.show("error", "Could not retrieve all products.")
        }
    }




    return (
        <div className="App">
            <Navigation/>
            <main>
                { /*showModalDialog()*/ }
                <div className="profileGrid">

                    <div className="content">
                        <div className="posts">
                            {loadingPosts ?
                                <div style={{position: "relative", left: "45%", marginTop: "50px"}}>
                                    <Spinner type="MutatingDots" height="100" width="100"/>
                                </div> :
                                <ProductPreviewGrid posts={products}/>
                            }
                        </div>
                    </div>
                </div>
            </main>
        </div>
    );
}

export default Home;