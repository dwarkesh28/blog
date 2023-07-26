package models2

import (
	"blog/internal/modules/user/models"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm: "varchar:191"`
	Content string `gorm: "varchar:191"`
	UserID  uint
	User    models.User
}
