package account

import (
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
)

type OutputData struct {
	Account domain.User
	Err     error
}
