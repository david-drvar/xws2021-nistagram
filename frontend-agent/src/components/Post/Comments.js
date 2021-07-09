import React, {useState} from "react";
import axios from "axios";
import {Button} from "react-bootstrap";
import "../../style/comments.css";

function Comments(props) {
    const{post}=props;
    const[comments,setComments]=useState([]);
    const[newComment, setNewComment]=useState('');

    /*treba da dobavimo komentare posta na osnovu id-a
        useEffect(() => {
                getComments();
        });
*/
    function getComments(){
        axios
            .get('http://localhost:8080/api/comments/'+post.id)
            .then(res => {
                console.log("RADI")
                setComments([
                    ...comments,{
                        id:'',
                        userId:'',
                        content:'',
                        createdAt:'',
                    }
                ]);
            }).catch(res => {
            console.log("NE RADI")
        })
    }
    //mokovani podaci za proveru
    const commentsOne = [
        {
            userId: "raffagrassetti",
            content: "Woah dude, this is awesome! 🔥",
            id: 1,
            createdAt:'f'
        },
        {
            userId: "therealadamsavage",
            content: "Like!",
            id: 2,
            createdAt:'f'
        },
        {
            userId: "mapvault",
            content: "Niceeeee!",
            id: 3,
            createdAt:'f'
        },
    ];
    function handleSubmitComment(){
        axios
            .post('http://localhost:8005/api/comments',{
                PostId: post.id,
                UserId: post.userId,
                Content:newComment,
                CreatedAt:new Date()
            })
            .then(res => {
                console.log("RADI")
                //treba da nam vrati username od ussera koji su lajkovali
            }).catch(res => {
            console.log("NE RADI")
        })
    }




    return (
        <div className="comments" >
            <div style={{backgroundColor:'#f5f1f2'}}>
            {commentsOne.map((comment) => {
                return (
                    <div className="commentContainer">
                        <div className="accountName">{comment.userId}</div>
                        <div className="comment">{comment.content}</div>
                    </div>

                );
            })}
            </div>

            <div className="row">
                <div style={{ marginLeft: '3em'}}className="accountName"><strong>IDUSERA</strong></div>
                <input  style={{marginLeft:'2em', height:'1.4em'}} aria-label="Add a comment" autoComplete="off"  type="text"
                        name="add-comment" placeholder="Add a comment..." value={newComment} onChange={({ target }) => setNewComment(target.value)} />
                <Button size="sm" style={{marginLeft:'0.6em',fontSize :'small', width:'9em'}} variant="light" onClick={handleSubmitComment}><strong> Add </strong></Button>
            </div>

        </div>

    );
}export default Comments;