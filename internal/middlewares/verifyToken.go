package middlewares

import (
	"log"
	"main/internal/models"
	"main/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func InitMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) VerifyToken(ctx *gin.Context) {
	// 1. ambil token dari header
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Msg: "Please login first"})
		return
	}
	// verifikasi bearer token
	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Msg: "Please login first"})
		return
	}
	// 2. pisahkan token dari bearer
	token := strings.Split(bearerToken, " ")[1]
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Msg: "Please login first"})
		return
	}
	// 3. verifikasi token
	claims := &pkg.Claims{}
	if err := claims.VerifyToken(token); err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "expired") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Msg: "Session expired. Please re-login"})
			return
		}
		if strings.Contains(err.Error(), "malformed") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Msg: "Login identity is malformed. Please re-login"})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	// masukkan claims/payload ke gin context
	ctx.Set("Payload", claims)
	ctx.Set("user_id", claims.UserID)
	ctx.Set("role", claims.Role)
	ctx.Next()
}
