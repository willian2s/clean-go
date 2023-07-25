package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/willian2s/clean-go/adapter/http/middleware"
	"github.com/willian2s/clean-go/adapter/postgres"
	"github.com/willian2s/clean-go/config"
	"github.com/willian2s/clean-go/di"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/willian2s/clean-go/adapter/http/docs"
)

// @title Clean GO API Docs
// @version 1.0.0
// @contact.name Willian Silva
// @license.name MIT
// @license.url https://mit-license.org/
// @host localhost:3000
// @BasePath /
func main() {
	config.ConfigRuntime()

	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	productService := di.ConfigProductDI(conn)

	mode := config.EnvConfigs.Mode
	gin.SetMode(mode)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(middleware.RateLimit, gin.Recovery())
	router.Use(middleware.Cors)

	product := router.Group("/product")
	{
		product.POST("/", func(c *gin.Context) {
			productService.Create(c.Writer, c.Request)
		})
		product.GET("/", func(c *gin.Context) {
			productService.Fetch(c.Writer, c.Request)
		})
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := config.EnvConfigs.ServerPort
	router.Run(fmt.Sprintf(":%s", port))
}
