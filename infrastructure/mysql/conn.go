package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/CA22-game-creators/cookingbomb-apiserver/config"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.DSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
