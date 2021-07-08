import "./../../style/PostHeader.css";
import ProfileIcon from "../ProfileComponent/ProfileIcon";
import { NavLink } from 'react-router-dom'

const PostHeader = (props) => {
    const {
      username,
      caption,
      captionSize,
      iconSize,
      hideUsername,
      image,
      isAd
    } = props;
    
    return (
      <div className="post-header">
        <ProfileIcon
          iconSize={iconSize}
          //storyBorder={storyBorder}
          image={image}
        />
        {(username || caption) && !hideUsername && (
          <div className="textContainer">
            {/* Make username clickable */}
            <span> <NavLink className="username" to={{pathname: `/profile/${username}`,}}>{username}</NavLink> {isAd ? "· Sponsored" : ""} </span>
            <span className={`caption ${captionSize}`}>{caption}</span>
          </div>
        )}
      </div>
    );
  }
  
  export default PostHeader;