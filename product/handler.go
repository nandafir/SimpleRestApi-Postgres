package product

import (
	"anaconda/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	// "xapiens.id/shared/errs"
	// "xapiens.id/shared/auth"
	// "xapiens.id/shared/errs"
	// "xapiens.id/shared/utils"
)

type Handler struct {
	service *Service
}

func (h Handler) handleAddProduct(c *gin.Context) {

	//Binding to params var
	var params ProductRequest
	// username := c.GetString("UserName")

	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	err := h.service.SubmitProduct(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.Response("success"))
}

func (h Handler) handleGetProduct(c *gin.Context) {

	res, err := h.service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.Response(res))
}
