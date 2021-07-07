import RootService from './root.service'
import {agentService} from '../backendPaths';
import axios from "axios";

class AgentService extends RootService {
    constructor(){
        super(agentService() + "/api/agent")
    }

    async login(data){
        const { email, password } = data
        const response = axios.post('http://localhost:8080/api/agent/login', {
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

    async getUserByUsername(data){
        const { username, jwt } = data
        const headers = this.setupHeaders(jwt)
        const response = axios.get('http://localhost:8080/api/agent/getUserByUsername/'+username,{
            headers: headers
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

    async getAllUsers(data) {
        const { jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('',{
            headers
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
        const { id, jwt } = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/username/'+id,{
            headers: headers
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


    async createUser(data){
        const { id, firstName, lastName, email, username, password, role, birthdate, profilePhoto,
        phoneNumber, sex, isActive, address} = data
        const response = axios.post('http://localhost:8080/api/agent/create-user',{
            id, firstName, lastName, email, username, password, role, birthdate, profilePhoto,
            phoneNumber, sex, isActive, address
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

const userService = new AgentService()

export default userService;