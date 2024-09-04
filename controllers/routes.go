package controllers

import (
    "data-curation-squad/service"
    "github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, classMaterialService *service.ClassMaterialService, transcriptTimeService *service.TranscriptTimeService,
	keywordService *service.KeywordService) {
	
	// Inicializando os controladores
    classMaterialController := NewClassMaterialController(classMaterialService)
    transcriptTimeController := NewTranscriptTimeController(transcriptTimeService)
    keywordController := NewKeywordController(keywordService)

    // Rotas para ClassMaterial
    e.POST("/class-material", classMaterialController.Create)
    e.GET("/class-material", classMaterialController.FindAll)
    e.GET("/class-material/:id",classMaterialController.FindByID)
    //e.PUT("/class-material/:id", classMaterialController.Update)
    //e.DELETE("/class-material/:id", classMaterialController.Delete)
    //e.PATCH("/class-material/update-isSuccessful/:id", classMaterialController.UpdateIsSuccessfulClassMaterial)

    // Rotas para TranscriptTime
    e.POST("/transcript-time", transcriptTimeController.Create)
    e.GET("/transcript-time", transcriptTimeController.FindAll)
    e.GET("/transcript-time/:id", transcriptTimeController.FindByID)
    //e.PUT("/transcript-time/:id", transcriptTimeController.Update)
    //e.DELETE("/transcript-time/:id", transcriptTimeController.Delete)
    //e.PATCH("/transcript-time/:id", transcriptTimeController.UpdateStartEndTimeTranscriptTime)
    //e.PATCH("/transcript-time/transcripts", transcriptTimeController.UpdateTranscripts)

    // Rotas para Keyword
    e.POST("/keyword", keywordController.Create)
    e.GET("/keyword", keywordController.FindAll)
    e.GET("/keyword/:id", keywordController.FindByID)
    //e.PUT("/keyword/:id", keywordController.Update)
    //e.DELETE("/keyword/:id", keywordController.Delete)
    //e.PATCH("/keywords/update-by-transcript-time-id", keywordController.UpdateKeywordsByTranscriptTimeID)
}
