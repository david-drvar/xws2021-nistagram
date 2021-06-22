import React, {useEffect, useState} from 'react';
import { useDispatch, useSelector } from 'react-redux';
import axios from "axios";
import Navigation from "../HomePage/Navigation";
import {Button} from "react-bootstrap";
import {user} from "../../store/reducers/user.reducer";
import userService from "../../services/user.service";
import {useHistory} from "react-router-dom";
import verificationRequestService from "../../services/verificationRequest.service";
import toastService from "../../services/toast.service";


function SubmitVerificationRequest() {
    const categories = ["Sports", "Influencer", "News","Brand","Business","Organization","Government"]
    const [selectedCategory, setSelectedCategory] = useState("")
    const [image, setImage] = useState("")
    const [categoryErr, setCategoryErr] = useState("Select category")
    const [imageErr, setImageErr] = useState("Select document photo")
    const [submitted, setSubmitted] = useState(false);

    const dispatch = useDispatch()
    const store = useSelector(state => state);
    const history = useHistory()


    useEffect(() => {
        setImageErr( image !== "" ? '' : 'Select document photo')
        setCategoryErr( selectedCategory !== "" ? '' : 'Select category')
    }, [image,selectedCategory])

    const handleInputChange = (event) => {
        const target = event.target;
        setSelectedCategory(target.value)
        validationErrorMessage(event);
    }

    function validationErrorMessage(event) {
        const { name, value } = event.target;

        switch (name) {
            case 'category':
                setCategoryErr(selectedCategory !== "" ? '' : 'Select category')
                break;
            case 'file':
                setImageErr(image !== "" ? '' : 'Select document photo')
                break;
        }
    }

    function handleChangeImage(evt) {
        console.log("Uploading");
        var self = this;
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function(upload) {
            setImage(upload.target.result)
        };
        reader.readAsDataURL(file);
        validationErrorMessage(evt);

    }

    function createTarget (error) {
        return {target : {value : error, name : error}}
    }

    function validateForm(errors) {
        let valid = true;
        for(const Error of errors) {
            validationErrorMessage(createTarget(Error));
        }
        if(categoryErr !== "" || imageErr !== "")
            return !valid;
        return valid;
    }

    async function submitVerificationRequest() {
        setSubmitted(true);

        const errors = ['category', 'file'];
        if (validateForm(errors)) {
            const response = await verificationRequestService.submitVerificationRequest({
                userId: store.user.id,
                documentPhoto: image,
                category: selectedCategory,
                jwt : store.user.jwt
            });
            if (response.status === 200) {
                toastService.show("success", "Verification request submitted successfully")
                history.push({
                    pathname: '/view-my-verification-request'
                })
            }
            else
                toastService.show("error", "Something went wrong. Try again")



        }
    }

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation/>
            <div style={{marginTop:'5%', marginLeft: '5%'}}>
                <h1>Submit verification request</h1>
                <br/>
                <div className="row" style={{marginTop: '1rem'}}>
                    <label  className="col-sm-2 col-form-label">Category</label>
                    <div className="col-sm-6 mb-2">
                        <select onChange={(e) => handleInputChange(e)} name={"category"} value={selectedCategory}>
                            <option disabled={true} value="">Select category</option>
                            {categories.map(category => {
                                return (
                                    <option value={category}>{category}</option>
                                )
                            })}
                        </select>
                        {submitted && categoryErr.length > 0 && <span className="text-danger">{categoryErr}</span>}
                    </div>
                    <div className="col-sm-4">
                    </div>
                </div>

                <div className="row" style={{marginTop: '1rem'}}>
                    <label  className="col-sm-2 col-form-label">Document photo</label>
                    <div className="col-sm-6 mb-2">
                        <input type="file" name="file"
                               className="upload-file"
                               id="file"
                               onChange={handleChangeImage}
                               formEncType="multipart/form-data"
                               required />
                        {submitted && imageErr.length > 0 &&  <span className="text-danger">{imageErr}</span>}

                    </div>

                    <div className="col-sm-4">
                    </div>
                </div>

                <br/>
                <Button type={"primary"} onClick={submitVerificationRequest}>Submit</Button>


            </div>
        </div>

    );
}

export default SubmitVerificationRequest;