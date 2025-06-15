package middlewares

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) CORSMiddleware(ctx *gin.Context) {
	whitelistOrigin := []string{"http://localhost:5173", "https://tickitz-react.vercel.app"}
	origin := ctx.GetHeader("Origin")

	if slices.Contains(whitelistOrigin, origin) {
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
	} else {
		// Origin not allowed
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "CORS origin denied"})
		return
	}

	ctx.Next()
}
