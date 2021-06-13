package auth

import (
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"

	auth "github.com/CA22-game-creators/cookingbomb-apiserver/application/common/auth/authenticate"
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	myError "github.com/CA22-game-creators/cookingbomb-apiserver/errors"
	dbModel "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/mysql/model/session"
)

type impl struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.AccountIDFetcher {
	return impl{
		db: db,
	}
}

func (i impl) Handle(sessionToken string) (domain.ID, error) {
	var session dbModel.Session
	if err := i.db.Where("session_token = ?", sessionToken).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ID{}, nil
		}
		return domain.ID{}, myError.Internal(err.Error())
	}

	if session.ExpiredAt.Before(time.Now()) {
		return domain.ID{}, myError.Unauthenticated("session timeout")
	}

	return domain.ID(ulid.MustParse(session.UserID)), nil
}
