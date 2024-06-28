BINARY_NAME=/go/bin/wodm8-core-api

compose-build:
	go mod tidy
	docker-compose build

compose-up:
	docker-compose up

build:
	CGO_ENABLED=0 go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm -f ${BINARY_NAME}

test:
	go test ./internal/platform/server/handler/exercise/...

test-coverage:
	go test -coverprofile=coverage.out ./internal/platform/server/handler/exercise/...
	go tool cover -html=coverage.out