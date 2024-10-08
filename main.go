package main

import (
	"go-boilerplate/database"
	"go-boilerplate/handlers"
	"go-boilerplate/repositories"
	"go-boilerplate/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Postgresql()
	defer db.Close()

	r := repositories.NewRepositorie(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	router.Use(cors.New(config))

	router.GET("/api/text", h.Test)

	err := router.Run(":8888")
	if err != nil {
		panic(err.Error())
	}
}
