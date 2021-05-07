package user

import (
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/oklog/ulid/v2"
)

type User struct {
	ID              string `gorm:"primary_key"`
	Name            string
	HashedAuthToken string
}

func New(entity domain.User) User {
	return User{
		ID:              ulid.ULID(entity.ID).String(),
		Name:            string(entity.Name),
		HashedAuthToken: string(entity.HashedAuthToken),
	}
}
