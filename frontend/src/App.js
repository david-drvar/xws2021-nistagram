import  React from "react";
import {IndexPage} from './pages/IndexPage.js'
import {ForgotPasswordPage} from './pages/forgotPass/ForgotPasswordPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import RegistrationPage from "./pages/RegistrationPage";
import PostsAndStories from "./components/FrontPageComponents/PostsAndStories";
import ProfilePage from "./components/FrontPageComponents/ProfilePage";
import Chats from "./components/FrontPageComponents/Chats";
import Saved from "./components/FrontPageComponents/Saved";
import HomePage from "./components/FrontPageComponents/HomePage";



function App () {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact={true} component={IndexPage}/>
                <Route path='/forgotten' exact={true} component={ForgotPasswordPage}/>
                <Route path='/registration' exact={true} component={RegistrationPage}/>
                <Route path='/home' exact={true} component={HomePage}/>
                <Route path='/posts' exact component={PostsAndStories} />
                <Route path='/profile' exact component={ProfilePage} />
                <Route path='/chats' exact component={Chats} />
                <Route path='/saved' exact component={Saved} />
             </Router>
        </div>
    );
}
export default App