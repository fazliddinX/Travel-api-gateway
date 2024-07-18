package handler

import (
	"api-gateway/generated/stories"
	"api-gateway/models"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Travel Story
// @Description Create new Travel Story
// @Tags Stories
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param TravelStory body stories.CreateTravelStoryRequest true "TravelStory"
// @Success 200 {object} stories.CreateTravelStoryResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/stories [post]
func (h *Handler) CreateTravelStory(ctx *gin.Context) {
	var req stories.CreateTravelStoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}
	resp, err := h.StoriesClient.CreateTravelStory(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik sayohat hikoyasini yaratishda", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat hikoyasini yaratishda ",
		})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// @Summary Update Travel Story
// @Description Update an existing Travel Story
// @Tags Stories
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param TravelStory body stories.UpdateTravelStoryRequest true "TravelStory"
// @Success 200 {object} stories.UpdateTravelStoryResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/stories/{storyId} [put]
func (h *Handler) UpdateTravelStory(ctx *gin.Context) {
	var req stories.UpdateTravelStoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}

	resp, err := h.StoriesClient.UpdateTravelStory(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik sayohat hikoyasini yangilashda", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat hikoyasini yangilashda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Travel Story
// @Description Delete an existing Travel Story
// @Tags Stories
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param storyId path string true "Story ID"
// @Success 200 {object} stories.DeleteTravelStoryResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/stories/{storyId} [delete]
func (h *Handler) DeleteTravelStory(ctx *gin.Context) {
	storyId := ctx.Param("storyId")
	req := &stories.DeleteTravelStoryRequest{TravelStoryId: storyId}

	resp, err := h.StoriesClient.DeleteTravelStory(ctx, req)
	if err != nil {
		h.Logger.Error("xatolik sayohat hikoyasini o'chirishda", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat hikoyasini o'chirishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary List Travel Stories
// @Description List all Travel Stories
// @Tags Stories
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param page query int fakse "Page number"
// @Param limit query int false "Page limit"
// @Success 200 {object} stories.ListTravelStoryResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/stories [get]
func (h *Handler) ListTravelStory(ctx *gin.Context) {
	var req stories.ListTravelStoryRequest
	// Get 'page' query parameter and convert it to int32
	pageStr := ctx.Query("page")
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.Logger.Error("Error converting page to int", slog.String("error", err.Error()))
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing page from string to int",
			})
			return
		}
		req.Page = int32(page)
	} else {
		req.Page = 1
	}
	limitStr := ctx.Query("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			h.Logger.Error("Error converting limit to int", "error", err.Error())
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing limit from string to int",
			})
			return
		}
		req.Limit = int32(limit)
	} else {
		req.Limit = 10
	}
	resp, err := h.StoriesClient.ListTravelStory(ctx, &req)
	if err != nil {
		h.Logger.Error("error in getting stories", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "error in getting stories",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Travel Story
// @Description Get a single Travel Story by ID
// @Tags Stories
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param storyId path string true "Story ID"
// @Success 200 {object} stories.GetTravelStoryResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/stories/{storyId} [get]
func (h *Handler) GetTravelStory(ctx *gin.Context) {
	storyId := ctx.Param("storyId")
	req := &stories.GetTravelStoryRequest{StoryId: storyId}
	resp, err := h.StoriesClient.GetTravelStory(ctx, req)
	if err != nil {
		h.Logger.Error("error in get story", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "error in get story",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Add Comment to Travel Story
// @Description Add a comment to a Travel Story
// @Tags Stories
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param comment body stories.AddCommentRequest true "Comment"
// @Success 200 {object} stories.AddCommentResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/stories/{storyId}/comments [post]
func (h *Handler) AddComment(ctx *gin.Context) {
	var req stories.AddCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}
	resp, err := h.StoriesClient.AddCommment(ctx, &req)
	if err != nil {
		h.Logger.Error("error in post comment", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "error in post comment",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary List Comments of Travel Story
// @Description List all comments of a Travel Story
// @Tags Stories
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param storyId path string true "Story ID"
// @Param page query int false "Page number"
// @Param limit query int false "Page limit"
// @Success 200 {object} stories.ListCommentsResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/stories/{storyId}/comments [get]
func (h *Handler) ListComments(ctx *gin.Context) {
	id := ctx.Param("storyId")
	var req stories.ListCommentsRequest
	// Get 'page' query parameter and convert it to int32
	pageStr := ctx.Query("page")
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.Logger.Error("Error converting page to int", "error", err.Error())
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing page from string to int",
			})
			return
		}
		req.Page = int32(page)
	} else {
		req.Page = 1
	}
	req.StoryId = id
	limitStr := ctx.Query("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			h.Logger.Error("Error converting limit to int", slog.String("error", err.Error()))
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing limit from string to int",
			})
			return
		}
		req.Limit = int32(limit)
	} else {
		req.Limit = 10
	}
	resp, err := h.StoriesClient.ListComments(ctx, &req)
	if err != nil {
		h.Logger.Error("error in get comments", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "error in get comments",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Add Like to Travel Story
// @Description Add a like to a Travel Story
// @Tags Stories
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param like body stories.AddLikeRequest true "Like"
// @Success 200 {object} stories.AddLikeResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/stories/{storyId}/like [post]
func (h *Handler) AddLike(ctx *gin.Context) {
	var req stories.AddLikeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", slog.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}
	resp, err := h.StoriesClient.AddLike(ctx, &req)
	if err != nil {
		h.Logger.Error("error in like", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "error in like",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
