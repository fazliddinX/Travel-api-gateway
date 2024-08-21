package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/api/handler/docs"
)

// @title Api Get Eway
// @version 1.0
// @description Api-geteway for user and story.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
// @schemes http
func NewRouter(handle *handler.Handler) *gin.Engine {
	router := gin.Default()

	// Swagger endpointini sozlash
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.AuthMiddleware())
	//router.Use(middleware.LoggerMiddleware())

	st := router.Group("/api/v1/stories")
	st.POST("/", handle.CreateTravelStory)
	st.PUT("/:storyId", handle.UpdateTravelStory)
	st.DELETE("/:storyId", handle.DeleteTravelStory)
	st.GET("/", handle.ListTravelStory)
	st.GET("/:storyId", handle.GetTravelStory)
	st.POST("/:storyId/comments", handle.AddComment)
	st.GET("/:storyId/comments", handle.ListComments)
	st.POST("/:storyId/like", handle.AddLike)

	iti := router.Group("/api/v1/itineraries")
	iti.POST("/", handle.CreateItinerary)
	iti.PUT("/:itineraryId", handle.UpdateItinerary)
	iti.DELETE("/:itineraryId", handle.DeleteItinerary)
	iti.GET("/", handle.ListItineraries)
	iti.GET("/:itineraryId", handle.GetItinerary)
	iti.POST("/:itineraryId/comments", handle.LeaveComment)

	destin := router.Group("/api/v1/destinations")
	destin.GET("/", handle.ListTravelDestinations)
	destin.GET("/:destinationId", handle.GetTravelDestination)
	destin.GET("/trending", handle.GetTrendDestinations)

	msg := router.Group("/api/v1/messages")
	msg.POST("/", handle.SendMessageUser)
	msg.GET("/", handle.ListMessage)

	tips := router.Group("/api/v1/travel-tips")
	tips.POST("/", handle.AddTravelTips)
	tips.GET("/", handle.GetTravelTips)

	us := router.Group("/api/v1/users")
	us.GET("/profile", handle.GetUserProfileHandle)
	us.PUT("/profile", handle.UpdateUserProfileHandle)
	us.GET("/", handle.ListUsersHandle)
	us.DELETE("/:userId", handle.DeleteUser)
	us.GET("/:userId/activity", handle.GetUserActivity)
	us.POST("/:userId/follow", handle.FollowUser)
	us.GET("/:userId/followers", handle.ListFollowers)
	us.GET(":userId/statics", handle.GetUserStatics)

	return router
}
