package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoherera/know-vegan-service/src/api/controller"
	"github.com/rodrigoherera/know-vegan-service/src/api/service"
)

func CategoryRoutes(r *gin.Engine, categoryService service.ICategoryService) {
	c := controller.NewCategoryController(categoryService)

	r.POST("/v1/category", c.CreateCategory)
	r.GET("/v1/category", c.GetAllCategories)
}
