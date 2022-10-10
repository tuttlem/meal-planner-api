package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(server *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	server.POST("/meals", h.AddMeal)
	server.GET("/meals", h.ListMeals)
	server.GET("/meals/:id", h.GetMeal)
	server.PUT("/meals/:id", h.UpdateMeal)
	server.DELETE("/meals/:id", h.DeleteMeal)

	server.GET("/meals/:id/ingredients", h.ListMealIngredients)
	server.PUT("/meals/:id/ingredients", h.UpdateMealIngredients)

	server.POST("/ingredients", h.AddIngredient)
	server.GET("/ingredients", h.ListIngredients)
	server.GET("/ingredients/:id", h.GetIngredient)
	server.PUT("/ingredients/:id", h.UpdateIngredient)
	server.DELETE("/ingredients/:id", h.DeleteIngredient)
}
