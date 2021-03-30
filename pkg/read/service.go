package reading

type Respository interface {
	ListUsers() ([]string, error)
}

type Service interface {
	ListUsers() ([]string, error)
}

type service struct {
	r Respository
}

func NewService(r Respository) *service {
	return &service{r}
}

func (s *service) ListUsers() ([]string, error) {
	users, err := s.r.ListUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
