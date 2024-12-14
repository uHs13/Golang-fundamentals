package userusecase_test

import (
	userusecase "socialMedia/application/domain/useCase/user"
	userrepository "socialMedia/infra/database/repository/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldProperlyCreateAnUser(t *testing.T) {
	userRepository := userrepository.NewCreateUserRepositoryMemory()
	useCase := userusecase.NewCreateUserUseCase(userRepository)

	input := userusecase.CreateUserUseCaseInput{
		Name:     "John Cena",
		Nickname: "john.cena",
		Email:    "john.cena@email.com",
		Password: "A234#Qwerty",
	}

	output, err := useCase.Execute(input)

	require.Nil(t, err)
	require.NotEmpty(t, output.Id)
	require.Equal(t, output.Name, input.Name)
	require.Equal(t, output.Nickname, input.Nickname)
	require.Equal(t, output.Email, input.Email)
}

func TestShouldThrowAnErrorWhenNameIsEmpty(t *testing.T) {
	userRepository := userrepository.NewCreateUserRepositoryMemory()
	useCase := userusecase.NewCreateUserUseCase(userRepository)

	input := userusecase.CreateUserUseCaseInput{
		Name:     "",
		Nickname: "john.cena",
		Email:    "john.cena@email.com",
		Password: "A234#Qwerty",
	}

	_, err := useCase.Execute(input)

	require.Equal(t, err.Error(), "the name cannot be empty")
}

func TestShouldThrowAnErrorWhenNameIsNotComposedByAtLeastTwoWords(t *testing.T) {
	userRepository := userrepository.NewCreateUserRepositoryMemory()
	useCase := userusecase.NewCreateUserUseCase(userRepository)

	input := userusecase.CreateUserUseCaseInput{
		Name:     "John",
		Nickname: "john.cena",
		Email:    "john.cena@email.com",
		Password: "A234#Qwerty",
	}

	_, err := useCase.Execute(input)

	require.Equal(t, err.Error(), "the name must contain at least two words")
}

func TestShouldThrowAnErrorWhenNameIsTooLong(t *testing.T) {
	userRepository := userrepository.NewCreateUserRepositoryMemory()
	useCase := userusecase.NewCreateUserUseCase(userRepository)

	input := userusecase.CreateUserUseCaseInput{
		Name:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia",
		Nickname: "john.cena",
		Email:    "john.cena@email.com",
		Password: "A234#Qwerty",
	}

	_, err := useCase.Execute(input)

	require.Equal(t, err.Error(), "the name cannot be longer than '255' characters")
}

func TestShouldThrowAnErrorWhenNicknameIsEmpty(t *testing.T) {
	userRepository := userrepository.NewCreateUserRepositoryMemory()
	useCase := userusecase.NewCreateUserUseCase(userRepository)

	input := userusecase.CreateUserUseCaseInput{
		Name:     "John Cena",
		Nickname: "",
		Email:    "john.cena@email.com",
		Password: "A234#Qwerty",
	}

	_, err := useCase.Execute(input)

	require.Equal(t, err.Error(), "the nickname cannot be empty")
}

func TestShouldThrowAnErrorWhenNicknameIsTooLong(t *testing.T) {
	userRepository := userrepository.NewCreateUserRepositoryMemory()
	useCase := userusecase.NewCreateUserUseCase(userRepository)

	input := userusecase.CreateUserUseCaseInput{
		Name:     "John Cena",
		Nickname: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia",
		Email:    "john.cena@email.com",
		Password: "A234#Qwerty",
	}

	_, err := useCase.Execute(input)

	require.Equal(t, err.Error(), "the nickname length must not exceed '255' characters")
}

func TestShouldThrowAnErrorWhenEmailIsInvalid(t *testing.T) {
	userRepository := userrepository.NewCreateUserRepositoryMemory()
	useCase := userusecase.NewCreateUserUseCase(userRepository)

	input := userusecase.CreateUserUseCaseInput{
		Name:     "John Cena",
		Nickname: "john.cena",
		Email:    "john.cena",
		Password: "A234#Qwerty",
	}

	_, err := useCase.Execute(input)

	require.Equal(t, err.Error(), "the provided email 'john.cena' is not valid")
}

func TestShouldThrowAnErrorWhenPasswordIsInvalid(t *testing.T) {
	userRepository := userrepository.NewCreateUserRepositoryMemory()
	useCase := userusecase.NewCreateUserUseCase(userRepository)

	input := userusecase.CreateUserUseCaseInput{
		Name:     "John Cena",
		Nickname: "john.cena",
		Email:    "john.cena@email.com",
		Password: "querty",
	}

	_, err := useCase.Execute(input)

	require.Equal(t, err.Error(), "the password must have at least '8' characters")
}
