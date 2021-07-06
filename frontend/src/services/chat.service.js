import RootService from './root.service'
import { chatService  as getChatService} from '../backendPaths';

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

}

const chatService = new ChatService();

export default chatService;