#!/bin/bash
cd .. && cd gateway && docker build -t gateway .
cd .. && cd frontend && docker build -t frontend .
cd .. && cd frontend-agent && docker build -t frontend-agent .
cd .. && cd backend
docker build -t user_service -f UserServiceDockerfile .
docker build -t content_service -f ContentServiceDockerfile .
docker build -t recommendation_service -f RecommendationServiceDockerfile .
docker build -t agent_application -f AgentApplicationDockerfile .
docker build -t chat_service -f ChatServiceDockerfile .
