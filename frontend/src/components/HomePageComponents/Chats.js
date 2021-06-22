import React, {useEffect} from 'react';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";

function Chats() {
    const dispatch = useDispatch()
    const store = useSelector(state => state);

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation/>
            <h1>Cet</h1>
        </div>
    );
}

export default Chats;