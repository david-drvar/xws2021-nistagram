import RootService from './root.service'
import { contentService } from '../../backendPaths';

class HighlightsService extends RootService {
    constructor(){
        super(contentService() + "/highlights")
    }

    async getUserHighlights(data){
        const { userId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get(`/user/${userId}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async createHighlight(data){
        const { userId, name, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('', { userId, name }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async removeHighlight(data){
        const { id, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post(`/${id}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async saveStoryToHighlight(data){
        const { userId, highlightId, storyId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('/create', { userId, highlightId, storyId }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async removeStoryFromHighlight(data){
        const { userId, highlightId, storyId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('/delete', { userId, highlightId, storyId }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getHighlight(data){
        const { id, jwt } = data;
        const headers = this.setupHeaders(jwt);

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

const highlightsService = new HighlightsService();

export default highlightsService;