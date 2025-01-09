package tasksRepository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"

	"github.com/aspirin100/TaskManager/internal/tasks"
	tasksRepository "github.com/aspirin100/TaskManager/internal/tasks/repository"
	mock_create_test "github.com/aspirin100/TaskManager/internal/tasks/repository/mocks"
)

type taskCreator interface {
	CreateTask(ctx context.Context, params tasks.CreateTaskRequest) (uuid.UUID, error)
}

func TestCreateTaskHandler(t *testing.T) {
	ctrl := gomock.NewController(t)

	cases := []struct {
		name          string
		expectedError error
		request       tasks.CreateTaskRequest
	}{
		{
			name:          "ok case",
			expectedError: nil,
			request: tasks.CreateTaskRequest{
				UserID: uuid.MustParse("e05fa11d-eec3-4fba-b223-d6516800a047"),
			},
		},
		{
			name:          "bad case",
			expectedError: tasksRepository.ErrUserNotFound,
			request: tasks.CreateTaskRequest{
				UserID: uuid.Nil,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			creator := mock_create_test.NewMocktaskCreator(ctrl)

			creator.EXPECT().CreateTask(gomock.Any(), tasks.CreateTaskRequest{
				UserID: uuid.MustParse("e05fa11d-eec3-4fba-b223-d6516800a047"),
			}).Return(uuid.New(), nil).AnyTimes()

			creator.EXPECT().CreateTask(gomock.Any(), tasks.CreateTaskRequest{
				UserID: uuid.Nil,
			}).Return(uuid.Nil, tasksRepository.ErrUserNotFound).AnyTimes()

			_, err := creator.CreateTask(context.Background(), tt.request)
			if err != nil {
				switch {
				case !errors.Is(err, tt.expectedError):
					t.Fail()
				}
			}
		})
	}

}
