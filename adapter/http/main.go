package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willian2s/clean-go/adapter/postgres"
	"github.com/willian2s/clean-go/config"
	"github.com/willian2s/clean-go/di"

	httpSwagger "github.com/swaggo/http-swagger"
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
	config.InitEnvConfigs()

	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	productService := di.ConfigProductDI(conn)

	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.Handle("/product", http.HandlerFunc(productService.Create)).Methods("POST")
	router.Handle("/product", http.HandlerFunc(productService.Fetch)).Queries().Methods("GET")

	port := config.EnvConfigs.ServerPort
	log.Printf("Listening on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
