package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pangdfg/MagicStream/Server/db"
	"github.com/pangdfg/MagicStream/Server/env"
	"github.com/pangdfg/MagicStream/Server/routes"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		})
	})

	if err := env.Check(); err != nil {
			log.Fatalf("env check failed: %v", err)
		}
	
	allowedOrigins := env.GetString("ALLOWED_ORIGINS", "")
	var origins []string
	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
			log.Println("Allowed Origin:", origins[i])
		}
	} else {
		origins = []string{"http://localhost:5173"}
		log.Println("Allowed Origin: http://localhost:5173")
	}

	config := cors.Config{}
	config.AllowOrigins = origins
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	app.Use(cors.New(config))
	app.Use(gin.Logger())

	var client *mongo.Client = db.Connect()
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to reach server: %v", err)
	}
	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}

	}()
	
	routes.DefaultRoutes(app, client)
	routes.LoginedRoutes(app, client)

	if err := app.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}