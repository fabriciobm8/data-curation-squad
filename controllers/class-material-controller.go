package controllers

import (
    "context"
    "data-curation-squad/model"
    "data-curation-squad/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type ClassMaterialController struct {
    service *service.ClassMaterialService
}

func NewClassMaterialController(service *service.ClassMaterialService) *ClassMaterialController {
    return &ClassMaterialController{service: service}
}

func (c *ClassMaterialController) Create(ctx echo.Context) error {
    var classMaterial model.ClassMaterial
    if err := ctx.Bind(&classMaterial); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    err := c.service.Create(context.Background(), &classMaterial)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusCreated, classMaterial)
}

func (c *ClassMaterialController) FindAll(ctx echo.Context) error {
    classMaterials, err := c.service.FindAll(context.Background())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, classMaterials)
}

func (c *ClassMaterialController) FindByID(ctx echo.Context) error {
    id := ctx.Param("id")
    classMaterial, err := c.service.FindByID(context.Background(), id)
    if err != nil {
        if err.Error() == "id é obrigatório" {
            return ctx.JSON(http.StatusBadRequest, err.Error())
        }
        if err.Error() == "classMaterial não encontrado" {
            return ctx.JSON(http.StatusNotFound, err.Error())
        }
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, classMaterial)
}

func (c *ClassMaterialController) UpdateTranscriptTime(ctx echo.Context) error {
    id := ctx.Param("id")
    var transcriptTime []model.TranscriptTime
    if err := ctx.Bind(&transcriptTime); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    err := c.service.UpdateTranscriptTime(context.Background(), id, transcriptTime)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, "TranscriptTime atualizado com sucesso")
}

func (c *ClassMaterialController) GetByCourseId(ctx echo.Context) error {
    courseId := ctx.Param("courseId")
    classMaterials, err := c.service.GetByCourseId(context.Background(), courseId)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, classMaterials)
}

func (c *ClassMaterialController) GetByObjectiveId(ctx echo.Context) error {
    objectiveId := ctx.Param("objectiveId")
    classMaterials, err := c.service.GetByObjectiveId(context.Background(), objectiveId)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, classMaterials)
}

func (c *ClassMaterialController) GetByMaterialId(ctx echo.Context) error {
    materialId := ctx.Param("materialId")
    classMaterials, err := c.service.GetByMaterialId(context.Background(), materialId)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, classMaterials)
}
