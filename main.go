package main

import (
	"context"
	"data-curation-squad/controllers"
	"data-curation-squad/repository"
	"data-curation-squad/service"

	//"data-curation-squad/repository"
	//"data-curation-squad/service"
	//"data-curation-squad/controllers"
	"log"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    e := echo.New()

    // MongoDB setup
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Initialize repositories and services
    classMaterialRepo := repository.NewClassMaterialRepository(client)
    classMaterialService := service.NewClassMaterialService(classMaterialRepo)
    transcriptTimeRepo := repository.NewTranscriptTimeRepository(client)
    transcriptTimeService := service.NewTranscriptTimeService(transcriptTimeRepo)
    keywordRepo := repository.NewKeywordRepository(client)
    keywordService := service.NewKeywordService(keywordRepo)

    // Registrar rotas
	controllers.RegisterRoutes(e, classMaterialService, transcriptTimeService, keywordService)
	

    e.Logger.Fatal(e.Start(":8080"))
}
