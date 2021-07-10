import "../../style/sidebar.css";
import Sticky from "react-sticky-el";
import Suggestions from "./Suggestions";
import {useDispatch, useSelector} from "react-redux";


function Sidebar() {
    const dispatch = useDispatch()
    const store = useSelector(state => state);
    return (
        <div style={{marginLeft:'0%'}}>
        <Sticky topOffset={-20}>
            <div className="sidebar" style={{display:'block'}}>{/*style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>*/}
                <Suggestions/>

            </div>
        </Sticky>
        </div>
    );
}

export default Sidebar;