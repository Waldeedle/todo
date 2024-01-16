package todos_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	mock_todos "github.com/waldeedle/todo/internal/mocks/todos"
	"github.com/waldeedle/todo/internal/models"
	"go.uber.org/mock/gomock"
)

//todo: update these tests to go through a handler where the mock is injected

var testTodoID = 1
var testTodoAccountID = 1
var testTodoTitle = "test"
var testTodoIsCompleted = false
var testTodoIsDeleted = false
var testTodoCreatedAt = time.Now()
var testTodoUpdatedAt = time.Now()
var testTodoDeletedAt = time.Now()

var testTodo = &models.Todo{
	ID:          &testTodoID,
	AccountID:   &testTodoAccountID,
	Title:       &testTodoTitle,
	IsCompleted: &testTodoIsCompleted,
	IsDeleted:   &testTodoIsDeleted,
	CreatedAt:   &testTodoCreatedAt,
	UpdatedAt:   &testTodoUpdatedAt,
	DeletedAt:   &testTodoDeletedAt,
}

// use the mock in /internal/mocks/todos/service.go and test all the methods above
func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_todos.NewMockService(ctrl)

	mockService.EXPECT().Create(*testTodo.GetTitle()).Return(testTodo, nil)

	todo, err := mockService.Create(*testTodo.GetTitle())
	assert.Nil(t, err)

	assert.Equal(t, *todo.GetTitle(), *testTodo.GetTitle())
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_todos.NewMockService(ctrl)

	mockService.EXPECT().Get(*testTodo.GetID()).Return(testTodo, nil)

	todo, err := mockService.Get(*testTodo.GetID())
	assert.Nil(t, err)

	assert.Equal(t, *todo.GetTitle(), *testTodo.GetTitle())
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_todos.NewMockService(ctrl)

	mockService.EXPECT().GetAll().Return([]*models.Todo{testTodo}, nil)

	todos, err := mockService.GetAll()
	assert.Nil(t, err)

	assert.Equal(t, *todos[0].GetTitle(), *testTodo.GetTitle())
}

// use table driven test for the title and is completed fields
func TestUpdate(t *testing.T) {
	updatedTodoAccountID := 2
	updatedTodoTitle := "updated"
	updatedTodoIsCompleted := true
	updatedTodoIsDeleted := true
	updatedTodoCreatedAt := time.Now()
	updatedTodoUpdatedAt := time.Now()
	updatedTodoDeletedAt := time.Now()
	testcases := []struct {
		testcase    string
		updatedTodo models.Todo
	}{
		{
			testcase: "valid payload",
			updatedTodo: models.Todo{
				ID:          testTodo.GetID(),
				AccountID:   &updatedTodoAccountID,
				Title:       &updatedTodoTitle,
				IsCompleted: &updatedTodoIsCompleted,
				IsDeleted:   &updatedTodoIsDeleted,
				CreatedAt:   &updatedTodoCreatedAt,
				UpdatedAt:   &updatedTodoUpdatedAt,
				DeletedAt:   &updatedTodoDeletedAt,
			},
		},
		{
			testcase: "invalid payload",
			updatedTodo: models.Todo{
				ID:          testTodo.GetID(),
				AccountID:   &updatedTodoAccountID,
				Title:       nil,
				IsCompleted: &updatedTodoIsCompleted,
				IsDeleted:   &updatedTodoIsDeleted,
				CreatedAt:   &updatedTodoCreatedAt,
				UpdatedAt:   &updatedTodoUpdatedAt,
				DeletedAt:   &updatedTodoDeletedAt,
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.testcase, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mock_todos.NewMockService(ctrl)

			mockService.EXPECT().Update(&tc.updatedTodo).Return(&tc.updatedTodo, nil)

			todo, err := mockService.Update(&tc.updatedTodo)
			assert.Nil(t, err)

			assert.Equal(t, *todo.GetAccountID(), *tc.updatedTodo.GetAccountID())
			assert.Equal(t, *todo.GetTitle(), *tc.updatedTodo.GetTitle())
			assert.Equal(t, *todo.GetIsCompleted(), *tc.updatedTodo.GetIsCompleted())
			assert.Equal(t, *todo.GetIsDeleted(), *tc.updatedTodo.GetIsDeleted())
			assert.Equal(t, *todo.GetCreatedAt(), *tc.updatedTodo.GetCreatedAt())
			assert.Equal(t, *todo.GetUpdatedAt(), *tc.updatedTodo.GetUpdatedAt())
			assert.Equal(t, *todo.GetDeletedAt(), *tc.updatedTodo.GetDeletedAt())
		})
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_todos.NewMockService(ctrl)

	mockService.EXPECT().Delete(*testTodo.GetID()).Return(nil)

	err := mockService.Delete(*testTodo.GetID())
	assert.Nil(t, err)
}
