# Transaction Service

## Overview
The Transaction Service is a Go-based application designed to handle account and transaction management. It uses PostgreSQL as its database and Docker for containerization.

## Prerequisites
- Go 1.19+
- Docker
- Docker Compose

## Project Structure
- `deploy/docker-compose.yml`: Docker Compose configuration for the database.
- `Makefile`: Contains commands to build, run, test, and clean the application.
- `pkg/handler/handler.go`: HTTP handlers for account and transaction operations.
- `pkg/handler/handler_test.go`: Unit tests for the HTTP handlers.

## Setup and Usage

### Build the Application
To build the application, run:
```sh
make build
```

### Run the Application
To run the application, execute:
```sh
make run
```

### Test the Application
To test the application, execute:
```sh
make test
```

### Run on local
You can simply
```shell
go run cmd/main.go
```

### Rest all the details are self explanatory in makefile 