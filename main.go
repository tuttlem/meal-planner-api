package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tuttlem/meal-planner-api/pkg/common"
	"github.com/tuttlem/meal-planner-api/pkg/routes"
	"log"
)

func main() {
	appConfig, err := common.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to read config: %s", err.Error())
	}

	server := gin.New()
	// server.RedirectTrailingSlash = false
	// server.RedirectFixedPath = false
	// server.HandleMethodNotAllowed = false
	// server.ForwardedByClientIP = true

	// server.UseRawPath = false
	// server.UnescapePathValues = true

	server.Use(common.JSONLogger())

	db := common.DbInit(appConfig.MealsDB)
	routes.RegisterRoutes(server, db)

	server.Run(fmt.Sprintf(":%d", appConfig.Server.Port))
}
