import RootService from './root.service'
import { chatService  as getChatService} from '../../backendPaths';

class ChatService extends RootService {
    constructor(){
        super(getChatService())
    }

    async CreateChatRoom(data){
        const { person1, person2, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = await this.apiClient.post(`/room`, {
            person1,person2
        },{ headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response;
    }

    async StartConversation(data){
        const { person1, person2, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = await this.apiClient.post(`/room/conversation`, {
            person1,person2
        },{ headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response;
    }

    async SendMessage(data){
        const { person1, person2, roomId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post(`/ws/` + roomId, {
            person1,person2
        },{ headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async GetMessagesForChatRoom(data){
        const { roomId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = await this.apiClient.get(`/room/` + roomId + `/messages`, {
        },{ headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response;
    }

    async acceptMessageRequest(data) {
        const headers = this.setupHeaders(data.jwt);
        const response = await this.apiClient.post('/request/accept', {
            SenderId : data.SenderId,
            ReceiverId : data.ReceiverId
        }, {headers})
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async declineMessageRequest(data) {
        const headers = this.setupHeaders(data.jwt);
        const response = await this.apiClient.post('/request/decline', {
            SenderId : data.SenderId,
            ReceiverId : data.ReceiverId
        }, {headers})
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const chatService = new ChatService();

export default chatService;