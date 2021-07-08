import RootService from './root.service'
import { contentService } from './../backendPaths';


class PostService extends RootService {
    constructor(){
        super(contentService() + "/posts")
    }

    async createPost(data){
        const { id, userId, isAd, type, description, location, createdAt,
            media, comments, likes, dislikes, hashtags, jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = this.apiClient.post('', {
            id, userId, isAd, type, description, location, createdAt,
            media, comments, likes, dislikes, hashtags
        }, {headers})
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getHomepagePosts(data){
        const { jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = this.apiClient.get('', { headers })
        .then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }

    async getPostsForUser(data){
        const { userId, jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = this.apiClient.get(`/user/${userId}`, { headers })
        .then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }
    async getPostById(data){
        const { id, jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = this.apiClient.get(`/${id}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async deletePost(data){
        const { id,jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = this.apiClient.delete(''+id, {headers})
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const postService = new PostService()

export default postService;