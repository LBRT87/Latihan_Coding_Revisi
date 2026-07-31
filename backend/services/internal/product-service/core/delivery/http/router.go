package http

import (
	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	Handler ProductHandler
	c       *gin.Context
	jwtMgr  jwt.Manager
}

func RegisterRouter(r *gin.Engine, h ProductHandler, c *gin.Context, jwtMgr *jwt.Manager) {
	prod := r.Group("/api/product")
	{
		prod.POST("/create", JWTAuth(c, jwtMgr), h.Create)
	}
}
