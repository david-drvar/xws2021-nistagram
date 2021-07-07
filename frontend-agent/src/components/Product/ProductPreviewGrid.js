import React, { useState, useEffect } from 'react';
import { Modal } from 'react-bootstrap';

import "./../../style/PostPreviewGrid.css"
import ProductPreviewThumbnail from "./ProductPreviewThumbnail";
import ProductPreviewModal from "./ProductPreviewModal";
import {useHistory} from "react-router-dom";

const ProductPreviewGrid = (props) => {
    const { posts, shouldReload } = props;

    const history = useHistory()


    const openPost = (post) => {
        history.push({ pathname: '/product/' + post.id })
    }

    return (
        ( posts && 
        <div class="postPreviewGrid__Wrapper">
            { posts.map(post => <ProductPreviewThumbnail post={post} openPost={openPost} /> ) }

        </div>
        )
    )
}

export default ProductPreviewGrid;