package user

import (
	"errors"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"

	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	myError "github.com/CA22-game-creators/cookingbomb-apiserver/errors"
	dbModel "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/mysql/model/user"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return repository{
		db: db,
	}
}

func (r repository) Save(entity domain.User) error {
	user := dbModel.New(entity)
	if err := r.db.Create(&user).Error; err != nil {
		return myError.Internal(err.Error())
	}
	return nil
}

func (r repository) Find(userID domain.ID) (domain.User, error) {
	var user dbModel.User
	if err := r.db.First(&user, ulid.ULID(userID).String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, nil
		}
		return domain.User{}, myError.Internal(err.Error())
	}
	return domain.FromRepository(
		user.ID, user.Name, user.HashedAuthToken,
	), nil
}
