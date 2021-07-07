import { userConstants } from "../constants";

const initialState = {
    user: {},   // Remove?
    jwt: "",
    id: "",
    role: "",
    username: "",
    isSSO: false,
    photo: "",
}

export const user = (state = initialState, action) => {
    switch(action.type) {
        case userConstants.REGISTER_REQUEST:
            return {
                ...state,
                jwt: action.data.jwt,
                id: action.data.id,
                role: action.data.role,
                isSSO: action.data.isSSO
            }
        case userConstants.LOGIN_REQUEST:
            return {
                ...state,
                jwt: action.data.jwt,
                id: action.data.id,
                role: action.data.role,
                isSSO: action.data.isSSO,
                username: action.data.username,
                photo: action.data.photo,
        }
        case userConstants.LOGOUT_REQUEST:
            return {
                ...initialState
            };
        default: 
            return state;
    }
}