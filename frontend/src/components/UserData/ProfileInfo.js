import Navigation from "../HomePage/Navigation";
import React, {useEffect} from "react";
import {ProSidebar, Menu, MenuItem, SubMenu, SidebarHeader, SidebarContent, SidebarFooter} from 'react-pro-sidebar';
import {
    CgProfile,
    FaBan,
    FaGem,
    FaHeart,
    FaHeartBroken, HiOutlinePhotograph,
    IoMdSettings, RiLockPasswordLine, SiGnuprivacyguard,
} from "react-icons/all";
import 'react-pro-sidebar/dist/css/styles.css';
import {Link} from "react-router-dom";
import CloseFriends from "./CloseFriends";
import {useDispatch, useSelector} from "react-redux";

function  ProfileInfo(){
    const dispatch = useDispatch()
    const store = useSelector(state => state);

    return(
        <div  style={{display: 'grid', gridTemplateColumns: '200px auto',marginLeft: '5%', marginRight: '20%',marginTop:'4.2%'}}>
            <Navigation/>
            <div style={{marginLeft: '5%', marginRight: '20%',marginTop:'4.2%',display: 'flex'}}>
                <ProSidebar >
                    <Menu iconShape="square">
                        <MenuItem icon={<FaGem/>} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>Close friends  <Link to="/close_friends" /> </MenuItem>
                        <MenuItem icon={<FaBan/>} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>Blocked users <Link to="/blocked" /> </MenuItem>
                        <MenuItem icon={<FaHeart/>} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>Liked posts <Link to="/liked" /> </MenuItem>
                        <MenuItem icon={<FaHeartBroken/>} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>Disliked posts <Link to="/disliked" /> </MenuItem>
                        <MenuItem icon={<IoMdSettings/>}>Edit profile info  <Link to="/edit_profile" /> </MenuItem>
                        <MenuItem icon={<RiLockPasswordLine/>}>Edit password  <Link to="/password" /> </MenuItem>
                        <MenuItem icon={<SiGnuprivacyguard/>} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>Edit privacy  <Link to="/privacy" /> </MenuItem>
                        <MenuItem icon={<HiOutlinePhotograph/>}>Edit profile photo  <Link to="/edit_photo" /> </MenuItem>
                    </Menu>
                </ProSidebar>
            </div>
        </div>
    );
}export default ProfileInfo;