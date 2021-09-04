package model

import (
	"github.com/Pauloo27/archvium/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"index;unique;not null"`
	Email    string `gorm:"index;unique;not null"`
	Password string `gorm:"not null"`
	IsAdmin  bool   `gorm:"not null"`
}

func (u *User) BeforeSave(db *gorm.DB) error {
	u.Password = utils.HashPassword(u.Password)
	return nil
}

func (u *User) ToDto() fiber.Map {
	return fiber.Map{
		"id": u.ID,
		"username": u.Username,
		"createdAt": u.CreatedAt,
		"deletedAt": u.DeletedAt,
		"updatedAt": u.UpdatedAt,
	}
}
