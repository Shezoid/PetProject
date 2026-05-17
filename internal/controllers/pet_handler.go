package controllers

import (
	"PetProject/internal/services"
	"PetProject/internal/services/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PetHandler struct {
	service *services.PetService
}

func NewPetHandler(service *services.PetService) *PetHandler {
	return &PetHandler{service: service}
}

func (handler *PetHandler) RegisterPetHandler(echo *echo.Echo) {
	api := echo.Group("/api")
	api.POST("/pets", handler.Create)
	api.GET("/pets", handler.GetAll)
	api.GET("/pets/:id", handler.Get)
	api.PUT("/pets", handler.Update)
	api.DELETE("/pets/:id", handler.Delete)
}

func (handler *PetHandler) Create(context echo.Context) error {
	var request models.CreatePetRequest
	if err := context.Bind(&request); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := handler.service.CreatePet(&request); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return context.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (handler *PetHandler) Get(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	response := handler.service.GetPet(id)
	if response == nil {
		return context.JSON(http.StatusNotFound, map[string]string{"error": "Pet not found"})
	}
	return context.JSON(http.StatusOK, response)
}

func (handler *PetHandler) GetAll(context echo.Context) error {
	Pets := handler.service.GetPets()
	return context.JSON(http.StatusOK, Pets)
}

func (handler *PetHandler) Update(context echo.Context) error {
	var request models.UpdatePetRequest
	if err := context.Bind(&request); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := handler.service.UpdatePet(&request); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return context.JSON(http.StatusOK, map[string]string{"status": "updated"})
}

func (handler *PetHandler) Delete(context echo.Context) error {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	handler.service.DeletePet(id)
	return context.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
