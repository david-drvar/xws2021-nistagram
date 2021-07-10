import  React from "react";
import IndexPage from './pages/IndexPage.js'
import {ForgotPasswordPage} from './components/forgotPass/ForgotPasswordPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import RegistrationPage from "./pages/RegistrationPage";
import Saved from "./components/HomePageComponents/Saved";
import Profile from "./components/ProfileComponent/Profile";
import NewProduct from "./components/Product/NewProduct";
import Home from "./components/HomePage/Home";
import UnauthorizedPage from "./helpers/UnauthorizedPage";

import ProfileInfo from "./components/UserData/ProfileInfo";
import APIKey from "./components/UserData/APIKey";
import EditProfile from "./components/UserData/EditProfile";
import ChangePassword from "./components/UserData/ChangePassword";

import EditProfileImage from "./components/UserData/EditProfileImage";
import Notifications from "./components/Notifications/Notifications";

import AuthenticatedRoute from './AuthenticatedRoute';
import AdminRoute from './AdminRoute';
import Product from "./components/Product/Product";
import Orders from "./components/UserData/Orders";
import CampaignsHome from "./components/Campaigns/CampaignsHome";
import CreateCampaign from "./components/Campaigns/CreateCampaign";
import CampaignPreview from "./components/Campaigns/CampaignPreview";
import CampaignRequests from "./components/AgentComponents/CampaignRequests";
import Influencers from "./components/AgentComponents/Influencers";

const App = () => {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact  component={Home}/>
                <Route path='/login' exact={true} component={IndexPage}/>
                <Route path='/unauthorized' exact={true} component={UnauthorizedPage}/>
                <Route path='/forgotten' exact={true} component={ForgotPasswordPage}/>
                <Route path='/registration' exact={true} component={RegistrationPage}/>
                <Route path='/info' exact component={ProfileInfo}/>
                <Route path='/api-key' exact component={APIKey}/>
                <Route path='/newproduct' exact component={NewProduct}/>
                <Route path='/profile/:username' exact component={Profile}/>
                <Route path='/product/:id' exact component={Product}/>
                <Route path='/my-orders' exact component={Orders}/>

                <Route path="/campaigns" exact component={CampaignsHome} />
                <Route path="/campaigns/create" exact component={CreateCampaign} />
                <Route path="/campaigns/preview/:id" component={CampaignPreview} />
                <Route path='/influencers' exact component={Influencers} />
                <Route path='/campaign-requests' exact component={CampaignRequests} />

                <AuthenticatedRoute path='/edit_profile' exact component={EditProfile} />
                <AuthenticatedRoute path='/password' exact component={ChangePassword} />
                <AuthenticatedRoute path='/edit_photo' exact component={EditProfileImage} />
            </Router>
        </div>
    );
}

export default App