import RootService from './root.service'
import { contentService } from '../../backendPaths';

class CommentService extends RootService {
    constructor(){
        super(contentService() + "/comments")
    }


    async addComment(data){
        const { userId, postId, content, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('', { userId, postId, content }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const commentService = new CommentService();

export default commentService;