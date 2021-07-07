import React, { useEffect, useState } from 'react';
import { useSelector } from "react-redux";
import Navigation from "../HomePage/Navigation";
import { Button, Modal, Dropdown } from "react-bootstrap";

import toastService from "../../services/toast.service";
import productService from "../../services/product.service";
import {useHistory} from "react-router-dom";



// Id        string `gorm:"primaryKey"`
// Name      string
// Price     float32
// IsActive  bool
// Quantity  int
// PhotoLink string
function NewProduct(props) {
    const [user, setUser] = useState({});

    const [quantity, setQuantity] = useState('');
    const [name, setName] = useState('');
    const [image, setImage] = useState('');
    const [price, setPrice] = useState("")
    const history = useHistory()

    const store = useSelector(state => state);

    useEffect(() => {
        console.log(store);
    }, []);



    const createProduct = async () => {
        const productInfo = {
            id: "1",
            name: name,
            price: price,
            isActive: true,
            quantity: quantity,
            photo: image,
            jwt: store.user.jwt,
            agentId : store.user.id
        };

        let response =  await productService.createProduct(productInfo)

        if (response.status === 200) {
            toastService.show("success", `New product successfully created!`);
            history.push({ pathname: '/profile/' + store.user.username });
        }

        else
            toastService.show("error", "Something went wrong, please try again!");
    }

    function handleChangeImage(evt) {
        console.log("Uploading");
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function (upload) {
            setImage(upload.target.result)
        };
        reader.readAsDataURL(file);
    }



    return (
        <div className='home'>
            <Navigation user={user} />

            <div className="card input-filed"
                style={{ margin: "30px auto", maxWidth: "500px", padding: "20px", textAlign: "center", marginTop: "5%" }} >

                <input type="text" placeholder="name" value={name} onChange={(e) => setName(e.target.value)} />
                <br />
                <input type="number" placeholder="price" value={price} onChange={(e) => setPrice(e.target.value)} />
                <br/>
                <input type="number" placeholder="quantity" value={quantity} onChange={(e) => setQuantity(e.target.value)} />
                <br/>

                <br /><br />
                <input type="file" name="file"
                       className="upload-file"
                       id="file"
                       onChange={handleChangeImage}
                       formEncType="multipart/form-data"
                       required />
                <br/>

                <Button type={"primary"} onClick={() => createProduct()}>Submit product</Button>

            </div>

        </div>
    );
}

export default NewProduct;