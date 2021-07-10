import "../../style/suggestions.css";
import ProfileForSug from "./ProfileForSug";
import React, {useEffect, useState} from "react";
import userService from "../../services/user.service";
import toastService from "../../services/toast.service";
import followersService from "../../services/followers.service";
import {useSelector} from "react-redux";

function Suggestions() {
    const[recommendations,setRecommendations]=useState([])
    const[users,setUsers]=useState([])
    const store = useSelector(state => state);

    useEffect(() => {
        getRecommendations()
    }, []);

    async function getRecommendations() {
        const response = await followersService.getRecommendations(
            {
                id:store.user.id,
                jwt: store.user.jwt
            })

        if (response.status === 200) {
           setRecommendations(response.data.recommendations)
            getUsers(response.data.recommendations)
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }

    function getUsers(requests) {
        requests.map((request, i) => {
            if(users.some(item => item.id === request.userId)) {
                return;
            }
            getUserById(request.userId)
            }
        );
    }

    async function getUserById(request) {
        const response = await userService.getUserById({
            id: request,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {

            setUsers(users=>[...users,response.data])
        } else {
            console.log("getuserbyid error")
        }
    }
    return (
        <div className="suggestions">
            <div className="titleContainer">
                <div className="title">Suggestions For You</div>
            </div>
            {users.map((result,index) =>
               <div><ProfileForSug
                        username={result.username}
                        firstName={result.firstName}
                        lastName={result.lastName}
                        caption="suggestion"
                        urlText="Follow"
                        iconSize="big"
                        captionSize="small"
                        image={result.profilePhoto} storyBorder={false} />
               </div>
            )}
        </div>
    );
}

export default Suggestions;