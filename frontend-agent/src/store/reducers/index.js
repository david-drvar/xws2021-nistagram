import { combineReducers } from 'redux';

import { user } from './user.reducer'
import { apiKey } from './apiKey.reducer'

const rootReducer = combineReducers({
    user, apiKey
})

export default rootReducer;