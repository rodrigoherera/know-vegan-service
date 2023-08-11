package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoherera/know-vegan-service/src/api/client/mysql"
	"github.com/rodrigoherera/know-vegan-service/src/api/client/s3"
	"github.com/rodrigoherera/know-vegan-service/src/api/config"
	"github.com/rodrigoherera/know-vegan-service/src/api/middleware"
	"github.com/rodrigoherera/know-vegan-service/src/api/persistance/repository"
	"github.com/rodrigoherera/know-vegan-service/src/api/routes"
	"github.com/rodrigoherera/know-vegan-service/src/api/service"
)

func Execute() {
	gin.SetMode(config.MODE)

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.AuthMiddleware())

	categoryRepository := repository.NewCategoryRepository(mysql.DB)
	productRepository := repository.NewProductRepository(mysql.DB)
	ingredientRepository := repository.NewIngredientRepository(mysql.DB)

	categoryService := service.NewCategoryService(categoryRepository)
	productService := service.NewProductService(productRepository, s3.GetSession)
	ingredientService := service.NewIngredientService(ingredientRepository)

	routes.ProductRoutes(r, productService)
	routes.CategoryRoutes(r, categoryService)
	routes.IngredientRoutes(r, ingredientService)

	err := r.Run(config.HTTPPORT)
	if err != nil {
		panic(err)
	}
}
