package user

import (
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (r *RepositoryMock) Create(u *User) (*User, error) {
	args := r.Mock.Called(u)

	if args.Get(0) == nil {
		return &User{}, nil
	}
	return args.Get(0).(*User), nil
}
