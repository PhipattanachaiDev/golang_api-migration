package application

import (
	"github.com/PhipattanachaiDev/golang_api-migration/internal/domain"
	"github.com/PhipattanachaiDev/golang_api-migration/internal/ports"
	"time"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(r ports.UserRepository) ports.UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(user *domain.User) error {
	user.CreatedAt = time.Now()
	return s.repo.Create(user)
}
func (s *userService) GetUser(id string) (*domain.User, error) {
	return s.repo.GetByID(id)
}
func (s *userService) GetUsers() ([]*domain.User, error) {
	return s.repo.GetAll()
}
func (s *userService) UpdateUser(user *domain.User) error {
	return s.repo.Update(user)
}
func (s *userService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
