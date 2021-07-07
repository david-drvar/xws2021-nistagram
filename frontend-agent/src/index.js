import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import { Provider } from 'react-redux'
import { PersistGate } from 'redux-persist/integration/react'
import { ToastContainer } from 'react-toastify';
import "./style/global.css"
import 'react-toastify/dist/ReactToastify.css';
import 'bootstrap/dist/css/bootstrap.css';
import "react-bootstrap";

import configureStore from './store'
const { store, persistor } = configureStore();
 
// Add ToastContainer
ReactDOM.render(
    <Provider store={store}>
        <PersistGate loading={null} persistor={persistor}>
            <ToastContainer className="toast" toastClassName="toast-wrapper" bodyClassName="toast-body" />
            <App/> 
        </PersistGate>
    </Provider>, 
document.getElementById('root'))
