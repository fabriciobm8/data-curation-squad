package controllers

import (
    "data-curation-squad/service"
    "github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, classMaterialService *service.ClassMaterialService, keywordService *service.KeywordService) {
	
	// Inicializando os controladores
    classMaterialController := NewClassMaterialController(classMaterialService)
    keywordController := NewKeywordController(keywordService)

    // Rotas para ClassMaterial
    e.POST("/class-material", classMaterialController.Create)
    e.GET("/class-material", classMaterialController.FindAll)
    e.GET("/class-material/:id",classMaterialController.FindByID)
    e.PATCH("/class-material/transcript-time/:id", classMaterialController.UpdateTranscriptTime)
    //e.PUT("/class-material/:id", classMaterialController.Update)
    //e.DELETE("/class-material/:id", classMaterialController.Delete)
    //e.PATCH("/class-material/update-isSuccessful/:id", classMaterialController.UpdateIsSuccessfulClassMaterial)

    
    // Rotas para Keyword
    e.POST("/keywords", keywordController.SaveKeywords)
    e.GET("/keyword", keywordController.FindAll)
    e.GET("/keyword/:id", keywordController.FindByID)
    //e.PUT("/keyword/:id", keywordController.Update)
    //e.DELETE("/keyword/:id", keywordController.Delete)
    //e.PATCH("/keywords/update-by-transcript-time-id", keywordController.UpdateKeywordsByTranscriptTimeID)
}
