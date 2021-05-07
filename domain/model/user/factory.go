//go:generate mockgen -source=$GOFILE -destination=../../../mock/domain/model/user/$GOFILE
package user

import (
	"github.com/google/uuid"

	"github.com/CA22-game-creators/cookingbomb-apiserver/util"
)

type factory struct {
	idGenerator    util.IDGenerator
	tokenGenerator util.TokenGenerator
	cryptoManager  util.CryptoManager
}

type Factory interface {
	Create(name Name) (user User, plainAuthToken uuid.UUID, err error)
}

func NewFactory(i util.IDGenerator, t util.TokenGenerator, c util.CryptoManager) Factory {
	return factory{
		idGenerator:    i,
		tokenGenerator: t,
		cryptoManager:  c,
	}
}

func (f factory) Create(userName Name) (User, uuid.UUID, error) {
	id, err := f.idGenerator.Generate()
	if err != nil {
		return User{}, uuid.Nil, err
	}
	userID, err := NewID(id)
	if err != nil {
		return User{}, uuid.Nil, err
	}

	authToken, err := f.tokenGenerator.Generate()
	if err != nil {
		return User{}, uuid.Nil, err
	}
	hashedAuthToken, err := f.cryptoManager.Encrypt(authToken.String())
	if err != nil {
		return User{}, uuid.Nil, err
	}
	userHashedAuthToken, err := NewHashedAuthToken(hashedAuthToken)
	if err != nil {
		return User{}, uuid.Nil, err
	}

	user, err := New(userID, userName, userHashedAuthToken)
	if err != nil {
		return User{}, uuid.Nil, err
	}

	return user, authToken, nil
}
