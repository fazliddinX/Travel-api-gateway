package main

import (
	"api-gateway/api"
	"api-gateway/api/handler"
	"api-gateway/config"
	"api-gateway/generated/communication"
	"api-gateway/generated/destination"
	"api-gateway/generated/itineraries"
	"api-gateway/generated/stories"
	"api-gateway/generated/user"
	"api-gateway/logs"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
)

func main() {
	cfg := config.Load()
	logs.InitLogger()

	//service, err := service.NewServiceManager(&cfg)
	//if err != nil {
	//	logs.Logger.Error("gRPC dial error", slog.String("error", err.Error()))
	//}
	//
	//handler := handler.NewHandler(service, logs.Logger)

	connUser, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	connStory, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}

	hd := &handler.Handler{
		UserClient:        user.NewAuthServiceClient(connUser),
		StoriesClient:     stories.NewTravelStoriesServiceClient(connStory),
		ItineraryClient:   itineraries.NewItinerariesServiceClient(connStory),
		CommunityClient:   communication.NewCommunicationServiceClient(connStory),
		DestinationClient: destination.NewTravelDestinationServiceClient(connStory),
		Logger:            logs.Logger,
	}

	router := api.NewRouter(hd)
	logs.Logger.Info("Server is running ... ", slog.String("PORT", cfg.HTTP_PORT))

	err = router.Run(":8080")
	if err != nil {
		logs.Logger.Error("Routerni run qilishda xatolik beryapti", slog.String("error", err.Error()))
		log.Fatal(err)
	}
}
