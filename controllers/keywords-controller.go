package controllers

import (
	"context"
	"data-curation-squad/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type KeywordController struct {
	service *service.KeywordService
}

func NewKeywordController(service *service.KeywordService) *KeywordController {
	return &KeywordController{service: service}
}

type KeywordRequest struct {
	ClassMaterialID string   `json:"classMaterialId"`
	Keywords        []string `json:"keywords"`
}

func (c *KeywordController) FindAll(ctx echo.Context) error {
	keywords, err := c.service.FindAll(context.Background())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, keywords)
}

func (c *KeywordController) FindByID(ctx echo.Context) error {
	id := ctx.Param("id")
	keyword, err := c.service.FindByID(context.Background(), id)
	if err != nil {
		if err.Error() == "id é obrigatório" {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		if err.Error() == "keyword não encontrado" {
			return ctx.JSON(http.StatusNotFound, err.Error())
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, keyword)
}

func (c *KeywordController) SaveKeywords(ctx echo.Context) error {
	var request KeywordRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if request.ClassMaterialID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "class material ID is required"})
	}

	err := c.service.SaveKeywords(context.Background(), request.ClassMaterialID, request.Keywords)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not save keywords"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "keywords saved successfully"})
}