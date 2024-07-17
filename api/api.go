package api

import (
	"Api-Gateway/api/handler"
	_ "Api-Gateway/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Gateway
// @version 1.0
// @description This is a sample server for API Gateway.
// @host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Router(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins: []string{"http://localhost", "http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}
	router.Use(cors.New(corsConfig))
	// Swagger endpointini sozlash

	r := router.Group("api/v1")

	//r.Use(middleware.MiddleWare())

	r.POST("/stories", h.PostStory)
	r.PUT("/stories/:id", h.PutStory)
	r.DELETE("/stories/:id", h.DeleteStory)
	r.GET("/stories", h.GetStory)
	r.GET("/stories/:id", h.GetFullInfoStory)
	r.POST("/stories/:id/comments", h.AddCommentToStory)
	r.GET("/stories/:id/comments", h.GetCommentsStory)
	r.POST("/stories/:id/like", h.AddLikeStory)

	r.POST("/itineraries", h.CreateItineraries)
	r.PUT("/itineraries/:id", h.UpdateItineraries)
	r.DELETE("/itineraries/:id", h.DeleteItineraries)
	r.GET("/itineraries", h.GetAllItineraries)
	r.GET("/itineraries/:id", h.GetFullInfoItineraries)
	r.POST("/itineraries/:id/comments", h.AddCommentToItineraries)

	r.GET("/destinations", h.GetDestinations)
	r.GET("/destinations/:id", h.GetDestinationsByID)

	r.POST("/messages", h.SendMessage)
	r.GET("/messages", h.GetMessage)

	r.POST("/travel_tips", h.CreateTravelTips)
	r.GET("/travel_tips", h.GetTravelTips)

	r.GET("/user/:id/statistics", h.GetUserStatistic)

	r.GET("/trending-destinations", h.GetTrendingDestination)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
