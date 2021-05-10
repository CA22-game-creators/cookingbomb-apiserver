//go:generate mockgen -source=$GOFILE -destination=../mock/util/$GOFILE
package util

import (
	"github.com/google/uuid"

	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
)

type tokenGenerator struct{}

type TokenGenerator interface {
	Generate() (uuid.UUID, error)
}

func NewTokenGenerator() TokenGenerator {
	return tokenGenerator{}
}

func (tokenGenerator) Generate() (uuid.UUID, error) {
	token, err := uuid.NewRandom()
	return token, errors.Internal(err.Error())
}
