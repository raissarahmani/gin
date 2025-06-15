package middlewares

import (
	"log"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) CORSMiddleware(ctx *gin.Context) {
	// setup whitelist origin
	whitelistOrigin := []string{"http://localhost:5173", "https://tickitz-react.vercel.app"}
	origin := ctx.GetHeader("Origin")
	log.Println("[DEBUG] Origin: ", origin)
	allowedOrigin := ""
	if slices.Contains(whitelistOrigin, origin) {
		allowedOrigin = origin
	}
	ctx.Header("Access-Control-Allow-Origin", allowedOrigin)
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	// handle preflight
	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()
}
