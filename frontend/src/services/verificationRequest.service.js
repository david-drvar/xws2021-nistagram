import RootService from './root.service'
import { userService as getUserService } from './../backendPaths';

class VerificationRequestService extends RootService {
    constructor(){
        super(getUserService() + "/api/users")
    }

    async getAllVerificationRequests(data){
        const {jwt} = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get('/get-all-verification-requests', { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async submitVerificationRequest(data){
        const {userId, documentPhoto, category, jwt} = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('/submit-verification-request', {
            userId, documentPhoto, category
        },{ headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getVerificationRequestsByUser(data){
        const {userId, jwt} = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('/get-verification-requests-by-user', {
            userId
        },{ headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getPendingVerificationRequests(data){
        const {jwt} = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get('/get-pending-verification-requests',{ headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async changeVerificationRequestStatus(data){
        const {id, status, jwt} = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('/change-verification-request-status',{
            id, status
        },{ headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const verificationRequestService = new VerificationRequestService();

export default verificationRequestService;