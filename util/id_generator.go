//go:generate mockgen -source=$GOFILE -destination=../mock/util/$GOFILE
package util

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type idGenerator struct{}

type IDGenerator interface {
	Generate() (ulid.ULID, error)
}

func NewIDGenerator() IDGenerator {
	return idGenerator{}
}

func (idGenerator) Generate() (ulid.ULID, error) {
	t := time.Unix(1000000, 0)
	return ulid.New(ulid.Timestamp(t), rand.Reader)
}
