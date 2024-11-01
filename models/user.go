package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Email    string `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Notes    []Note `gorm:"foreignKey:AuthorId;constraint:OnDelete:CASCADE" json:"notes"`
}
