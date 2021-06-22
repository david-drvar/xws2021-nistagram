import "../../style/profileSug.css";
import ProfileIcon from "../ProfileComponent/ProfileIcon";
import { useHistory } from "react-router-dom";

function ProfileForSug(props) {
    const { username, caption, iconSize,captionSize, storyBorder, hideAccountName, image, firstName, lastName} = props;
    const history = useHistory()

    function redirect(){
        history.push({
            pathname: '/profile/'+username,
        })
    }
    return (
        <div className="profile">
            <ProfileIcon iconSize={iconSize} storyBorder={storyBorder} image={image} />
            {(username || caption) && !hideAccountName && (
                <div className="textContainer">
                    <div className="accountName" onClick={redirect}>{firstName} {lastName} <span style={{color: "#8e8e8e"}}>@{username}</span></div>
                    <div className={`caption ${captionSize}`}>{caption.length > 52 ? caption.substring(0, 52).trim() + "..." : caption}</div>
                </div>
            )}
        </div>
    );
}

export default ProfileForSug;