package user

import (
	"errors"
)

type User struct {
	ID              ID
	Name            Name
	HashedAuthToken HashedAuthToken
}

func New(id ID, name Name, hashedAuthToken HashedAuthToken) (User, error) {
	if (id == ID{}) {
		return User{}, errors.New("user id is nil")
	}
	if name == "" {
		return User{}, errors.New("user name is nil")
	}
	if hashedAuthToken == nil {
		return User{}, errors.New("user hashed_auth_token is nil")
	}

	return User{id, name, hashedAuthToken}, nil
}
