package todos

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mock_todos "github.com/waldeedle/todo/internal/mocks/todos"
	"github.com/waldeedle/todo/internal/models"
	"go.uber.org/mock/gomock"
)

// use the mock in /internal/mocks/todos/service.go and test all the methods above
func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_todos.NewMockService(ctrl)

	todoTitle := "test"
	mockService.EXPECT().Create(todoTitle).Return(&models.Todo{Title: &todoTitle}, nil)

	todo, err := mockService.Create(todoTitle)
	assert.Nil(t, err)

	assert.Equal(t, *todo.GetTitle(), todoTitle)
}
