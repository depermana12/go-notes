package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Notes    []Note `json:"notes" gorm:"foreignKey:AuthorId;constraint:OnDelete:CASCADE"`
}
