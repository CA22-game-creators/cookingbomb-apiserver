package user

import (
	"gorm.io/gorm"

	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	dbModel "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/mysql/model/user"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Save(entity domain.User) error {
	user := dbModel.New(entity)
	return r.db.Create(&user).Error
}
