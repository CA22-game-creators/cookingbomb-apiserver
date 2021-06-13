package account

import (
	"time"

	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
	"github.com/CA22-game-creators/cookingbomb-apiserver/util"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

type interactor struct {
	repository            domain.Repository
	sessionTokenRefresher SessionTokenRefresher
	tokenGenerator        util.TokenGenerator
}

func New(r domain.Repository, s SessionTokenRefresher, t util.TokenGenerator) InputPort {
	return interactor{
		repository:            r,
		sessionTokenRefresher: s,
		tokenGenerator:        t,
	}
}

func (i interactor) Handle(input InputData) OutputData {
	// ID から User 取得
	ulid, err := ulid.Parse(input.UserID)
	if err != nil {
		return OutputData{Err: errors.InvalidArgument(err.Error())}
	}
	userID, err := domain.NewID(ulid)
	if err != nil {
		return OutputData{Err: err}
	}
	user, err := i.repository.Find(userID)
	if err != nil {
		return OutputData{Err: err}
	}
	if user.ID == (domain.ID{}) {
		return OutputData{Err: errors.InvalidArgument("user not found")}
	}

	// AuthToken の検証
	if err = bcrypt.CompareHashAndPassword(user.HashedAuthToken, []byte(input.AuthToken)); err != nil {
		return OutputData{Err: errors.Unauthenticated(err.Error())}
	}

	// SessionToken を更新・取得
	newSessionToken, err := i.tokenGenerator.Generate()
	if err != nil {
		return OutputData{Err: err}
	}
	newExpiredAt := time.Now().Add(20 * time.Hour)
	if err := i.sessionTokenRefresher.Handle(user, newSessionToken, newExpiredAt); err != nil {
		return OutputData{Err: err}
	}

	return OutputData{SessionToken: newSessionToken}
}
