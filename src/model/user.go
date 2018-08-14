package model

import "time"

type User struct {
	ID        int64 `gorm:"primary_key,AUTO_INCREMENT"`
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
