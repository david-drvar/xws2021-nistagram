import RootService from './root.service'

class SearchService extends RootService {
    constructor() {
        super(process.env.REACT_APP_CONTENT_SERVICE /*+ "/api/content"*/)
    }

    async searchByTag(data){
        const {text,jwt} = data
        const headers=this.setupHeaders(jwt)
        return this.apiClient.post('/posts-by-hashtag', {
            text
        }, {
            headers: headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
    }
    async searchByLocation(data){
        const {location,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/location',{
            location
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }


}

const searchService = new SearchService()

export default searchService;