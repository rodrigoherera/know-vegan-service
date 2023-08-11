package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rodrigoherera/know-vegan-service/src/api/domain"
	"github.com/rodrigoherera/know-vegan-service/src/api/service"
	"github.com/rodrigoherera/know-vegan-service/src/api/utils"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	// Validate required fields
	if c.PostForm("name") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Missing 'name' field",
			Data:    nil,
		})
		return
	}
	if c.PostForm("categoryID") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Missing 'categoryID' field",
			Data:    nil,
		})
		return
	}
	if len(c.PostFormArray("tags")) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "At least one 'tag' is required",
			Data:    nil,
		})
		return
	}
	if c.PostForm("type") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Missing 'type' field",
			Data:    nil,
		})
		return
	}
	if c.PostForm("base64") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Missing 'base64' field",
			Data:    nil,
		})
		return
	}
	if c.PostForm("imageName") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Missing 'imageName' field",
			Data:    nil,
		})
		return
	}

	// Convert values to domain types
	catID, err := strconv.Atoi(c.PostForm("categoryID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid 'categoryID' field",
			Data:    nil,
		})
		return
	}
	tagIDs := make([]uint, len(c.PostFormArray("tags")))
	for i, tagIDStr := range c.PostFormArray("tags") {
		tagID, err := strconv.Atoi(tagIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
				Success: false,
				Message: "Invalid 'tags' field",
				Data:    nil,
			})
			return
		}
		tagIDs[i] = uint(tagID)
	}
	photo := &domain.Photo{
		Base64: c.PostForm("base64"),
		Type:   c.PostForm("type"),
		Name:   c.PostForm("imageName"),
	}

	// Create product
	product := &domain.Product{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Ingredients: c.PostForm("ingredients"),
		CategoryID:  uint(catID),
		Tags:        make([]*domain.Tag, len(tagIDs)),
	}
	for i, tagID := range tagIDs {
		product.Tags[i] = &domain.Tag{ID: tagID}
	}

	err = pc.productService.CreateProduct(product, photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "",
		Data:    product,
	})
}
