package user_test

import (
	"testing"

	"github.com/arnoldtherigan15/Go-REST-API-Documentation/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &user.RepositoryMock{Mock: mock.Mock{}}
var userService = user.NewService(userRepository)

func TestUserService_Create(t *testing.T) {
	u := user.RegisterInput{
		Email:     "test@mail.com",
		Password:  "secret",
		FirstName: "test",
	}
	userData := &user.User{
		Email:     "test@mail.com",
		Password:  "secret",
		FirstName: "test",
	}
	userRepository.Mock.On("Create", userData).Return(userData, nil)
	newUser, err := userService.Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, newUser)
	assert.Equal(t, userData, newUser)
	userRepository.Mock.AssertExpectations(t)
}
