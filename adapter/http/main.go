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
)

// func init() {
// 	viper.SetConfigFile(`config.json`)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	config.InitEnvConfigs()

	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	productService := di.ConfigProductDI(conn)

	router := mux.NewRouter()
	router.Handle("/product", http.HandlerFunc(productService.Create)).Methods("POST")
	router.Handle("/product", http.HandlerFunc(productService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	port := config.EnvConfigs.ServerPort
	log.Printf("Listening on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
