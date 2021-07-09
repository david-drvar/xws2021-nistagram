import Navigation from "../HomePage/Navigation";
import React, {useEffect} from "react";
import {ProSidebar, Menu, MenuItem, SubMenu, SidebarHeader, SidebarContent, SidebarFooter} from 'react-pro-sidebar';
import {
    CgProfile,
    MdBookmarkBorder,
    FaHeart,
    FaHeartBroken, HiOutlinePhotograph,
    IoMdSettings, RiLockPasswordLine, SiGnuprivacyguard,
} from "react-icons/all";
import 'react-pro-sidebar/dist/css/styles.css';
import {Link} from "react-router-dom";
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
                        <MenuItem icon={<MdBookmarkBorder/>} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>My orders<Link to="/my-orders" /> </MenuItem>
                        <MenuItem icon={<IoMdSettings/>}>Edit profile info  <Link to="/edit_profile" /> </MenuItem>
                        <MenuItem icon={<RiLockPasswordLine/>}>Edit password  <Link to="/password" /> </MenuItem>
                        <MenuItem icon={<HiOutlinePhotograph/>}>Edit profile photo  <Link to="/edit_photo" /> </MenuItem>
                        {store.user.role === "Agent" &&
                            <MenuItem icon={<HiOutlinePhotograph/>}>API key <Link to="/api-key"/> </MenuItem>
                        }
                    </Menu>
                </ProSidebar>
            </div>
        </div>
    );
}export default ProfileInfo;