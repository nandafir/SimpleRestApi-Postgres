package product

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, service *Service) {
	// conf := config.Get()
	handler := &Handler{
		service: service,
	}
	v1 := router.Group("/v1/anaconda")

	v1survey := v1.Group("/product")
	{
		v1survey.POST("/add", handler.handleAddProduct)
		v1survey.GET("", handler.handleGetProduct)
		// v1survey.DELETE(":id", handler.handleDeleteSurvey)
		// v1survey.PUT(":id", handler.handleUpdateSurvey)
	}
}
