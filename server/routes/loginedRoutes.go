package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/pangdfg/MagicStream/Server/controllers"
	"github.com/pangdfg/MagicStream/Server/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func LoginedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie(client))
	router.POST("/addmovie", controller.AddMovie(client))
	router.GET("/recommendedmovies", controller.GetRecommendedMovies(client))
	router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate(client))
}