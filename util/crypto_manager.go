//go:generate mockgen -source=$GOFILE -destination=../mock/util/$GOFILE
package util

import (
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
	return bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
}
