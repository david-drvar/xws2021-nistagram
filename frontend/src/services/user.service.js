import RootService from './root.service'
import { userService as getUserService } from './../backendPaths';

class UserService extends RootService {
    constructor(){
        super(getUserService() + "/api/users")
    }

    async login(data){
        const { email, password } = data
        const response = this.apiClient.post('/login', {
            email,
            password
        }).then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }

    async googleLogin(data){
        const { googleToken } = data
        const response = this.apiClient.post('/auth/google', {
            token: googleToken
        }).then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }

    async checkIsApproved(data){
        const { id, jwt } = data
        const headers = this.setupHeaders(jwt)
        const response = this.apiClient.post('/checkIsApproved', {
            id
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }


    async approveAccount(data){
        const { id, oldPassword, newPassword, repeatedPassword, jwt } = data
        const headers = this.setupHeaders(jwt)
        const response = this.apiClient.post('/approveAccount', {
            password: { id, oldPassword, newPassword, repeatedPassword }
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getUserByUsername(data){
        const { username,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/getUserByUsername/'+username,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getUserById(data) {
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/'+id,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async editProfile(data){
        const { id,firstName,lastName,email,phoneNumber,username,profilePhoto,sex,website,biography,jwt,role} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/update_profile',{
            user:{id,firstName,lastName,email,phoneNumber,username,profilePhoto,sex,website,biography, role}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async changePassword(data){
        console.log(data)
        const {id,oldPassword,newPassword,repeatedPassword,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/update_password',{
            password:{id,oldPassword,newPassword,repeatedPassword}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getUsernameById(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/username/'+id,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async searchByUser(data){
        const { username,firstName,lastName,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/searchByUser',{
            username,firstName,lastName
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async updatePhoto(data){
        const { userId,photo,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/update_photo',{
            userId,photo
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getBlockedUsers(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/get_blocked/'+id,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getUserNotifications(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/get_notifications/'+id,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async deleteNotification(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/delete_notifications/'+id,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async readNotifications(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/read_notifications',{
            id
            },
        {
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async deleteByTypeAndCreator(data){
        const { creatorId,type,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/deleteBy_type_creator',{
                creatorId,type
            },
            {
                headers:headers
            }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

}

const userService = new UserService()

export default userService;