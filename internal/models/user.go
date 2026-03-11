package models

import "time"

type User struct {
	ID           uint64
	Username     string    `gorm:"size:50;not null;unique"`
	Password     string    `gorm:"size:300;not null"`
	Salt         string    `gorm:"size:300;not null"`
	Role         int       `gorm:"default:0"`
	Karma        int       `gorm:"default:0"`
	RegisteredAt time.Time `gorm:"autoCreateTime"`
}
