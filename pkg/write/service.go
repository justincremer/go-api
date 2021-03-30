package writing

type Respository interface {
	CreateUser(User) (string, error)
}

type Service interface {
	CreateUser(User) (string, error)
}

type service struct {
	repository Respository
}

func NewService(repository Respository) *service {
	return &service{repository}
}

func (service *service) CreateUser(user User) (string, error) {
	id, err := service.repository.CreateUser(user)
	if err != nil {
		return "", err
	}

	return id, nil
}
