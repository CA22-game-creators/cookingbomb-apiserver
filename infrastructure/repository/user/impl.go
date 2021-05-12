package user

import (
	"gorm.io/gorm"

	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
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
		return errors.Internal(err.Error())
	}
	return nil
}
