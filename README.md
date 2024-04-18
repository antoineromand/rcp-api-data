# RCP API Data

## Introduction

RCP API Data is a Go project designed to handle data operations for the Recup-Plast platform.

## Getting Started

To run the project, execute the following command:

```
go run cmd/rcp-api-data/main.go
```

## Migrations

To run migrations, set the environment variable `MIGRATION` to `true` before starting the project.

Example:

```
MIGRATION=true go run cmd/rcp-api-data/main.go
```

## Environment Variables

Ensure you have a `.env` file in the root directory of the project with the following variables:

```
PORT=3333
RCP_AUTH_PROTOCOL=http
RCP_AUTH_URL=localhost
RCP_AUTH_PORT=3330
DB_SERVER_USER=<username>
DB_SERVER_PASSWORD=<password>
DB_SERVER_DATABASE=data
DB_SERVER_PORT=7051
DB_SERVER_HOST=localhost
ENVIRONMENT=dev
CORS_ORIGIN=http://localhost:3000,...
MIGRATE=false
PREFIX=/api/recup-plast
RP_BROKERS=localhost:19092
```

Please replace `<username>` and `<password>` with your database credentials.

## Docker Compose

There is a Docker Compose file available for easy deployment. 

To run the project using Docker Compose, execute:

```
docker-compose up
```

This will start the project and its dependencies in Docker containers.
