//go:generate mockgen -source=$GOFILE -destination=../mock/util/$GOFILE
package util

import (
	"github.com/google/uuid"
)

type tokenGenerator struct{}

type TokenGenerator interface {
	Generate() (uuid.UUID, error)
}

func NewTokenGenerator() TokenGenerator {
	return tokenGenerator{}
}

func (tokenGenerator) Generate() (uuid.UUID, error) {
	return uuid.NewRandom()
}
