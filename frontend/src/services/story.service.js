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
}

const storyService = new StoryService()

export default storyService;