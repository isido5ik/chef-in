package main

import (
	"log"

	"github.com/isido5ik/StoryPublishingPlatform/configs"
	"github.com/isido5ik/StoryPublishingPlatform/internal/delivery/http"
	"github.com/isido5ik/StoryPublishingPlatform/internal/repository"
	"github.com/isido5ik/StoryPublishingPlatform/internal/usecase"
	"github.com/isido5ik/StoryPublishingPlatform/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := configs.InitConfigs(); err != nil {
		logrus.Fatal("error initializing configs", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db := server.InitDB()
	repo := repository.NewRepository(db)
	usecases := usecase.NewUsecase(repo)
	handlers := http.NewHandler(usecases)

	server := server.NewServer()
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
