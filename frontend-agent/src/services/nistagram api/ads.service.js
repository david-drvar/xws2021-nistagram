import RootService from './root.service'
import { contentService } from '../../backendPaths';

class AdsService extends RootService {
    constructor(){
        super(contentService() + "/ads")
    }

    async getAdCategories(data){
        const { jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get(`/categories`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async createAdCategory(data){
        const { jwt, name } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post(`/categories`, { name: name }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const adsService = new AdsService()
export default adsService;