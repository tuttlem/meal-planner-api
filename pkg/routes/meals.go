package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tuttlem/meal-planner-api/pkg/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type AddMealRequestBody struct {
	Name string `json:"name"`
}

type UpdateMealRequestBody struct {
	Name string `json:"name"`
}

type UpdateMealIngredientsRequestBody struct {
	Ingredients []int `json:"ingredients"`
}

func (h handler) AddMeal(c *gin.Context) {
	body := AddMealRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		log.Print(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var meal models.Meal

	meal.Name = body.Name

	if result := h.DB.Create(&meal); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &meal.ID)
}

func (h handler) ListMeals(c *gin.Context) {
	var meals []models.Meal

	if result := h.DB.Scopes(Paginate(c)).Find(&meals); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
	}

	c.JSON(http.StatusOK, &meals)
}

func (h handler) GetMeal(c *gin.Context) {
	id := c.Param("id")
	var meal models.Meal

	if result := h.DB.First(&meal, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, &meal)
}

func (h handler) UpdateMeal(c *gin.Context) {
	id := c.Param("id")
	body := UpdateMealRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var meal models.Meal

	if result := h.DB.First(&meal, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	meal.Name = body.Name

	h.DB.Save(&meal)
	c.JSON(http.StatusNoContent, &meal)
}

func (h handler) DeleteMeal(c *gin.Context) {
	id := c.Param("id")
	var meal models.Meal

	if result := h.DB.First(&meal, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	h.DB.Delete(&meal)
	c.JSON(http.StatusOK, &meal)
}

func (h handler) ListMealIngredients(c *gin.Context) {
	id := c.Param("id")
	var meal models.Meal

	if result := h.DB.Preload("Ingredients").First(&meal, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, &meal.Ingredients)
}

func (h handler) UpdateMealIngredients(c *gin.Context) {
	id := c.Param("id")
	body := UpdateMealIngredientsRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var meal models.Meal

	if result := h.DB.Preload("Ingredients").First(&meal, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	if len(body.Ingredients) == 0 {
		h.DB.Model(&meal).Association("Ingredients").Delete(meal.Ingredients)
		c.JSON(http.StatusNoContent, &meal)
		return
	}

	var ingredients []*models.Ingredient

	if result := h.DB.Find(&ingredients, body.Ingredients); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	meal.Ingredients = ingredients
	h.DB.Model(&meal).Association("Ingredients").Replace(meal.Ingredients)
	c.JSON(http.StatusNoContent, &meal)
}
