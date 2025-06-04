package ports

import "github.com/PhipattanachaiDev/golang_api-migration/internal/domain"

type UserService interface {
	CreateUser(user *domain.User) error
	GetUser(id string) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
}
