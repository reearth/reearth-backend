# reearth-backend

## Getting Started 

## Installation

To get started, follow these instructions:

 1. Clone to your local machine using `git`.
 1. Make sure that you have Go v1.16 or later. See instructions [here](https://golang.org/doc/install).
 1. Make sure that you have Docker and docker-compose. See instructions [here](https://docs.docker.com/get-docker/) and [here](https://docs.docker.com/compose/install/).
 1. Make sure that you have GolangCI-Lint installed. See instructions [here](https://github.com/golangci/golangci-lint#install).
 
 ## Building and running the app 
 
Reearth is a Go web application built using [Echo](https://github.com/golangci/golangci-lint#install) as a framework and MongoDB as DBMS.
 
To run DB with docker-compose:
 
```make run-db```

This will run the DB. Basically, docker-compose will find any containers from previous runs, or recreate a new container with the same configurations

To run the app: 

```make run-app```

This will start `reearth` app
 
To create a local build:

```make build``` 

This will generate an executable `reearth` which you can run by `./reearth`




 

## License

[Apache License 2.0](LICENSE)
