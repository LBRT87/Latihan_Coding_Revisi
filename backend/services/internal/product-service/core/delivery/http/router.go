package http

import (
	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, h ProductHandler, c *gin.Context, jwtMgr *jwt.Manager) {
	prod := r.Group("/api/icecream")
	{
		prod.GET("", h.GetList)
		prod.POST("/create", h.Create)
		prod.DELETE("/detail/:id", h.Delete)
		prod.DELETE("/delete/:id", h.Delete)
		prod.PATCH("/price", h.Update)
	}
}
