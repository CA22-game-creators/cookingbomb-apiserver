package auth

import (
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
)

type interactor struct {
	repository       domain.Repository
	accountIDFetcher AccountIDFetcher
}

func New(r domain.Repository, a AccountIDFetcher) InputPort {
	return interactor{
		repository:       r,
		accountIDFetcher: a,
	}
}

func (i interactor) Handle(input InputData) OutputData {
	userID, err := i.accountIDFetcher.Handle(input.SessionToken)
	if err != nil {
		return OutputData{Err: err}
	}
	if userID == (domain.ID{}) {
		return OutputData{Err: errors.Unauthenticated("session not found")}
	}

	user, err := i.repository.Find(userID)
	if err != nil {
		return OutputData{Err: err}
	}
	if user.ID == (domain.ID{}) {
		return OutputData{Err: errors.Internal("user not found")}
	}

	return OutputData{Account: user}
}
