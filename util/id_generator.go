//go:generate mockgen -source=$GOFILE -destination=../mock/util/$GOFILE
package util

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
)

type idGenerator struct{}

type IDGenerator interface {
	Generate() (ulid.ULID, error)
}

func NewIDGenerator() IDGenerator {
	return idGenerator{}
}

func (idGenerator) Generate() (ulid.ULID, error) {
	id, err := ulid.New(ulid.Timestamp(time.Unix(1000000, 0)), rand.Reader)
	if err != nil {
		return ulid.ULID{}, errors.Internal(err.Error())
	}
	return id, nil
}
