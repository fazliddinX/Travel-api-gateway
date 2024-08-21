package handler

import (
	"api-gateway/generated/communication"
	"api-gateway/generated/destination"
	"api-gateway/models"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary List Travel Destinations
// @Description List all Travel Destinations
// @Tags Destinations
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param page query int fale "Page number"
// @Param limit query int false "Page limit"
// @Success 200 {object} destination.ListDetinationResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/destinations [get]
func (h *Handler) ListTravelDestinations(ctx *gin.Context) {
	var req destination.ListDetinationRequest
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

	// Get 'limit' query parameter and convert it to int32
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

	resp, err := h.DestinationClient.ListTravelDestnations(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik sayohat manzillarini ro'yxatlashda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat manzillarini ro'yxatlashda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Travel Destination
// @Description Get a single Travel Destination by ID
// @Tags Destinations
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param destinationId path string true "Destination ID"
// @Success 200 {object} destination.GetDestinationResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/destinations/{destinationId} [get]
func (h *Handler) GetTravelDestination(ctx *gin.Context) {
	destinationId := ctx.Param("destinationId")
	req := &destination.GetDestinationRequest{Id: destinationId}

	resp, err := h.DestinationClient.GetTravelDestination(ctx, req)
	if err != nil {
		h.Logger.Error("xatolik sayohat manzilini olishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat manzilini olishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Trend Travel Destinations
// @Description Get trending Travel Destinations
// @Tags Destinations
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param Top query int false "Top Destination"
// @Success 200 {object} destination.GetTrendDestinationResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/destinations/trending [get]
func (h *Handler) GetTrendDestinations(ctx *gin.Context) {
	req := &destination.GetTrendDestinationRequest{}

	topStr := ctx.Query("Top")

	if topStr != "" {
		page, err := strconv.Atoi(topStr)
		if err != nil {
			h.Logger.Error("Error converting page to int", slog.String("error", err.Error()))
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing page from string to int",
			})
			return
		}
		req.Limit = int32(page)
	} else {
		req.Limit = 10
	}

	resp, err := h.DestinationClient.GetTrendDestinations(ctx, req)
	if err != nil {
		h.Logger.Error("xatolik trend sayohat manzillarini olishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik trend sayohat manzillarini olishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Send Message to User
// @Description Send a message to a user
// @Tags Messages
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param message body communication.SendMessageRequest true "Message"
// @Success 200 {object} communication.SendMessageResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/messages [post]
func (h *Handler) SendMessageUser(ctx *gin.Context) {
	fmt.Println("Keldi")
	fmt.Println("Keldi")
	fmt.Println("Keldi")
	fmt.Println("Keldi")
	var req communication.SendMessageRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", slog.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}

	resp, err := h.CommunityClient.SendMessageUser(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik xabarni yuborishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik xabarni yuborishda",
		})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// @Summary List Messages
// @Description List all messages
// @Tags Messages
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page limit"
// @Success 200 {object} communication.ListMessageResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/messages [get]
func (h *Handler) ListMessage(ctx *gin.Context) {
	var req communication.ListMessageRequest
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

	// Get 'limit' query parameter and convert it to int32
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
	fmt.Println("Keldi")
	resp, err := h.CommunityClient.ListMessage(ctx, &req)
	if err != nil {
		fmt.Println("error", err)
		h.Logger.Error("xatolik xabarlarni ro'yxatlashda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik xabarlarni ro'yxatlashda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Add Travel Tips
// @Description Add new travel tips
// @Tags TravelTips
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param tips body communication.AddTravelTipsRequest true "Travel Tips"
// @Success 200 {object} communication.AddTravelTipsResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/travel-tips [post]
func (h *Handler) AddTravelTips(ctx *gin.Context) {
	fmt.Println("hello")
	var req communication.AddTravelTipsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", slog.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}

	resp, err := h.CommunityClient.AddTravelTips(ctx, &req)
	if err != nil {
		fmt.Println(err)
		h.Logger.Error("xatolik sayohat tavsiyalarini qo'shishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat tavsiyalarini qo'shishda",
		})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// @Summary Get Travel Tips
// @Description Get travel tips by ID
// @Tags TravelTips
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page limit"
// @Param category query string true "Category"
// @Success 200 {object} communication.GetTravelTipsResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/travel-tips/ [get]
func (h *Handler) GetTravelTips(ctx *gin.Context) {
	fmt.Println("hello")
	fmt.Println("hello")

	var req communication.GetTravelTipsRequest
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
	fmt.Println("Keldi")
	fmt.Println("Keldi")
	fmt.Println("Keldi")
	fmt.Println("Keldi")

	resp, err := h.CommunityClient.GetTravelTips(ctx, &req)
	if err != nil {
		fmt.Println(err)
		h.Logger.Error("xatolik sayohat tavsiyalarini olishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat tavsiyalarini olishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get User Statistics
// @Description Get user statistics
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} communication.GetUserStaticsResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/users/{userId}/statics [get]
func (h *Handler) GetUserStatics(ctx *gin.Context) {
	userId := ctx.Param("userId")
	req := &communication.GetUserStaticsRequest{UserId: userId}

	resp, err := h.CommunityClient.GetUserStatics(ctx, req)
	if err != nil {
		h.Logger.Error("xatolik foydalanuvchi statistikasini olishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik foydalanuvchi statistikasini olishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
