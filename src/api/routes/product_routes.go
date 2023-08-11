package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoherera/know-vegan-service/src/api/controller"
	"github.com/rodrigoherera/know-vegan-service/src/api/service"
)

func ProductRoutes(r *gin.Engine, productService service.IProductService) {
	c := controller.NewProductController(productService)

	r.POST("/v1/product", c.CreateProduct)
}
