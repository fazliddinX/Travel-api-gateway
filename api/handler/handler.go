package handler

import (
	"Api-Gateway/genproto/auth_service"
	"Api-Gateway/genproto/content-service"
)

type Handler struct {
	User *auth_service.AuthServiceClient
	Content *content_service.ContentClient
}