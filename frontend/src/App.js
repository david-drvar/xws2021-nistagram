import  React from "react";
import {LoginPage} from './pages/LoginPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import RegistrationPage from "./pages/RegistrationPage";
import RegistrationPageHooks from "./pages/RegistrationPageHooks";
import Search from "./pages/Search";


function App () {
    return (
        <div className="App">
            <Router>
                <Route path='/login' exact={true} component={LoginPage}/>
                <Route path='/registration' exact={true} component={RegistrationPage}/>
                <Route path='/registration-hooks' exact={true} component={RegistrationPageHooks}/>
                <Route path='/search' exact={true} component={Search}/>

             </Router>
        </div>
    );
}
export default App