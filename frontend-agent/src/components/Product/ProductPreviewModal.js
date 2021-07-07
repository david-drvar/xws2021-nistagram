import React from 'react';
import { Modal } from 'react-bootstrap';

import "./../../style/PostPreviewModal.css"
import Product from "./Product";

const ProductPreviewModal = (props) => {
    const { postUser, post, showModal, setShowModal, shouldReload } = props;

    return (
        <Modal 
            className="PostPreviewModal__Wrapper"
            contentClassName="content" 
            show={showModal} 
            onHide={() => setShowModal(false)}>
            <Product className="Post" shouldReload={shouldReload} post={post} postUser={postUser} />
        </Modal>
    )
}

export default ProductPreviewModal;