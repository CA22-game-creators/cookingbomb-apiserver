package account

import (
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
)

type interactor struct {
	factory    domain.Factory
	repository domain.Repository
}

func New(f domain.Factory, r domain.Repository) InputPort {
	return interactor{
		factory:    f,
		repository: r,
	}
}

func (i interactor) Handle(input InputData) OutputData {
	userName, err := domain.NewName(input.Name)
	if err != nil {
		return OutputData{Err: err}
	}

	user, authToken, err := i.factory.Create(userName)
	if err != nil {
		return OutputData{Err: err}
	}

	if err = i.repository.Save(user); err != nil {
		return OutputData{Err: err}
	}

	return OutputData{Account: user, AuthToken: authToken}
}
