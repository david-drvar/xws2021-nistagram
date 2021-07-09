import RootService from './root.service'
import { userService as getUserService } from '../../backendPaths';

class APIService extends RootService {
    constructor() {
        super(getUserService() + "/api/apiKey")
    }

    async GetKeyByUserId(data) {
        const {id, jwt} = data
        const headers = this.setupHeaders(jwt)
        const response = await this.apiClient.get('/' + id,  {
            headers: headers
        }).then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }
}

const apiService = new APIService()

export default apiService;

