//go:generate mockgen -destination=../mocks/todos/service.go -package=mock_todos github.com/waldeedle/todo/internal/todos Service

package todos

import (
	"github.com/waldeedle/todo/internal/models"
)

type Service interface {
	Create(title string) (*models.Todo, error)
	Get(id int) (*models.Todo, error)
	GetAll() ([]*models.Todo, error)
	Update(id int, title string, completed bool) (*models.Todo, error)
	Delete(id int) error
}

//todo: use uber mock
//todo: implement dead code checker
//todo: implement posthog later when complexity increases

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(title string) (*models.Todo, error) {
	return s.repository.Create(title)
}

func (s *service) Get(id int) (*models.Todo, error) {
	return s.repository.GetById(id)
}

func (s *service) GetAll() ([]*models.Todo, error) {
	return s.repository.GetAll()
}

func (s *service) Update(id int, title string, completed bool) (*models.Todo, error) {
	return s.repository.Update(id, title, completed)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
