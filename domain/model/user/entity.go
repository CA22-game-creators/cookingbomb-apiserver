package user

import (
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
)

type User struct {
	ID              ID
	Name            Name
	HashedAuthToken HashedAuthToken
}

func New(id ID, name Name, hashedAuthToken HashedAuthToken) (User, error) {
	if (id == ID{}) {
		return User{}, errors.Internal("user id is nil")
	}
	if name == "" {
		return User{}, errors.Internal("user name is nil")
	}
	if hashedAuthToken == nil {
		return User{}, errors.Internal("user hashed_auth_token is nil")
	}

	return User{id, name, hashedAuthToken}, nil
}
