package controller

import (
	"API_Rest_GO/model"
	"API_Rest_GO/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProducterController(usecase usecase.ProductUseCase) productController {
	return productController{
		ProductUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID : 1,
			Name: "batata frita",
			Price: 20,

		},
	}

	ctx.JSON(http.StatusOK, products)
}