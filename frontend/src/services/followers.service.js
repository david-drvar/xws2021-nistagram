import RootService from './root.service'
import { recommendationService } from './../backendPaths';

class FollowersService extends RootService {
    constructor() {
        super(recommendationService() + "/api/followers")
    }

    async getFollowers(data){
        const { userId,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/get_followers',{
            user: { UserId: userId}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getFollowing(data){
        const { userId,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/get_followings',{
            user: { UserId: userId}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async createConnection(data){
        const { userId,followerId,isApprovedRequest,isCloseFriends,isMuted,
            isMessageNotificationEnabled,isPostNotificationEnabled, isStoryNotificationEnabled, isCommentNotificationEnabled,
            requestIsPending,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/create_connection',{
          follower:{  userId,followerId,isApprovedRequest,isCloseFriends,isMuted,
              isMessageNotificationEnabled,isPostNotificationEnabled, isStoryNotificationEnabled, isCommentNotificationEnabled,
              requestIsPending}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async unfollow(data){
        const { userId,followerId,isApprovedRequest,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/delete_directed',{
             userId,followerId,isApprovedRequest
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }
    async updateUserConnection(data){
        const { userId,followerId,isApprovedRequest,isCloseFriends,isMuted,
            isMessageNotificationEnabled,isPostNotificationEnabled, isStoryNotificationEnabled, isCommentNotificationEnabled,
            requestIsPending,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/update_follower',{
             userId,followerId,isApprovedRequest,isCloseFriends,isMuted,isMessageNotificationEnabled,isPostNotificationEnabled, isStoryNotificationEnabled, isCommentNotificationEnabled,requestIsPending
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }


    async getFollowersConnection(data){
        const { userId,followerId,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/connection',{
             userId,followerId
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getCloseFriends(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/close/'+id,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async acceptRequest(data){
        const { userId,followerId,isApprovedRequest,isCloseFriends,isMuted,
            isMessageNotificationEnabled,isPostNotificationEnabled, isStoryNotificationEnabled, isCommentNotificationEnabled,
            requestIsPending,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/accept_request',{
            userId,followerId,isApprovedRequest,isCloseFriends,isMuted, isMessageNotificationEnabled,isPostNotificationEnabled, isStoryNotificationEnabled, isCommentNotificationEnabled,requestIsPending
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getRecommendations(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/recommendation/'+id,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

}const followersService = new FollowersService()

export default followersService;