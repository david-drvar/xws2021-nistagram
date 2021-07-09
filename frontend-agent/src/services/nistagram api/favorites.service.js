import RootService from './root.service'
import { contentService } from '../../backendPaths';
class FavoritesService extends RootService {
    constructor(){
        super(contentService() + "/favorites")
    }

    async createFavorite(data){
        const { userId, postId, collectionId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post("/create", { userId, postId, collectionId }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async removeFavorite(data){
        const { userId, postId, collectionId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post("/remove", { userId, postId, collectionId }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getUserFavorites(data){
        const { userId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get(`/${userId}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
    
    async getUserFavoritesOptimized(data){
        const { userId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get(`/optimized/${userId}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const favoritesService = new FavoritesService();

export default favoritesService;