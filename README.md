# pokeapi
Pokemon Restful API written in native golang. This project implements Domain Driven Design.

## ENDPOINTS

There are 2 endpoints exposed --> GetPokemon and FetchPokemon

1- GetPokemon --> It gets a pokemon by id from the storage. 
It can return a 404 NotFound or 500 InternalServer

2- FetchPokemon --> It gets a pokemon by id from the external api and saves it in the storage. 
It can return a 500 InternalServer, either because of the storage or the external service.

You can fetch a pokemon with the following curl:

`curl --location --request GET 'http://localhost:8080/pokeapi/fetch?id=9'`

You can get a pokemon with the following curl:

`curl --location --request GET 'http://localhost:8080/pokeapi/get?id=9'`

## DEPENDENCIES

### Rest Client
The "https://pokeapi.co/" is used as an external service. This app retrieves pokemon data from that api

### Storage
It is used a mongo database for storing pokemon data by id

## FLAGS
You can run the application using a custom flag for the server port.
For example, you can run the app in port 9090 by using the following command:

`go run ./api/cmd/main.go --port 9090`

By default, it runs on 8080

## DOCKER

For building the image project, use the following command from content root:

`docker build -t [image-name] ./api`

For running an instance, use the following command:

`docker run -d -p 8080:8080 --name [container-name] [image-name]`

Make sure you are running a mongo db instance before running the application:

`docker run -d --name [mongo-container] mongo:tag`

For building both mongodb and the app run a docker-compose command:

`docker-compose up -d`