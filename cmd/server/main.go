package main

import (
	"fmt"
	"go-farms/configs"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
)

func init() {
	configs.LoadConfig()
	configs.ConnectDB()
}

func main() {
	r := chi.NewRouter()

	StartRoutes(r)

	port := viper.Get("SERVER_PORT")

	fmt.Printf("Server running on port %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		panic(err)
	}
}

func StartRoutes(r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/farm", func(r chi.Router) {
			// r.Post("/", CreateFarm)
			// r.Get("/", GetFarms)
		})
	})
}
