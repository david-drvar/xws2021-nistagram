import React, { useEffect, useState } from "react";
import toastService from './../../services/toast.service';

import '../../style/Profile.css';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";
import {useHistory, useParams} from 'react-router-dom'
import Spinner from "../../helpers/spinner";
import ProductPreviewGrid from "./ProductPreviewGrid";
import productService from "../../services/product.service";
import {Button, Modal} from "react-bootstrap";
import RegistrationPage from "../../pages/RegistrationPage";


const Product = () => {
    const {id} = useParams()

    const dispatch = useDispatch()
    const store = useSelector(state => state);
    const [product, setProduct] = useState({});
    const [productNameEdit, setProductNameEdit] = useState("");
    const [productNameEditErr, setProductNameEditErr] = useState("Enter product name");
    const [productPriceEdit, setProductPriceEdit] = useState("");
    const [productPriceEditErr, setProductPriceEditErr] = useState("Enter price");
    const [productQuantityEdit, setProductQuantityEdit] = useState("");
    const [productQuantityEditErr, setProductQuantityEditErr] = useState("Enter quantity");
    const [productImageEdit, setProductImageEdit] = useState("");
    const [showEditModal, setShowEditModal] = useState(false);
    const [showOrderModal, setShowOrderModal] = useState(false);
    const [submitted, setSubmitted] = useState(false);

    const [orderQuantity, setOrderQuantity] = useState("");

    const history = useHistory()

    useEffect(() => {
        getProductById(id);
    },[]);

    async function getProductById(id) {
        const response = await productService.getProductById({
            id : id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setProduct(response.data)
            console.log(response.data)
        } else {
            console.log(response);
            toastService.show("error", "Could not retrieve products")
        }
    }

    async function deleteProduct() {
        if (window.confirm('Are you sure you want to delete this product?')) {
            const response = await productService.deleteProduct({
                id : id,
                jwt: store.user.jwt,
            })

            if (response.status === 200) {
                toastService.show("success", "Product successfully deleted")
                history.push({ pathname: '/profile/' + store.user.username })
            } else {
                console.log(response);
                toastService.show("error", "Could not retrieve products")
            }
        }
    }

    function handleEditModal() {
        setProductNameEdit(product.name);
        setProductQuantityEdit(product.quantity);
        setProductPriceEdit(product.price);
        setShowEditModal(!showEditModal)
    }

    function handleOrderModal() {
        setShowOrderModal(!showOrderModal);
    }

    function handleInputChange(event) {
        const target = event.target;
        switch (target.name) {
            case "productQuantityEdit" :
                setProductQuantityEdit(target.value);
                break;
            case "productPriceEdit" :
                setProductPriceEdit(target.value);
                break;
            case "productNameEdit" :
                setProductNameEdit(target.value);
                break;

        }
        validationErrorMessage(event);
    }

    function validationErrorMessage(event) {
        const { name, value } = event.target;

        switch (name) {
            case 'productQuantityEdit':
                setProductQuantityEditErr(productQuantityEdit !== "" ? '' : 'Enter quantity')
                break;
            case 'productPriceEdit':
                setProductPriceEditErr(productPriceEdit !== "" ? '' : 'Enter price')
                break;
            case 'productNameEdit':
                setProductNameEditErr(productNameEdit !== "" ? '' : 'Enter name')
                break;
            default:
                /*this.setState({
                    validForm: true
                })*/
                break;
        }

    }

    function validateForm(errors) {
        let valid = true;
        for(const Error of errors) {
            validationErrorMessage(createTarget(Error));
        }

        if(productQuantityEditErr !== "" || productPriceEditErr !== "" || productNameEditErr !== "")
            return !valid;
        return valid;
    }

    function createTarget (error) {
        return {target : {value : error, name : error}}
    }

    async function submitEdit() {
        setSubmitted(true);

        const errors = ['productQuantityEdit', 'productPriceEdit', 'productNameEdit'];
        if (validateForm(errors)) {
            const response = await productService.updateProduct({
                id : id,
                name : productNameEdit,
                quantity: productQuantityEdit,
                price : productPriceEdit,
                photo : productImageEdit,
                jwt: store.user.jwt,
            })

            if (response.status === 200) {
                toastService.show("success", "Product successfully updated")
                setShowEditModal(!showEditModal);
                setSubmitted(false);
                getProductById(id);
            } else {
                console.log(response);
                toastService.show("error", "Could not retrieve products")
            }
        } else {
            console.log('Invalid Form')
        }
    }

    function handleChangeImage(evt) {
        console.log("Uploading");
        var self = this;
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function(upload) {
            setProductImageEdit(upload.target.result)
        };
        reader.readAsDataURL(file);
    }

    async function submitOrder() {
        if (orderQuantity > product.quantity) {
            alert("Cannot exceed quantity of product");
            return;
        }

        const response = await productService.orderProduct({
            productId : id,
            userId : store.user.id,
            quantity: orderQuantity,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            toastService.show("success", "Product successfully ordered")
            setShowOrderModal(!showOrderModal);

            history.push({ pathname: '/my-orders' })


        } else {
            console.log(response);
            toastService.show("error", "Could not retrieve products")
        }
    }


    return (
        <div>
            <Navigation/>
            <div className="profileGrid">
                <div className="card">
                    {console.log(product)}
                    <img src={product.photo} alt="product photo" style={{width: "50%"}}/>
                        <h1>{product.name}</h1>

                    {product.agentId === store.user.id && <Button style={{width: '250px'}} variant={"outline-primary"} onClick={handleEditModal}>Edit</Button>}
                    {product.agentId === store.user.id && <Button style={{width: '250px'}} variant={"outline-danger"} onClick={deleteProduct}>Delete</Button>}
                        <p className="price">${product.price}</p>
                        <p>Quantity : {product.quantity}</p>
                        <p>
                            {store.user.role === 'Basic' && <Button onClick={handleOrderModal}>Add to Cart</Button>}
                        </p>
                </div>
            </div>

            <Modal show={showEditModal} onHide={setShowEditModal} style={{ 'height': 650 }} >
                <Modal.Header closeButton style={{ 'background': 'silver' }}>
                    <Modal.Title>Edit product</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{ 'background': 'silver' }}>
                    <div className="row">
                        <label className="col-sm-2 col-form-label">*Name</label>
                        <div className="col-sm-5 mb-2">
                            <input  type="text" value={productNameEdit} name="productNameEdit" onChange={(e) =>
                                handleInputChange(e) } className="form-control" placeholder="product name"/>
                            {submitted && productNameEditErr.length > 0 && <span className="text-danger">{productNameEditErr}</span>}

                        </div>
                        <div className="col-sm-5 mb-2">
                            <input   type="number" value={productPriceEdit} name="productPriceEdit" onChange={(e) => handleInputChange(e) } className="form-control" placeholder="Last Name"/>
                            {submitted && productPriceEditErr.length > 0 && <span className="text-danger">{productPriceEditErr}</span>}
                        </div>
                        <div className="col-sm-4">
                            <input   type="number" value={productQuantityEdit} name="productQuantityEdit" onChange={(e) => handleInputChange(e) } className="form-control" placeholder="Last Name"/>
                            {submitted && productQuantityEditErr.length > 0 && <span className="text-danger">{productQuantityEditErr}</span>}
                        </div>
                    </div>

                    <div className="row" style={{marginTop: '1rem'}}>
                        <label  className="col-sm-2 col-form-label">*photo</label>
                        <div className="col-sm-6 mb-2">
                            {/*<input type="file" onChange={(e) => setProfilePhoto(e.target.files[0])} />*/}
                            <input type="file" name="file"
                                   className="upload-file"
                                   id="file"
                                   onChange={handleChangeImage}
                                   formEncType="multipart/form-data"
                                   required />
                        </div>
                        <div className="col-sm-4">
                        </div>
                    </div>
                    <br/>
                    <Button onClick={submitEdit}>Submit</Button>
                </Modal.Body>
            </Modal>



            <Modal show={showOrderModal} onHide={setShowOrderModal} style={{ 'height': 650 }} >
                <Modal.Header closeButton style={{ 'background': 'silver' }}>
                    <Modal.Title>Order product</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{ 'background': 'silver' }}>
                    <div className="row">
                        <label className="col-sm-2 col-form-label">*Quantity</label>
                        <div className="col-sm-5 mb-2">
                            <input  type="number" value={orderQuantity} name="orderQuantity" onChange={(e) =>
                                setOrderQuantity(e.target.value) } className="form-control" placeholder="order quantity"/>
                        </div>
                    </div>
                    <br/>
                    <Button onClick={submitOrder}>Submit</Button>
                </Modal.Body>
            </Modal>
        </div>
    );
}
export default Product;
