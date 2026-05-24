package http

import (
	"PetProject/internal/services"
	"PetProject/internal/services/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Owner struct {
	service *services.Owner
}

func NewOwnerHandler(service *services.Owner) *Owner {
	return &Owner{service: service}
}

func (h *Owner) RegisterOwnerHandler(echo *echo.Echo) {
	api := echo.Group("/api")
	api.POST("/owners", h.Create)
	api.GET("/owners", h.GetAll)
	api.GET("/owners/:id", h.Get)
	api.PUT("/owners", h.Update)
	api.DELETE("/owners/:id", h.Delete)
}

func (h *Owner) Create(c echo.Context) error {
	var request models.CreateOwnerRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := h.service.Create(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (h *Owner) Get(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	response, err := h.service.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if response == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "owner not found"})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Owner) GetAll(c echo.Context) error {
	owners, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, owners)
}

func (h *Owner) Update(c echo.Context) error {
	var request models.UpdateOwnerRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := h.service.Update(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "updated"})
}

func (h *Owner) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	err = h.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
