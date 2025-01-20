.PHONY: run
run:
		docker-compose up -d && \
		go run ./cmd/task-server/main.go

.DEFAULT_GOAL: run

.PHONY: cover
cover:
		go test -short -race -coverprofile=coverage.out ./... 
		go tool cover -html=coverage.out
		rm coverage.out

.PHONY: mock
mock:
		mockgen -source=./internal/tasks/handlers/tests/create_task_test.go \
		-destination=./internal/tasks/handlers/tests/mocks/task_creator_mock.go