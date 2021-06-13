package account

import (
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"

	application "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/refresh_session_token"
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	dbModel "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/mysql/model/session"
)

type impl struct {
	db *gorm.DB
}

func New(db *gorm.DB) application.SessionTokenRefresher {
	return impl{
		db: db,
	}
}

func (i impl) Handle(user domain.User, newSessionToken uuid.UUID, expiredAt time.Time) error {
	session := dbModel.Session{
		UserID:       ulid.ULID(user.ID).String(),
		SessionToken: newSessionToken.String(),
		ExpiredAt:    expiredAt,
	}
	return i.db.Save(&session).Error
}
