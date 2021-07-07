import { createStore, compose, applyMiddleware } from 'redux';
import thunkMiddleware from 'redux-thunk'
import { persistStore, persistReducer } from 'redux-persist'
import localForage from 'localforage'
import rootReducer from './reducers'

const middlewares = [thunkMiddleware];

const composeEnhancer = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

// Persist and rehydrate a redux store
const storage = localForage.createInstance({
    name: "xws-localForage"
})
const persistedReducer = persistReducer({ key: 'root', storage }, rootReducer)

export default () => {
    const store = createStore(
        persistedReducer,
        composeEnhancer( applyMiddleware(...middlewares) )
    );
    const persistor = persistStore(store);

    return { store, persistor }
}