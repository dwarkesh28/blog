package migration

import (
	articleModel "blog/internal/modules/article/models2"
	userModel "blog/internal/modules/user/models"
	"blog/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModel.User{}, &articleModel.Article{})

	if err != nil {
		log.Fatal("migration Failed")
	}

	fmt.Println("migration Done!!")
}
