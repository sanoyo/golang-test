package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/viniciuswebdev/golang-unit-tests/entity"
)

type UserRepositoryStub struct {
	mock.Mock
}

func (ur *UserRepositoryStub) Add(user entity.User) error {
	// 引数である user を忘れずに渡す
	args := ur.Called(user)
	return args.Error(0)
}

type BadWordsRepositoryStub struct {
	mock.Mock
}

func (br *BadWordsRepositoryStub) FindAll() ([]string, error) {
	args := br.Called()
	return args.Get(0).([]string), args.Error(1)
}

func TestShouldRegister(t *testing.T) {
	user := entity.User{
		Name:        "Yo",
		Email:       "test@test.com",
		Description: "Developver",
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)

	badWordsRepository := &BadWordsRepositoryStub{}
	badWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)

	userService := NewUserService(userRepository, badWordsRepository)
	err := userService.Register(user)

	userRepository.AssertCalled(t, "Add", user)
	assert.Nil(t, err)
}

// errors.New("Bad word found")
func TestShouldNotRegister(t *testing.T) {
	user := entity.User{
		Name:        "Yo",
		Email:       "test@test.com",
		Description: "potato Developver", // badWords である potato を含ませる
	}

	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)

	badWordsRepository := &BadWordsRepositoryStub{}
	badWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)

	userService := NewUserService(userRepository, badWordsRepository)
	err := userService.Register(user)

	userRepository.AssertNotCalled(t, "Add", user)
	assert.Error(t, err)
}
