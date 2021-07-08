import React from 'react';
import { useSelector } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';

const AuthenticatedRoute = ({ component: Component, isAdminProhibited, ...rest }) => {
    const store = useSelector(state => state);

    return (
        <Route {...rest} render = {props => {
            return store.user.jwt ? // If user is logged in, let him through
                (store.user.role !== "Admin" ? 
                    <Component {...props} {...rest} /> :    // If user is not an admin, let him access
                    ( isAdminProhibited ?
                        <Redirect to={{ pathname: "/" }} /> : // If admin cannot access that route, redirect him to home page 
                        <Component {...props} {...rest} /> 
                    )
                ) :
                <Redirect to={{ pathname: "/" }} />  // If the user is not logged in, redirect him to home page
            }}
        />
    );
};

export default AuthenticatedRoute;