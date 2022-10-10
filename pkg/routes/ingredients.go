package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tuttlem/meal-planner-api/pkg/models"
	"gorm.io/gorm"
	"net/http"
)

type AddIngredientRequestBody struct {
	Name string `json:"name"`
}

type UpdateIngredientRequestBody struct {
	Name string `json:"name"`
}

func (h handler) AddIngredient(c *gin.Context) {
	body := AddIngredientRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var Ingredient models.Ingredient

	Ingredient.Name = body.Name

	if result := h.DB.Create(&Ingredient); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &Ingredient.ID)
}

func (h handler) ListIngredients(c *gin.Context) {
	var Ingredients []models.Ingredient

	if result := h.DB.Scopes(Paginate(c)).Find(&Ingredients); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
	}

	c.JSON(http.StatusOK, &Ingredients)
}

func (h handler) GetIngredient(c *gin.Context) {
	id := c.Param("id")
	var Ingredient models.Ingredient

	if result := h.DB.First(&Ingredient, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, &Ingredient)
}

func (h handler) UpdateIngredient(c *gin.Context) {
	id := c.Param("id")
	body := UpdateIngredientRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var Ingredient models.Ingredient

	if result := h.DB.First(&Ingredient, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	Ingredient.Name = body.Name

	h.DB.Save(&Ingredient)
	c.JSON(http.StatusOK, &Ingredient)
}

func (h handler) DeleteIngredient(c *gin.Context) {
	id := c.Param("id")
	var Ingredient models.Ingredient

	if result := h.DB.First(&Ingredient, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	h.DB.Delete(&Ingredient)
	c.JSON(http.StatusOK, &Ingredient)
}
