import "../../style/profileSug.css";
import ProfileIcon from "../ProfileComponent/ProfileIcon";
import {Button, NavLink} from "react-bootstrap";
import {useHistory} from "react-router-dom";

function ProfileForAutocomplete(props) {
    const {user, username, caption,  urlText, iconSize,captionSize, storyBorder,hideAccountName,image, firstName, lastName} = props;
    const history = useHistory()

    function redirect(){
        history.push({
            pathname: '/profile',
            state: { user:user, follow:true }
        })
    }
    return (
        <div className="profile">
            <ProfileIcon iconSize={iconSize} storyBorder={storyBorder} image={image} />
            {('maja' || caption) && !hideAccountName && (
                <div className="textContainer">
                    <span className="accountName">{firstName} {lastName} @{username}</span>
                    <span className={`caption ${captionSize}`}>{caption}</span>
                </div>
            )}
        </div>
    );
}

export default ProfileForAutocomplete;