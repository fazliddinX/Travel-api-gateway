package main

import (
	"Api-Gateway/api"
	"Api-Gateway/api/handler"
	"Api-Gateway/logger"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	logg := logger.InitLogger()

	content, err := grpc.NewClient(":5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logg.Error("new grpc client error: %v", err)
		log.Fatalln(err)
	}

	user, err := grpc.NewClient(":5052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logg.Error("new grpc client error: %v", err)
		log.Fatalln(err)
	}

	hnd := handler.NewHandler(user, content, logg)
	run := api.Router(hnd)

	err = run.Run(":8080")
	fmt.Println(err)
}
