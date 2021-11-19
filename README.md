

# Ništagram

Ništagram is a social network modeled by popular Instagram.
It is based on microservices architecture and developed by using the programming language Go.

## Installation

These instructions will get you a copy of the project up and running on your local machine for development purposes.

### Prerequisites (backend)
To run the server-side of Ništagram on your local machine, you will need:
1. [Golang](https://golang.org/doc/install)
2. [gRPC for Golang](https://grpc.io/docs/languages/go/quickstart/)
2. [Postgres](https://www.postgresql.org/download/)
	and the following databases: `xws_users`,`xws_content`,`xws_agent`, `xws_chat`
	
3. [Neo4j](https://neo4j.com/download/) and database `neo4j`
4. [Redis](https://redis.io/download)


### Prerequisites (frontend)
To run the frontend (React/Redux) of Ništagram on your local machine you will need:
1. [Node.js](https://nodejs.org/en/)

Afterward, run these commands in the terminal in folders `frontend` and `frontend-agent` :
```sh
npm install
```

```bash
npm start
```

### Prerequisites (gateway)
To run the gateway of Ništagram on your local machine you will need:
1. [Java](https://java.com/en/download/help/download_options.html)
2. [Spring Boot](https://spring.io/guides/gs/spring-boot/)
