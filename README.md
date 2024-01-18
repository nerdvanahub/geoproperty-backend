# GeoProperty Backend

## Description

This Service is responsible for managing the properties and the users of the application. It also provides the authentication and authorization for the users.

## Installation

### Requirements

- Golang > 1.18
- Docker
- Docker Compose
- Make
- Grpc
- Protobuf
- MinIO
- Postgres (Postgis)

## Steps to run the project

1. Clone the repository
2. Run `make install` to install the dependencies
3. Run `make run` to run the project or `go run main.go server -i 0.0.0.0 -p 3000` to run the project without docker
4. Run `make build` to build the project or `go build -o bin/server main.go` to build the project without docker
5. If you want to run the project with docker, run `make docker-build` to build the docker image and `make docker-run` to run the docker image
