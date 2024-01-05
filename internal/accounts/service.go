package accounts

import (
	"github.com/waldeedle/todo/internal/models"
)

type Service interface {
	Create(email string) (*models.Account, error)
	GetById(id int) (*models.Account, error)
	GetByEmail(email string) (*models.Account, error)
	Update(id int, email string) (*models.Account, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(email string) (*models.Account, error) {
	return s.repository.Create(email)
}

func (s *service) GetById(id int) (*models.Account, error) {
	return s.repository.GetById(id)
}

func (s *service) GetByEmail(email string) (*models.Account, error) {
	return s.repository.GetByEmail(email)
}

func (s *service) Update(id int, email string) (*models.Account, error) {
	return s.repository.Update(id, email)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
