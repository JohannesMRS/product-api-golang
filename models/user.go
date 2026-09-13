package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name     string    `gorm:"not null" json:"name"`
	Email    string    `gorm:"unique; not null" json:"email"`
	Password string    `gorm:"not null" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
