package user

type service struct {
	repository Repository
}

type Service interface {
	Create(input RegisterInput) (*User, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input RegisterInput) (*User, error) {
	user := &User{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	if err := user.BeforeSave(); err != nil {
		return &User{}, err
	}
	user, err := s.repository.Create(user)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
