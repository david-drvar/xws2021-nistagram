import  React from "react";
import IndexPage from './pages/IndexPage.js'
import {ForgotPasswordPage} from './components/forgotPass/ForgotPasswordPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import RegistrationPage from "./pages/RegistrationPage";
import Chats from "./components/HomePageComponents/Chats";
import Saved from "./components/HomePageComponents/Saved";
import StoryArchive from "./components/HomePageComponents/StoryArchive";
import Search from "./pages/Search";
import Profile from "./components/ProfileComponent/Profile";
import NewPost from "./components/Post/NewPost";
import Home from "./components/HomePage/Home";
import UnauthorizedPage from "./helpers/UnauthorizedPage";
import BlockedUsers from "./components/UserData/BlockedUsers";
import CloseFriends from "./components/UserData/CloseFriends";
import Liked from "./components/UserData/Liked";
import Disliked from "./components/UserData/Disliked";
import Archived from "./components/UserData/Archived";
import ProfileInfo from "./components/UserData/ProfileInfo";
import EditProfile from "./components/UserData/EditProfile";
import ChangePassword from "./components/UserData/ChangePassword";
import EditUserPrivacy from "./components/UserData/EditUserPrivacy";
import SubmitVerificationRequest from "./components/VerificationRequest/SubmitVerificationRequest";
import ViewMyVerificationRequests from "./components/VerificationRequest/ViewMyVerificationRequests";
import ViewPendingVerificationRequests from "./components/VerificationRequest/ViewPendingVerificationRequests";
import ViewAllVerificationRequests from "./components/VerificationRequest/ViewAllVerificationRequests";
import EditProfileImage from "./components/UserData/EditProfileImage";
import Notifications from "./components/Notifications/Notifications";
import CampaignsHome from './components/Campaigns/CampaignsHome'
import CampaignPreview from './components/Campaigns/CampaignPreview'
import CreateCampaign from './components/Campaigns/CreateCampaign';
import AuthenticatedRoute from "./routes/AuthenticatedRoute";
import AgentRoute from "./routes/AgentRoute";
import CampaignRequests from "./components/AgentComponents/CampaignRequests";
import Influencers from "./components/AgentComponents/Influencers";
import ComplaintPreview from "./components/AdminComponents/ComplaintPreview";
import AgentCheck from "./components/AdminComponents/AgentCheck";
import AgentReg from "./components/AdminComponents/AgentReg";
import AdminRoute from "./routes/AdminRoute";


const App = () => {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact  component={Home}/>
                <Route path='/login' exact={true} component={IndexPage}/>
                <Route path='/unauthorized' exact={true} component={UnauthorizedPage}/>
                <Route path='/forgotten' exact={true} component={ForgotPasswordPage}/>
                <Route path='/registration' exact={true} component={RegistrationPage}/>
                <Route path='/search' exact={true} component={Search}/>
                <Route path='/profile/:username' exact component={Profile}/>
                <Route path='/info' exact component={ProfileInfo}/>

                <AuthenticatedRoute path='/new_post' exact component={NewPost} isAdminProhibited={true} />
                <AuthenticatedRoute path='/chats' exact component={Chats} isAdminProhibited={true} />
                <AuthenticatedRoute path='/saved' exact component={Saved} isAdminProhibited={true} />
                <AuthenticatedRoute path='/story-archive' exact component={StoryArchive} isAdminProhibited={true} />
                <AuthenticatedRoute path='/notifications' exact component={Notifications} isAdminProhibited={true} />
                <AuthenticatedRoute path='/submit-verification-request' exact component={SubmitVerificationRequest} isAdminProhibited={true} />
                <AuthenticatedRoute path='/view-my-verification-request' exact component={ViewMyVerificationRequests} isAdminProhibited={true} />
                <AdminRoute path='/view-pending-verification-request' exact component={ViewPendingVerificationRequests} />
                <AdminRoute path='/view-all-verification-request' exact component={ViewAllVerificationRequests} />

                <AuthenticatedRoute path='/blocked' exact component={BlockedUsers} />
                <AuthenticatedRoute path='/close_friends' exact component={CloseFriends} />
                <AuthenticatedRoute path='/liked' exact component={Liked} />
                <AuthenticatedRoute path='/disliked' exact component={Disliked} />
                <AuthenticatedRoute path='/archive' exact component={Archived} />
                <AuthenticatedRoute path='/edit_profile' exact component={EditProfile} />
                <AuthenticatedRoute path='/password' exact component={ChangePassword} />
                <AuthenticatedRoute path='/privacy' exact component={EditUserPrivacy} />
                <AuthenticatedRoute path='/edit_photo' exact component={EditProfileImage} />
                <Route path='/agent_registration' exact component={AgentReg} />
                <AuthenticatedRoute path='/agent_check' exact component={AgentCheck} />
                <AuthenticatedRoute path='/complaints' exact component={ComplaintPreview} />
                <AuthenticatedRoute path='/influencers' exact component={Influencers} />
                <AuthenticatedRoute path='/campaign-requests' exact component={CampaignRequests} />
                <AgentRoute path="/campaigns" exact component={CampaignsHome} />
                <AgentRoute path="/campaigns/create" exact component={CreateCampaign} />
                <AgentRoute path="/campaigns/preview/:id" component={CampaignPreview} />
            </Router>
        </div>
    );
}

export default App