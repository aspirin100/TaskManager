.PHONY: build
build:
		go run ./cmd/task-server/main.go

.DEFAULT_GOAL: build