package middlewares

import (
	UserRepository "blog/internal/modules/user/repositiories"
	"blog/pkg/sessions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(ctx *gin.Context) {
		authID := sessions.Get(ctx, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID != 0 {
			ctx.Redirect(http.StatusFound, "/")
		}
		// before request

		ctx.Next()

	}
}
