import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import Navigation from './../HomePage/Navigation'
import Spinner from './../../helpers/spinner'
import { Form, Button } from 'react-bootstrap'
import './../../style/AdCategories.css'

import adsService from './../../services/nistagram api/ads.service'
import toastService from '../../services/toast.service';
import ProfileInfo from "../UserData/ProfileInfo";

const AdCategories = () => {
    const [allCategories, setAllCategories] = useState([])
    const [usersCategories, setUsersCategories] = useState([])
    const [isLoading, setIsLoading] = useState(true)

    const store = useSelector(state => state)

    useEffect(() => {
        adsService.getAdCategories({ jwt: store.apiKey.jwt })
            .then(response => {
                response.data && setAllCategories([...response.data.categories])
            })
            .catch(err => {
                toastService.show("err", "Could not retrieve ad categories.")
            })
            
        adsService.getUsersAdCategories({ jwt: store.apiKey.jwt })
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
        if(usersCategories.length < 2) {
            toastService.show("error", "You cannot have less that 2 ad categories selected")
            return
        }

        adsService.updateUsersAdCategories({ jwt: store.apiKey.jwt, categories: usersCategories })
            .then(response => {
                toastService.show("success", "Successfully updated your ad categories.")
            })
            .catch(err => {
                toastService.show("error", "Could not update your ad categories.")
            })
    }

    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div className="AdCategories__Wrapper">
                <div className="title">Your Ad Categories</div>
                { isLoading ? <Spinner type="MutatingDots" /> : renderList() }
                <Button className="submit" onClick={saveChanges}>Save Changes</Button>
            </div>
        </div>
    )
}

export default AdCategories;