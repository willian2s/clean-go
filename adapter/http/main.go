package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	"github.com/willian2s/clean-go/adapter/postgres"
	"github.com/willian2s/clean-go/config"
	"github.com/willian2s/clean-go/di"

	httpSwagger "github.com/swaggo/http-swagger/v2"
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

	app := chi.NewRouter()

	app.Use(middleware.RequestID)
	app.Use(middleware.Recoverer)
	app.Use(middleware.Logger)
	app.Use(httprate.LimitByIP(100, 1*time.Minute))
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	app.Route("/product", func(r chi.Router) {
		r.Get("/", productService.Fetch)
		r.Post("/", productService.Create)
	})

	app.Get("/swagger/*", httpSwagger.WrapHandler)

	port := config.EnvConfigs.ServerPort
	http.ListenAndServe(fmt.Sprintf(":%s", port), app)
}
