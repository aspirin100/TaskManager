.PHONY: build
build:
		go run ./cmd/task-server/main.go

.DEFAULT_GOAL: build

.PHONY: cover
cover:
		go test -short -race -coverprofile=coverage.out ./... 
		go tool cover -html=coverage.out
		rm coverage.out