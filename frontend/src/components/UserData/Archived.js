import React, {useEffect, useState} from 'react';
import {useDispatch, useSelector} from "react-redux";

function Archived() {
    const dispatch = useDispatch()
    const store = useSelector(state => state);

    return (
        <div style={{marginTop:'55%'}}>
        </div>
    );
}

export default Archived;