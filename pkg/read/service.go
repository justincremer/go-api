package reading

type Respository interface {
	ListUsers() ([]string, error)
}

type Service interface {
	ListUsers() ([]string, error)
}

type service struct {
	repository Respository
}

func NewService(repository Respository) *service {
	return &service{repository}
}

func (service *service) ListUsers() ([]string, error) {
	users, err := service.repository.ListUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
