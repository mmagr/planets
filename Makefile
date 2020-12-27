
BINARY_FILE=service
BINARY_FILE_DIR=build/bin
DOCKER_IMAGE_TAG=mmagr/planets

.PHONY: build-api build-cli
build-api: tests
	go build -o $(BINARY_FILE_DIR)/$(BINARY_FILE) cmd/weather/api.go

build-cli: tests
	go build -o $(BINARY_FILE_DIR)/$(BINARY_FILE) cmd/weather/cli.go

.PHONY: build-docker-jenkins
build-docker-jenkins:
	docker build -f build/package/docker/Dockerfile -t $(DOCKER_IMAGE_TAG) .

.PHONY: clean
clean:
	rm -rf $(BINARY_FILE_DIR)
	go clean

.PHONY: run-api
run-api:
	go run cmd/weather/api.go

.PHONY: tests
tests:
	go test -v ./...
