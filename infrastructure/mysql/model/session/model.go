package session

import "time"

type Session struct {
	UserID       string `gorm:"primary_key"`
	SessionToken string
	ExpiredAt    time.Time
}
