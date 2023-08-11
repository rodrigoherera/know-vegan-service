package controller

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rodrigoherera/know-vegan-service/src/api/domain"
	"github.com/rodrigoherera/know-vegan-service/src/api/service"
	"github.com/rodrigoherera/know-vegan-service/src/api/utils"
)

type IngredientController struct {
	ingredientService service.IIngredientService
}

func NewIngredientController(ingredientService service.IIngredientService) *IngredientController {
	return &IngredientController{
		ingredientService: ingredientService,
	}
}

func (ic *IngredientController) CreateIngredient(c *gin.Context) {
	var (
		ingredient *domain.Ingredient
	)

	if err := c.ShouldBindJSON(&ingredient); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ic.ingredientService.CreateIngredient(ingredient)

	c.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "",
		Data:    ingredient,
	})
}

func (ic *IngredientController) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid 'page' field",
			Data:    nil,
		})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid 'limit' field",
			Data:    nil,
		})
		return
	}

	offset := int64((page - 1) * limit)

	ingredients, total := ic.ingredientService.GetAll(offset, int64(limit))

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	if ingredients == nil {
		c.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "Ingredients not found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.ResponsePagination{
		Success: true,
		Metadata: utils.Metadata{
			Page:       int64(page),
			Limit:      int64(limit),
			Total:      total,
			TotalPages: int64(totalPages),
		},
		Data: ingredients,
	})
}
