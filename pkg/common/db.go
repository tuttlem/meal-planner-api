package common

import (
	"gorm.io/gorm/logger"
	"log"

	"github.com/tuttlem/meal-planner-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbInit(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Meal{})
	db.AutoMigrate(&models.Ingredient{})

	return db
}
