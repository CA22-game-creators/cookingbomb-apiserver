package account

import (
	"github.com/google/uuid"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
)

type OutputData struct {
	Account   user.User
	AuthToken uuid.UUID
	Err       error
}
