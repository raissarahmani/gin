package middlewares

import (
	"main/internal/models"

	// "main/pkg"
	"net/http"
	// "slices"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AccessGateAdmin(ctx *gin.Context) {
	// 1. ambil payload/claims dari context gin
	role, exist := ctx.Get("role")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Msg: "Please login first"})
		return
	}
	// type assertion claims menjadi pkg.claims
	roleStr, ok := role.(string)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Msg: "Login identity is malformed. Please re-login"})
		return
	}
	// cek role yang ada di claims
	if roleStr != "admin" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, models.Response{Msg: "You can not access this page"})
		return
	}
	ctx.Next()
}

// untuk setting beberapa role
// func (m *Middleware) AccessGate(allowedRole ...string) func(*gin.Context) {
// 	// 1. ambil payload/claims dari context gin
// 	claims, exist := ctx.Get("Payload")
// 	if !exist {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Msg: "Please login first"})
// 		return
// 	}
// 	// type assertion claims menjadi pkg.claims
// 	userClaims, ok := claims.(*pkg.Claims)
// 	if !ok {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Msg: "Login identity is malformed. Please re-login"})
// 		return
// 	}
// 	// cek role yang ada di claims
// 	if !slices.Contains(allowedRole, userClaims.Role) {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, models.Response{Msg: "You do not have permission to access this page"})
// 		return
// 	}
// 	ctx.Next()
// }
