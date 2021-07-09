import { userConstants } from "../constants";
import jwt_decode from "jwt-decode";


const initialState = {
    jwt : "",
    id : "",
    role: "",
    username: "",
    photo: "",
}

export const apiKey = (state = initialState, action) => {

    switch(action.type) {
        case userConstants.SUBMIT_TOKEN:
            let decoded = jwt_decode(action.data.token);
            // console.log("decoded")
            // console.log(decoded)
            return {
                ...state,
                jwt: action.data.token,
                id: decoded.userId,
                role : action.data.role,
                username : action.data.username,
                photo : action.data.photo
            }
        default: 
            return state;
    }
}