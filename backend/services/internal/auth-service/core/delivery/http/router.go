package http

import (
	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	Handler AuthHandler
	c       *gin.Context
	jwtMgr  jwt.Manager
}

func RegisterRouter(r *gin.Engine, h AuthHandler, c *gin.Context, jwtMgr *jwt.Manager) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
		auth.PATCH("/update/email", JWTAuth(c, jwtMgr), h.UpdateEmail)
		auth.PATCH("/update/username", h.UpdateUsername)
		auth.POST("/refreshToken", h.RefreshToken)
	}
}
