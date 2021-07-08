import React from 'react';
import { useSelector } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';

const AgentRoute = ({ component: Component, ...rest }) => {
    const store = useSelector(state => state);

    return (
        <Route {...rest} render = {props => {
            return store.user.role === "Agent" ? 
                <Component {...props} {...rest} /> :
                <Redirect to={{ pathname: "/" }} /> 
            }}
        />
    );
};

export default AgentRoute;