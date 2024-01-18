# Geoproperty GPT

## Description

This service is responsible for generating queries from prompt to SQL queries with LLM (Gemini-Pro) and return the results to the client (geoproperty-be) for execute SQL query command.

## Installation

### Requirements

- Python 3.11
- Docker
- Docker Compose
- Make

### Steps

1. Clone this repository
2. Run `make install` to install dependencies
3. Run `make run` to start the service in development mode (with hot reload) and withouth docker
4. Run `make docker-build` to build the docker image and `make docker-run` to start the service with docker
