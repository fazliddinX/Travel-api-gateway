package handler

import (
	"Api-Gateway/genproto/content_service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Post Story
// @Description Post Story
// @Tags /api/v1
// @Accept json
// @Produce json
// @Param user body content_service.Story true "Login Request"
// @Success 200 {object} content_service.Story
// @Failure 400 {object} error "Error"
// @Failure 500 {object} error "Error"
// @Router /api/v1/stories [post]
func (h *Handler) PostStory(c *gin.Context) {
	story := content_service.Story{}
	err := c.ShouldBindJSON(&story)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in ShouldBindJSON", "error", err)
		return
	}

	result, err := h.Content.CreateStory(context.Background(), &story)
	if err != nil {
		h.Logger.Error("error in CreateStory", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) PutStory(c *gin.Context) {
	story := content_service.PutStory{}
	err := c.ShouldBindJSON(&story)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in ShouldBindJSON", "error", err)
		return
	}
	story.Id = c.Param("id") // Use c.Param("id") to get the story ID from the path parameter
	result, err := h.Content.UpdateStory(context.Background(), &story)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in UpdateStory", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) DeleteStory(c *gin.Context) {
	story := content_service.StoryId{}
	story.Id = c.Query("story_id")
	result, err := h.Content.DeleteStory(context.Background(), &story)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in DeleteStory", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}
func (h *Handler) GetStory(c *gin.Context) {
	filter := content_service.FilterStories{}

	title := c.Query("title")
	location := c.Query("location")
	a := c.Query("limit")
	b := c.Query("offset")

	limit, err := strconv.Atoi(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetStory", "error", err)
		return
	}
	offset, err := strconv.Atoi(b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetStory", "error", err)
		return
	}
	filter.Title = title
	filter.Location = location
	filter.Limit = int64(limit)
	filter.Offset = int64(offset)

	result, err := h.Content.GetStories(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetStory", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}
func (h *Handler) GetFullInfoStory(c *gin.Context) {
	story_id := c.Query("id")
	user_id, exists := c.Get("user_id")

	fmt.Println(user_id)
	fmt.Println(user_id)
	fmt.Println(user_id)

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		h.Logger.Error("error in user_id not found", "user_id", user_id)
		return
	}
	in := content_service.LikeReq{StoryId: story_id, UserId: user_id.(string)}

	result, err := h.Content.GetFullStoryInfo(context.Background(), &in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetFullInfoStory", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) AddCommentToStory(c *gin.Context) {
	storyComment := content_service.StoryComment{}
	err := c.ShouldBindJSON(&storyComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in ShouldBindJSON", "error", err)
		return
	}
	id, a := c.Get("user_id")
	if !a {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		h.Logger.Error("user_id not found", "user_id", id)
		return
	}
	storyComment.UserId = id.(string)

	result, err := h.Content.AddCommentStory(context.Background(), &storyComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in AddCommentToStory", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)

}
func (h *Handler) GetCommentsStory(c *gin.Context) {
	filter := content_service.FilterComment{}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetCommentsStory", "error", err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetCommentsStory", "error", err)
		return
	}

	filter.Limit = int64(limit)
	filter.Offset = int64(offset)
	filter.StoryId = c.Query("story_id")

	result, err := h.Content.GetComments(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetCommentsStory", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) AddLikeStory(c *gin.Context) {
	like := content_service.LikeReq{}

	us, ex := c.Get("user_id")
	if !ex {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		h.Logger.Error("user_id not found", "user_id", us)
		return
	}
	like.StoryId = c.Query("story_id")
	like.UserId = us.(string)

	result, err := h.Content.AddLikeToStory(context.Background(), &like)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in AddLikeToStory", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) CreateItineraries(c *gin.Context) {
	itineraries := content_service.Itineraries{}
	err := c.ShouldBindJSON(&itineraries)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in ShouldBindJSON", "error", err)
		return
	}
	id, a := c.Get("user_id")
	if !a {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found in context"})
		h.Logger.Error("user_id not found", "user_id", id)
		return
	}
	itineraries.AuthorId = id.(string)
	result, err := h.Content.CreateItineraries(context.Background(), &itineraries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in CreateItineraries", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateItineraries(c *gin.Context) {
	putItineraries := content_service.PutItineraries{}
	err := c.ShouldBindJSON(&putItineraries)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in ShouldBindJSON", "error", err)
		return
	}
	putItineraries.Id = c.Query("id")

	result, err := h.Content.UpdateItineraries(context.Background(), &putItineraries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in UpdateItineraries", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) DeleteItineraries(c *gin.Context) {
	itineraries := content_service.ItinerariesId{}
	itineraries.Id = c.Query("id")

	result, err := h.Content.DeleteItineraries(context.Background(), &itineraries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in DeleteItineraries", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetAllItineraries(c *gin.Context) {
	filter := content_service.FilterItineraries{}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetAllItineraries", "error", err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetAllItineraries", "error", err)
		return
	}
	likeCount, err := strconv.Atoi(c.Query("like_count"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetAllItineraries", "error", err)
		return
	}
	filter.Limit = int64(limit)
	filter.Offset = int64(offset)
	filter.Title = c.Query("title")
	filter.LikeCount = int64(likeCount)

	result, err := h.Content.GetItineraries(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetAllItineraries", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetFullInfoItineraries(c *gin.Context) {
	c.JSON(200, "nma gap e")
}
func (h *Handler) AddCommentToItineraries(c *gin.Context) {
	comment := content_service.ItinerariesComment{}
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in AddCommentToItineraries", "error", err)
		return
	}
	comment.ItineraryId = c.Query("itinerary_id")

	result, err := h.Content.AddCommentItineraries(context.Background(), &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in AddCommentToItineraries", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetDestinations(c *gin.Context) {
	filter := content_service.FilterDestinations{}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetDestinations", "error", err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetDestinations", "error", err)
		return
	}
	filter.Limit = int64(limit)
	filter.Offset = int64(offset)
	filter.Name = c.Query("name")
	filter.Country = c.Query("country")

	result, err := h.Content.GetDescriptions(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetDestinations", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetDestinationsByID(c *gin.Context) {
	id := content_service.DestinationId{}
	id.Id = c.Query("destination_id")

	result, err := h.Content.GetDestinationById(context.Background(), &id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetDestinationsByID", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) SendMessage(c *gin.Context) {
	message := content_service.MessageReq{}
	err := c.ShouldBindJSON(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in SendMessage", "error", err)
		return
	}
	user_id, ex := c.Get("user_id")
	if !ex {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		h.Logger.Error("unauthorized error in SendMessage", "error", err)
		return
	}
	message.SenderId = user_id.(string)

	result, err := h.Content.CreateMessage(context.Background(), &message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in SendMessage", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetMessage(c *gin.Context) {
	message := content_service.GetMessage{}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetMessage", "error", err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetMessage", "error", err)
		return
	}
	message.Limit = int64(limit)
	message.Offset = int64(offset)

	result, err := h.Content.GetMessages(context.Background(), &message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetMessage", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) CreateTravelTips(c *gin.Context) {
	travel := content_service.TravelTipReq{}
	err := c.ShouldBindJSON(&travel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.Logger.Error("error in CreateTravelTips", "error", err)
		return
	}
	id, ex := c.Get("user_id")
	if !ex {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		h.Logger.Error("unauthorized error in CreateTravelTips", "error", err)
		return
	}
	travel.AuthorId = id.(string)

	result, err := h.Content.CreateTravelTips(context.Background(), &travel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in CreateTravelTips", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetTravelTips(c *gin.Context) {
	filter := content_service.FilterTravelTip{}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetTravelTips", "error", err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Atoi"})
		h.Logger.Error("error in GetTravelTips", "error", err)
		return
	}
	filter.Limit = int64(limit)
	filter.Offset = int64(offset)
	filter.Category = c.Query("category")
	result, err := h.Content.GetTravelTips(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetTravelTips", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetUserStatistic(c *gin.Context) {
	id := content_service.UserId{}
	id.UserId = c.Query("user_id")

	result, err := h.Content.GetUserStatistics(context.Background(), &id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetUserStatistics", "error", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetTrendingDestination(c *gin.Context) {
	a := content_service.Void{}
	result, err := h.Content.GetTrendingDestinations(context.Background(), &a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.Logger.Error("error in GetTrendingDestinations", "error", err)
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		return
	}
	c.JSON(http.StatusOK, result)
}
