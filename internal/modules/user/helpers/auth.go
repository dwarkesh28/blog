package helpers

import (
	UserRepository "blog/internal/modules/user/repositiories"
	"blog/internal/modules/user/responses"
	"blog/pkg/sessions"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) responses.User {
	var response responses.User
	authID := sessions.Get(ctx, "auth")
	userID, _ := strconv.Atoi(authID)
	var userRepo = UserRepository.New()
	user := userRepo.FindByID(userID)

	if user.ID == 0 {
		return response
	}
	return responses.ToUser(user)
}
