package http

import (
	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/auth-service/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	Handler AuthHandler
	c *gin.Context
	jwtMgr jwt.Manager
}

func NewRouter(h AuthHandler, c *gin.Context, jwtMgr *jwt.Manager) *gin.Engine{
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
		auth.PATCH("/update/email", JWTAuth(c, jwtMgr), h.UpdateEmail)
		auth.PATCH("/update/username", h.UpdateUsername)
		auth.POST("/refreshToken", h.RefreshToken)
	}
	product := r.Group("/api/ice-cream")
	{
		product.POST("/create",)
	}
	return r
}
