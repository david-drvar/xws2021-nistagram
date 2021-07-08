import React, { useEffect, useState } from 'react';
import {Button, Modal, ListGroup, Table, Dropdown} from "react-bootstrap";
import { useSelector } from 'react-redux';
import toastService from './../../services/toast.service'
import './../../style/CollectionsModal.css'
import Navigation from "../HomePage/Navigation";
import complaintService from "../../services/complaint.service";
import userService from "../../services/user.service";
import storyService from "../../services/story.service";
import postService from "../../services/post.service";
import PostPreviewModal from "../Post/PostPreviewModal";
import {ReactComponent as Plus} from "../../images/icons/plus.svg";
import DatePicker from "react-datepicker";
import ProfileForSug from "../HomePage/ProfileForSug";

const ComplaintPreview = (props) => {
    const [complaints, setComplaints] = useState([]);
    const [users,setUsers]=useState([]);
    const [user,setUser]=useState([]);
    const store = useSelector(state => state);
    const [post, setPost] = useState({})
    const [story, setStory] = useState({})
    const [showModal, setShowModal] = useState(false);
    const [showModalStory, setShowModalStory] = useState(false);
    const [modalStory, setModalStory] = useState(false);
    const [modalUser, setModalUser] = useState(false);
    const [userForModal, setUserForModal] = useState(false);

    useEffect(() => {
        getComplaints()
    }, []);

    function getComplaints() {
        complaintService.getAllContentComplaints({ jwt: store.user.jwt })
            .then(response => {
                if(response.status === 200)
                    setComplaints(response.data.contentComplaints)
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }


     async function getPostById(id,flag) {
         await postService.getPostById({id: id, jwt: store.user.jwt})
             .then(response => {
                 setPost(response.data)
                 setUser(response.data.userId)
                 if(flag===true)
                     setShowModal(true);
             })
             .catch(err => {
                 toastService.show("error", "Error")
             })
    }

    async function getStoryById(id,flag) {
        await storyService.getStoryById({id: id, jwt: store.user.jwt})
            .then(response => {
                setStory(response.data)
                setUser(response.data.userId)

                if(flag===true){
                    setModalStory(response.data.media[0].content)
                    setShowModalStory(!showModalStory)
                }
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }
     function getUser(complaint) {
       if (complaint.isPost==true){
           getPostById(complaint.postId,false)
       }else{
           getStoryById(complaint.postId,false)
       }
       getUserById(user)
       setModalUser(!modalUser)
    }

    async function getUserById(id) {
        const response = await userService.getUserById({
            id: id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUserForModal(response.data)
        } else {
            console.log("getuserbyid error")
        }
    }

    function  changeStatus(post,status){
        if(status==="Refused"){
            rejectComplaint(post.id)
        }else if(status==="Block"){
            changeUserActiveStatus(post.userId)
        }else if(status==="Delete"){
            if(post.isPost===false){
                deleteStory(post.postId)
            }else{
                deletePost(post.postId)
            }
        }
    }

    async function rejectComplaint(complaintId) {
        await complaintService.rejectById({id: complaintId, jwt: store.user.jwt})
            .then(response => {
                toastService.show("success", "Successfully rejected!")
                getComplaints()
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }

    async function changeUserActiveStatus(userId) {
        await userService.changeUserActiveStatus({id: userId, jwt: store.user.jwt})
            .then(response => {
                toastService.show("success", "Successfully updated!")
                getComplaints()
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }

    async function deleteStory(storyId) {
        await storyService.deleteStory({id: storyId, jwt: store.user.jwt})
            .then(response => {
                toastService.show("success", "Successfully updated!")
                getComplaints()
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }

    async function deletePost(postId) {
        await postService.deletePost({id: postId, jwt: store.user.jwt})
            .then(response => {
                toastService.show("success", "Successfully updated!")
                getComplaints()
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }

    function closeModalStory() {
        setShowModalStory(!showModalStory)
    }

    function closeModalUser() {
        setModalUser(!modalUser)
    }

    return (
        <div>
            <Navigation/>
            <div style={{marginTop:'5%',marginLeft:'10%', marginRight:'10%', marginBottom:'20%'}}>
                <h3 style={{borderBottom:'1px solid black'}}>Complaints</h3>
                <Table striped bordered hover>
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>Category</th>
                        <th>Post/Story</th>
                        <th>User</th>
                        <th>Status</th>
                        <th></th>
                    </tr>
                    </thead>
                    <tbody>
                    {complaints.map((complaint,index) => {
                        return (
                            <tr>
                                <td>{index+1}</td>
                                <td>{complaint.category}</td>
                                <td>
                                    { complaint.isPost===true ?
                                        <Button variant="link" style={{color: 'black'}} onClick={() =>getPostById(complaint.postId,true)} >post</Button>
                                        :
                                        <Button variant="link" style={{color: 'black'}} onClick={() =>getStoryById(complaint.postId,true)} >story</Button>
                                    }
                                </td>
                                <td>
                                    <Button variant="link" style={{color: 'black'}} onClick={() =>getUser(complaint)} >click for info</Button>

                                   </td>
                                <td>{complaint.status}</td>
                                {complaint.status === "Pending" && <td>
                                    <Button variant={"danger"} onClick={() => changeStatus(complaint, 'Delete')}>Delete
                                        content</Button>
                                    <span style={{marginLeft: '5%'}}/>
                                    <Button variant={"danger"} onClick={() => changeStatus(complaint, 'Block')}>Block
                                        account</Button>
                                    <span style={{marginLeft: '5%'}}/>
                                    <Button variant={"success"} onClick={() => changeStatus(complaint, 'Refused')}>Refuse
                                        complaint</Button>

                                </td>
                                }
                            </tr>
                        )
                    })}
                    </tbody>
                </Table>
            </div>

            <PostPreviewModal
                post={post}
                postUser={{id: post.userId}}
                showModal={showModal}
                setShowModal={setShowModal}/>

            <Modal show={showModalStory} onHide={closeModalStory}>
                <Modal.Header closeButton>
                    <Modal.Title>Story</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <img src={modalStory}  style={{width:'400px', height: '400px'}}/>
                </Modal.Body>
            </Modal>
            <Modal show={modalUser} onHide={closeModalUser}>
                <Modal.Header closeButton>
                    <Modal.Title>User</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <ProfileForSug
                                   username={userForModal.username}
                                   firstName={userForModal.firstName}
                                   lastName={userForModal.lastName}
                                   caption="see profile"
                                   urlText="Follow"
                                   iconSize="big"
                                   captionSize="small"
                                   image={userForModal.profilePhoto} storyBorder={true} />
                </Modal.Body>
            </Modal>
        </div>


    );
}

export default ComplaintPreview;