package main

import (
	"log"

	"github.com/isido5ik/RecipePublishingPlatform/configs"
	"github.com/isido5ik/RecipePublishingPlatform/internal/delivery/http"
	"github.com/isido5ik/RecipePublishingPlatform/internal/repository"
	"github.com/isido5ik/RecipePublishingPlatform/internal/usecase"
	"github.com/isido5ik/RecipePublishingPlatform/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

//@title Chef-In
//@version 1.0
//@description Api Server for NFactorial

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
