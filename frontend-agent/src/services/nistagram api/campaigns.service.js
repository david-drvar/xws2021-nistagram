import RootService from './root.service'
import { contentService } from '../../backendPaths';

class CampaignsService extends RootService {
    constructor(){
        super(contentService() + "/campaigns")
    }

    async getAgentsCampaigns(data){
        const { jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get(``, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getCampaignById(data){
        const { jwt, id } = data;
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

    async updateCampaign(data){
        const { jwt, campaign } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.put(``, campaign, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async createCampaign(data){
        const { jwt, campaign } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post(``, campaign, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async deleteCampaign(data){
        const { jwt, id } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.delete(`/${id}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const campaignsService = new CampaignsService()
export default campaignsService;