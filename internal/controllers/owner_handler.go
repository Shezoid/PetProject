package controllers

import (
	"PetProject/internal/services"
	"PetProject/internal/services/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OwnerHandler struct {
	service *services.OwnerService
}

func NewOwnerHandler(service *services.OwnerService) *OwnerHandler {
	return &OwnerHandler{service: service}
}

func (handler *OwnerHandler) RegisterOwnerHandler(echo *echo.Echo) {
	api := echo.Group("/api")
	api.POST("/owners", handler.Create)
	api.GET("/owners", handler.GetAll)
	api.GET("/owners/:id", handler.Get)
	api.PUT("/owners", handler.Update)
	api.DELETE("/owners/:id", handler.Delete)
}

func (handler *OwnerHandler) Create(context echo.Context) error {
	var request models.CreateOwnerRequest
	if err := context.Bind(&request); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := handler.service.CreateOwner(&request); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return context.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (handler *OwnerHandler) Get(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	response := handler.service.GetOwner(id)
	if response == nil {
		return context.JSON(http.StatusNotFound, map[string]string{"error": "owner not found"})
	}
	return context.JSON(http.StatusOK, response)
}

func (handler *OwnerHandler) GetAll(context echo.Context) error {
	owners := handler.service.GetOwners()
	return context.JSON(http.StatusOK, owners)
}

func (handler *OwnerHandler) Update(context echo.Context) error {
	var request models.UpdateOwnerRequest
	if err := context.Bind(&request); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := handler.service.UpdateOwner(&request); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return context.JSON(http.StatusOK, map[string]string{"status": "updated"})
}

func (handler *OwnerHandler) Delete(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	handler.service.DeleteOwner(id)
	return context.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
