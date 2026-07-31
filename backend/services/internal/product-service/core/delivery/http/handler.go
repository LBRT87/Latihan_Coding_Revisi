package http

import (
	"io"
	"net/http"
	"strconv"

	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/product-service/core/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase usecase.IceCreamUsecaseInterface
	cookieSecure   bool
}

func NewProductHandler(productUsecase usecase.IceCreamUsecaseInterface, cookieSecure bool) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
		cookieSecure:   cookieSecure,
	}
}

func (ph *ProductHandler) Create(c *gin.Context) {
	var req usecase.CreateIceCreamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	fileHeader, err := c.FormFile("Photo")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":"photo must be upload",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":"error while opening the file",
		})
	}

	defer file.Close()

	photoByte, err := io.ReadAll(file)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":"error while reading the file",
		})
	}

	if err := ph.productUsecase.Create(c.Request.Context(), req, photoByte); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success creating ice cream"})
}

func (ph *ProductHandler) GetDetail(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	if iceCream, err := ph.productUsecase.GetByID(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, iceCream)
	}
}

func (ph *ProductHandler) GetList(c *gin.Context) {
	var req usecase.ListIceCreamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	data, err := ph.productUsecase.GetAll(c.Request.Context(), req)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"iceCreams":data.IceCreams,
		"totalData":data.TotalData,
	})

}

func (ph *ProductHandler) Update(c *gin.Context) {
	var req usecase.UpdateIceCreamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := ph.productUsecase.Update(c.Request.Context(), req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (ph *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	if err := ph.productUsecase.Delete(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}