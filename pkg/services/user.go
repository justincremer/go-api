package services

import "github.com/justincremer/go-api/pkg/schemas"

type UserRespository interface {
	CreateUser(schemas.User) (string, error)
	ListUsers() ([]string, error)
}

type UserService interface {
	CreateUser(schemas.User) (string, error)
	ListUsers() ([]string, error)
}

type service struct {
	repository UserRespository
}

func NewUserService(repository UserRespository) *service {
	return &service{repository}
}

func (service *service) CreateUser(user schemas.User) (string, error) {
	id, err := service.repository.CreateUser(user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (service *service) ListUsers() ([]string, error) {
	users, err := service.repository.ListUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
