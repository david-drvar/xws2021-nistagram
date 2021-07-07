import React from 'react';
import { Link } from 'react-router-dom';

function UnauthorizedPage(){
        return (
            <div className='container'>
                <div>
                    <h1>403 - Unauthorized access</h1>
                    <p>Please log in to access this page!</p>
                </div>
                <p><Link to='/'>Back to Login page</Link></p>
            </div>
        );

}export default UnauthorizedPage;
