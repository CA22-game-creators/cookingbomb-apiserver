//go:generate mockgen -source=$GOFILE -destination=../mock/util/$GOFILE
package util

import (
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
	"golang.org/x/crypto/bcrypt"
)

type cryptoManager struct{}

type CryptoManager interface {
	Encrypt(string) ([]byte, error)
}

func NewCryptManager() CryptoManager {
	return cryptoManager{}
}

func (cryptoManager) Encrypt(v string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, errors.Internal(err.Error())
	}
	return hash, nil
}
