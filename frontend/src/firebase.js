import firebase from "firebase/app";
import "firebase/storage";

// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyA4_Unc7OGSg3Ym8q2Jyu9K25r86HeINzM",
    authDomain: "nistagram-97e89.firebaseapp.com",
    projectId: "nistagram-97e89",
    storageBucket: "nistagram-97e89.appspot.com",
    messagingSenderId: "847846922063",
    appId: "1:847846922063:web:44409e2488e2b46114173a",
    measurementId: "G-MT6M0PF345"
};

firebase.initializeApp(firebaseConfig);

var storage = firebase.storage();

export {
    storage, firebase as default
}