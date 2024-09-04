package controllers

import (
    "context"
    "data-curation-squad/model"
    "data-curation-squad/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type TranscriptTimeController struct {
    service *service.TranscriptTimeService
}

func NewTranscriptTimeController(service *service.TranscriptTimeService) *TranscriptTimeController {
    return &TranscriptTimeController{service: service}
}

func (c *TranscriptTimeController) Create(ctx echo.Context) error {
    var transcriptTime model.TranscriptTime
    if err := ctx.Bind(&transcriptTime); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    err := c.service.Create(context.Background(), &transcriptTime)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusCreated, transcriptTime)
}

func (c *TranscriptTimeController) FindAll(ctx echo.Context) error {
    transcriptTimes, err := c.service.FindAll(context.Background())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, transcriptTimes)
}

func (c *TranscriptTimeController) FindByID(ctx echo.Context) error {
    id := ctx.Param("id")
    transcriptTime, err := c.service.FindByID(context.Background(), id)
    if err != nil {
        if err.Error() == "id é obrigatório" {
            return ctx.JSON(http.StatusBadRequest, err.Error())
        }
        if err.Error() == "transcriptTime não encontrado" {
            return ctx.JSON(http.StatusNotFound, err.Error())
        }
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, transcriptTime)
}
