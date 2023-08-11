package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoherera/know-vegan-service/src/api/controller"
	"github.com/rodrigoherera/know-vegan-service/src/api/service"
)

func IngredientRoutes(r *gin.Engine, ingredientService service.IIngredientService) {
	i := controller.NewIngredientController(ingredientService)

	r.POST("/v1/ingredient", i.CreateIngredient)
	r.GET("/v1/ingredient", i.GetAll)
}
