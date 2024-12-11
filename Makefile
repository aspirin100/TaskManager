.PHONY: build
build:
		go run ./cmd/task-server/task-server.go

.DEFAULT_GOAL: build