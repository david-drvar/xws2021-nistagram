import ProfileInfo from "./ProfileInfo";
import React, {useEffect, useState} from "react";
import {Button} from "react-bootstrap";
import userService from "../../services/user.service";
import {useSelector} from "react-redux";
import toastService from "../../services/toast.service";

function EditProfileImage(){
    const [profilePhoto, setProfilePhoto] = useState("");
    const store = useSelector(state => state);

    function handleChangeImage(evt) {
        console.log("Uploading");
        var self = this;
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function(upload) {
            setProfilePhoto(upload.target.result)
        };
        reader.readAsDataURL(file);
    }

    async function updatePhoto() {
        const response = await userService.updatePhoto({
            userId: store.user.id,
            photo:profilePhoto,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            console.log("BRAVO")
            toastService.show("success", "Successfully updated!");
        } else {
            console.log("NE BRAVO")
            toastService.show("error", "Could not update image!");

        }
    }

    return(
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                    <label style={{marginTop:'20px',fontSize:'20px', fontWeight:'bold'}}>Please choose new profile photo</label>
                    <div className="col-sm-6 mb-2">
                        <input type="file" name="file"  className="upload-file" id="file"  onChange={handleChangeImage} formEncType="multipart/form-data" required />
                    </div>
                       <img src={profilePhoto}/>

                <Button style={{marginTop:'10px'}}variant="secondary" onClick={updatePhoto}>Update</Button>

            </div>
        </div>
    );

}export  default  EditProfileImage;