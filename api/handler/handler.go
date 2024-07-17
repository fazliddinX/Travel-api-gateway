package handler

import (
	"Api-Gateway/genproto/auth_service"
	"Api-Gateway/genproto/content_service"
	"google.golang.org/grpc"
	"log/slog"
)

type Handler struct {
	Logger  *slog.Logger
	User    auth_service.AuthServiceClient
	Content content_service.ContentClient
}

func NewHandler(user, content *grpc.ClientConn, logger *slog.Logger) *Handler {
	us := auth_service.NewAuthServiceClient(user)
	c := content_service.NewContentClient(content)
	return &Handler{Logger: logger, User: us, Content: c}
}
