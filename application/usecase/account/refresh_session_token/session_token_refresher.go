//go:generate mockgen -source=$GOFILE -destination=../../../../mock/application/usecase/account/refresh_session_token/$GOFILE
package account

import (
	"time"

	"github.com/google/uuid"

	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
)

type SessionTokenRefresher interface {
	Handle(user domain.User, newSessionToken uuid.UUID, expiredAt time.Time) error
}
