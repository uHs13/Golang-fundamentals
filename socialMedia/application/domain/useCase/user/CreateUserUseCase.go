package userusecase

import (
	port "socialMedia/application/port/repository/user"
	factory "socialMedia/infra/factory/user"
)

type CreateUserUseCaseInput struct {
	Name     string
	Nickname string
	Email    string
	Password string
}

type CreateUserUseCaseOutput struct {
	Id       string
	Name     string
	Nickname string
	Email    string
}

type CreateUserUseCase struct {
	userRepository port.CreateUserRepositoryInterface
}

func NewCreateUserUseCase(userRepository port.CreateUserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository}
}

func (useCase *CreateUserUseCase) Execute(input CreateUserUseCaseInput) (*CreateUserUseCaseOutput, error) {
	user, err := factory.NewUserFactory().MakeComplete(
		input.Name,
		input.Nickname,
		input.Email,
		input.Password,
	)

	if err != nil {
		return nil, err
	}

	createdUser, err := useCase.userRepository.Execute(user)

	if err != nil {
		return nil, err
	}

	output := &CreateUserUseCaseOutput{
		Id:       createdUser.GetId(),
		Name:     createdUser.GetName(),
		Nickname: createdUser.GetNickname(),
		Email:    createdUser.GetEmail(),
	}

	return output, nil
}
