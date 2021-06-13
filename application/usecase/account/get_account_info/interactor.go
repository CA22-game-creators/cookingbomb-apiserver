package account

import (
	auth "github.com/CA22-game-creators/cookingbomb-apiserver/application/common/auth/authenticate"
)

type interactor struct {
	auth auth.InputPort
}

func New(a auth.InputPort) InputPort {
	return interactor{
		auth: a,
	}
}

func (i interactor) Handle(input InputData) OutputData {
	authOutput := i.auth.Handle(auth.InputData{SessionToken: input.SessionToken})
	if authOutput.Err != nil {
		return OutputData{Err: authOutput.Err}
	}

	return OutputData{Account: authOutput.Account}
}
