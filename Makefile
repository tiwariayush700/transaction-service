# Variables
APP_NAME = transaction-service
DOCKER_IMAGE = $(APP_NAME):latest
DB_CONFIG = config/config.go

# Go commands
build:
	go mod vendor
	go build -o $(APP_NAME) cmd/main.go

run: build
	./$(APP_NAME)

test:
	go test ./...

clean:
	rm -f $(APP_NAME)

# Docker commands
docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)

# Database setup
db-setup:
	@echo "Configuring database..."
	@echo "Please ensure the database configuration is set in $(DB_CONFIG)"

# Phony targets
.PHONY: build run test clean docker-build docker-run db-setup