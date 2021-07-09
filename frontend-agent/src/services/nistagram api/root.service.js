import axios from 'axios';

class RootService {
    constructor(baseURL){
        this.apiClient = axios.create({
            baseURL: baseURL
        })
    }

    setupHeaders(jwt){
        const headers = { Accept: 'application/json', }
        if(jwt) headers["Authorization"] = "Bearer " + jwt 

        return headers
    }
}

export default RootService;