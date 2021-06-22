import React from "react";
import {Button} from "react-bootstrap";
import {useHistory} from "react-router-dom";
import Menu from "./Menu";
import "../../style/navigation.css";

function Navigation() {
    const history = useHistory()

    function redirectToSearch(){
        history.push({
            pathname: '/search'
        })
    }

    return (
        <div className="navigation">
            <div className="container">
                <font face = "Comic Sans MS" size = "5" style={{marginRight:'5em'}}>Ništagram</font>
                <Button variant="outline-dark" style={{marginRight:'25em'}} onClick={redirectToSearch}>Search...</Button>
            </div>
            <Menu className="menu" />
        </div>
    );

}

export default Navigation;