import React, {useEffect, useState} from "react";
import '../../style/Profile.css';
import HomePage from "./HomePage";
import {Button, Modal} from "react-bootstrap";
import Stories from "react-insta-stories";
import EditProfile from "./EditProfile";

function Profile(){
    const [image,setImage] = useState('');
    const[user,setUser]=useState({name:'Jovana',email:'Petrovic', followers:4, following:5});
    const [showModal, setModal]=useState(false);

    const updatePhoto = (file)=>{
        setImage(file)
    }
    useEffect(() => {
        //getUserByUsername
        //onda moras da mu dobavis i br followera i koliko on prati
        //moras da dobavis postove  usera
    });

    function handleModal() {
        setModal(!showModal)
    }

return (
    <div >
        <HomePage/>
    <div style={{marginLeft:'20%', marginRight:'20%'}}>
        <div style={{margin:"18px 0px",orderBottom:"1px solid grey"}}>
            <div style={{ display:"flex",justifyContent:"space-around",}}>
                <div>
                    <img style={{width:"180px",height:"160px",borderRadius:"80px"}}
                         src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                </div>
                <div>
                    <h4>{user.name}</h4>
                    <h5>{user.email}</h5>
                    <div style={{display:"flex",justifyContent:"space-between",width:"108%"}}>
                        <h6>15 posts</h6>
                        <h6>{user.followers} followers</h6>
                        <h6>{user.following} following</h6>
                    </div>
                    <Button variant="link" style={{  borderTop: '1px solid red',display:"flex",justifyContent:"space-between",width:"108%", color:'red',float : "right"}} onClick={handleModal}>Update profile info?</Button>

                </div>
            </div>

        </div>

        <div className="gallery">
            <img className="item"
                 src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>

                <img className="item"
                     src='https://images.unsplash.com/photo-1581882898166-634d30416957?ixid=MnwxMjA3fDB8MHxzZWFyY2h8OXx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>

                    <img className="item"
                         src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>

                    {/*  {
                mypics.map(item=>{
                    return(
                        <img key={item._id} className="item" src={item.photo} alt={item.title}/>
                    )
                })
            }
        */}

            <Modal show={showModal} onHide={handleModal}>
                <Modal.Header closeButton >
                    <Modal.Title>Edit profile</Modal.Title>
                </Modal.Header>
                <Modal.Body >
                    <EditProfile user={user}/>
                </Modal.Body>
                <Modal.Footer >

                </Modal.Footer>
            </Modal>

        </div>

    </div>
    </div>
)
}


export default Profile