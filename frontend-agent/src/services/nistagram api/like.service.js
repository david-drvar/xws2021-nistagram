import RootService from './root.service'
import { contentService } from '../../backendPaths';

class LikeService extends RootService {
    constructor(){
        super(contentService())
    }

    async addLike(data){
        const { userId, postId, isLike, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('/likes', { userId, postId, isLike }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getUserLikedOrDislikedPosts(data){
        const { userId, isLike, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('/liked-posts-user', { userId, isLike }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const likeService = new LikeService();

export default likeService;