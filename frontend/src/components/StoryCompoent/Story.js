import React, { useState, useEffect } from "react";
import Stories from "react-insta-stories";
import {Button, Modal} from "react-bootstrap";
import renderer from './StoryRenderer';
import storyService from './../../services/story.service'
import '../../style/Story.css';
import '../../style/ProfileIcon.css';
import complaintService from "../../services/complaint.service";
import toastService from "../../services/toast.service";
import {useSelector} from "react-redux";

function Story(props) {
    const { story, iconSize, hideUsername, fixMargins } = props;
    const [showModal, setModal] = useState(false);
    const [showModalReport, setModalReport] = useState(false);
    const [convertedStory, setConvertedStory] = useState([])
    const [header, setHeader] = useState({})
    const [reportCategory, setReportCategory] = useState("");
    const [reportCategoryErr, setReportCategoryErr] = useState("Select report category");
    const [storyId, setStoryId] = useState("");
    const store = useSelector(state => state)

    // Convert story with multiple media to multiple stories with single media, to comply with react-insta-stories
    useEffect(()=>{
        const convertedStories = [];
        story.stories.forEach(singleStory => {  
            convertedStories.push(...storyService.convertStory(singleStory))
        })
        setConvertedStory(convertedStories);

        setHeader({
            heading: story.username,
            profileImage: story.userPhoto
        })
    }, [])

    async function sendReport() {
        if (reportCategory === "") {
            setReportCategoryErr("Select report category");
            return;
        }
        const response = await complaintService.createComplaint({
            id: "",
            category: reportCategory,
            postId: storyId,
            status: "",
            isPost: false,
            userId: store.user.id,
            jwt: store.user.jwt
        });

        if(response.status === 200){
            toastService.show("success", "Story report successfully created.")
        } else{
            toastService.show("error", "Could not create report for this story.")
        }
        setModalReport(false);
    }

    const handleReportCategoryChange = (event) => {
        setReportCategory(event.target.value);
        setReportCategoryErr("");
    }

    const handleModalReport = () => {
        setReportCategory("");
        setReportCategoryErr("");
        setModalReport(!showModalReport)
    }

    return (
        <div>
            <div className={`story ${fixMargins ? "fixMargins" : ""}`}>
                <div className={true ? "storyBorder" : ""}>
                    <img className={`profileIcon ${iconSize ? iconSize : "big"}`} src={story.userPhoto} alt="" onClick={() => setModal(!showModal)}/>
                </div>
                {!hideUsername ? <span className="accountName">{story.username}</span> : null}
            </div>
            
            <Modal show={showModal} onHide={() => setModal(!showModal)}>
                <div  >
                    <Button variant={"outline-danger"} onClick={handleModalReport}>Report story</Button>
                <Stories
                    onAllStoriesEnd={() => setModal(!showModal)}
                    renderers={[renderer]}
                    stories={convertedStory} 
                    defaultInterval={10000} 
                    header={{...header, setStoryId}}
                    width={500}
                    height={700}
                />
                </div>
            </Modal>

            <Modal show={showModalReport} onHide={handleModalReport} style={{ 'height': 650 }} >
                <Modal.Header closeButton>
                    <Modal.Title>Report this post?</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <div className="col-sm-6 mb-2">
                        <select onChange={(e) => handleReportCategoryChange(e)} name={"reportCategory"} value={reportCategory}>
                            <option disabled={true} value="">Select report story</option>
                            <option value="Gore">Gore</option>
                            <option value="Nudity">Nudity</option>
                            <option value="Violence">Violence</option>
                            <option value="Suicide">Suicide</option>
                            <option value="Fake News">Fake News</option>
                            <option value="Spam">Spam</option>
                            <option value="Hate Speech">Hate Speech</option>
                            <option value="Terrorism">Terrorism</option>
                            <option value="Harassment">Harassment</option>
                            <option value="Other">Other</option>
                        </select>
                        {reportCategoryErr.length > 0 && <span className="text-danger">{reportCategoryErr}</span>}
                    </div>
                    <p >Are you sure you want to report this post? </p>
                    <div style={{display:'flex',float:'right'}}>
                        <Button variant="info" style={{marginRight:'10px'}} onClick={sendReport} >Confirm</Button>
                        <Button variant="secondary" onClick={() => setModalReport(!showModalReport)} >Cancel</Button>
                    </div>
                </Modal.Body>
            </Modal>
        </div>
    );
}

export default Story;