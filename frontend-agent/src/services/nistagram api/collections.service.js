import RootService from './root.service'
import { contentService } from '../../backendPaths';

class CollectionsService extends RootService {
    constructor(){
        super(contentService() + "/collections")
    }

    async getUserCollections(data){
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

    async getCollection(data){
        const { collectionId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get(`/${collectionId}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async createCollection(data){
        const { name, userId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post("", { name, userId }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async deleteCollection(data){
        const { collectionId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.delete(`/${collectionId}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
    
}

const collectionService = new CollectionsService();

export default collectionService;