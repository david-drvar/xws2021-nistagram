import RootService from './root.service'
import { userService } from '../../backendPaths';

class PrivacyService extends RootService {
    constructor() {
        super(userService() + "/api/privacy")
    }

    async getUserPrivacy(data){
        const { userId, jwt } = data
        const headers = this.setupHeaders(jwt)
        const response = this.apiClient.post('/isProfilePublic', {
            userId
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async updateUserPrivacy(data){
        const { Id,isProfilePublic,isDmPublic,isTagEnabled,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/update',{
            privacy:{Id,isProfilePublic,isDmPublic,isTagEnabled}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async blockUser(data){
        const { UserId,BlockedUserId,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/block_user',{
            UserId,BlockedUserId
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async unBlockUser(data){
        const { UserId,BlockedUserId,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/unblock_user',{
            UserId,BlockedUserId
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getUserPrivacyInfo(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/get_privacy/'+id,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }




}const privacyService = new PrivacyService()

export default privacyService;