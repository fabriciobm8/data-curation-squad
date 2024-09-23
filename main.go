package main

import (
	"context"
	"data-curation-squad/controllers"
	"data-curation-squad/repository"
	"data-curation-squad/service"
	"log"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    e := echo.New()

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.Background())

    classMaterialRepo := repository.NewClassMaterialRepository(client)
    keywordRepo := repository.NewKeywordRepository(client)

    keywordService := service.NewKeywordService(keywordRepo, nil)
    classMaterialService := service.NewClassMaterialService(classMaterialRepo, keywordService)

    keywordService.ClassMaterialService = classMaterialService

    controllers.RegisterRoutes(e, classMaterialService, keywordService)

    e.Logger.Fatal(e.Start(":8080"))
}