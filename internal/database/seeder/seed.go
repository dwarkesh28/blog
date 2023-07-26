package seeder

import (
	articalModel "blog/internal/modules/article/models2"
	userModel "blog/internal/modules/user/models"
	"blog/pkg/database"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("12345"), 12)

	if err != nil {
		log.Fatal("hash password error")
		return
	}
	user := userModel.User{Name: "Dwarkesh", Email: "123@mail.com", Password: string(hashPassword)}

	db.Create(&user)
	log.Printf("user created successfully with address %s", user.Email)

	for i := 1; i <= 10; i++ {
		article := articalModel.Article{Title: fmt.Sprintf("Random title %d", i), Content: fmt.Sprintf("Random context %d", i), UserID: 1}

		db.Create(&article)
		log.Printf("user created successfully with address %s", article.Title)
	}

	fmt.Println("seeder done")

}
