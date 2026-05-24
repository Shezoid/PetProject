package http

import (
	"PetProject/internal/services"
	"PetProject/internal/services/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Pet struct {
	service *services.Pet
}

func NewPetHandler(service *services.Pet) *Pet {
	return &Pet{service: service}
}

func (h *Pet) RegisterPetHandler(echo *echo.Echo) {
	api := echo.Group("/api")
	api.POST("/pets", h.Create)
	api.GET("/pets", h.GetAll)
	api.GET("/pets/:id", h.Get)
	api.PUT("/pets", h.Update)
	api.DELETE("/pets/:id", h.Delete)
}

func (h *Pet) Create(c echo.Context) error {
	var request models.CreatePetRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := h.service.Create(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (h *Pet) Get(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	response, err := h.service.GetByID(id)
	if response == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Pet not found"})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Pet) GetAll(c echo.Context) error {
	Pets, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, Pets)
}

func (h *Pet) Update(c echo.Context) error {
	var request models.UpdatePetRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := h.service.Update(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "updated"})
}

func (h *Pet) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	err = h.service.Delete(id)
	return c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
