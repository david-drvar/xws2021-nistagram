import RootService from './root.service'
import { contentService } from './../backendPaths';

class StoryService extends RootService {
    constructor(){
        super(contentService() + "/stories")
    }

    async getHomepageStories(data){
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

    async getMyStories(data){
        const { userId, jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = this.apiClient.get(`/archive/${userId}`, { headers })
        .then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }

    async getUsersStories(data){
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

    async getStoryById(data){
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


    async createStory(data){
        const { id, userId, isAd, type, description, location, createdAt, isCloseFriends,
            media, comments, likes, dislikes, hashtags, jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = this.apiClient.post('', {
            id, userId, isAd, type, description, location, createdAt, isCloseFriends,
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

    async deleteStory(data){
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

    // Convert story with multiple media to multiple stories with single media, to comply with react-insta-stories
    convertStory(story){
        if(story.media.length === 1) return [{
            ...story,
            orderNum: 1, // Important for our custom slider component 
        }]
        
        const stories = [];
        story.media.forEach(singleMedia => {
            singleMedia.orderNum = 1;   // Important for our custom slider component
            stories.push({
                ...story,
                media: [singleMedia] // Must be an array
            })
        })

        return stories;
    }

    async getStoryById(data){
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
}

const storyService = new StoryService()

export default storyService;