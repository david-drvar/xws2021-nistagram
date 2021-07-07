module.exports.userService = () => {
    return process.env.REACT_APP_USER_SERVICE ? process.env.REACT_APP_USER_SERVICE : "http://localhost:8001"
}

module.exports.contentService = () => {
    return process.env.REACT_APP_CONTENT_SERVICE ? process.env.REACT_APP_CONTENT_SERVICE : "http://localhost:8002"
}

module.exports.chatService = () => {
    return process.env.REACT_APP_CHAT_SERVICE ? process.env.REACT_APP_CHAT_SERVICE : "http://localhost:8003"
}

module.exports.agentService = () => {
    return process.env.REACT_APP_AGENT_SERVICE ? process.env.REACT_APP_AGENT_SERVICE : "http://localhost:8004"
}

module.exports.recommendationService = () => {
    return process.env.REACT_APP_RECOMMENDATION_SERVICE ? process.env.REACT_APP_RECOMMENDATION_SERVICE : "http://localhost:8005"
}