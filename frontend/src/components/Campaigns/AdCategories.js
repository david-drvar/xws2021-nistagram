import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import Navigation from './../HomePage/Navigation'
import Spinner from './../../helpers/spinner'
import { Form, Button } from 'react-bootstrap'
import './../../style/AdCategories.css'

import adsService from './../../services/ads.service'
import toastService from '../../services/toast.service';

const AdCategories = () => {
    const [allCategories, setAllCategories] = useState([])
    const [usersCategories, setUsersCategories] = useState([])
    const [isLoading, setIsLoading] = useState(true)

    const store = useSelector(state => state)

    useEffect(() => {
        adsService.getAdCategories({ jwt: store.user.jwt })
            .then(response => {
                response.data && setAllCategories([...response.data.categories])
            })
            .catch(err => {
                toastService.show("err", "Could not retrieve ad categories.")
            })
            
        adsService.getUsersAdCategories({ jwt: store.user.jwt })
            .then(response => {
                response.data && setUsersCategories([...response.data.categories])
                setIsLoading(false)
            })
            .catch(err => {
                toastService.show("err", "Could not retrieve ad categories.")
            })
    }, [])

    const changeCategories = (category, isChecked) => {
        if(isChecked){
            setUsersCategories([...usersCategories.filter(userCategory => userCategory.id !== category.id)])
        }else{
            setUsersCategories([...usersCategories, category])
        }
    }

    const renderList = () => {
        return allCategories.map(category => {
            const isChecked = usersCategories.some(userCategory => userCategory.id === category.id)
            return (
                <div className="custom-control custom-checkbox category" onClick={() => changeCategories(category, isChecked)}>
                    <input type="checkbox" checked={isChecked} />
                    <label className="categoryName">{ category.name }</label>
                </div>
            )
        })
    }

    const saveChanges = () => {
        adsService.updateUsersAdCategories({ jwt: store.user.jwt, categories: usersCategories })
            .then(response => {
                toastService.show("success", "Successfully updated your ad categories.")
            })
            .catch(err => {
                toastService.show("error", "Could not update your ad categories.")
            })
    }

    return (
        <div>
            <Navigation />
            <div className="AdCategories__Wrapper">
                <div className="title">Your Ad Categories</div>
                { isLoading ? <Spinner type="MutatingDots" /> : renderList() }
                <Button className="submit" onClick={saveChanges}>Save Changes</Button>
            </div>
        </div>
    )
}

export default AdCategories;