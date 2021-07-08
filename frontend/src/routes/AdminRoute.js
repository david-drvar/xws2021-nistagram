import React from 'react';
import { useSelector } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';

const AdminRoute = ({ component: Component, ...rest }) => {
    const store = useSelector(state => state);

    return (
        <Route {...rest} render = {props => {
            return store.user.role === "Admin" ? 
                <Component {...props} {...rest} /> :
                <Redirect to={{ pathname: "/" }} /> 
            }}
        />
    );
};

export default AdminRoute;