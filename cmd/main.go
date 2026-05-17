package main

import (
	"PetProject/internal/controllers"
	"PetProject/internal/infrastructure"
	"PetProject/internal/services"
	"PetProject/pkg"

	"github.com/labstack/echo/v4"
)

func main() {
	repository := pkg.NewRepository("postgres://admin:admin@haproxy:5432/$pet_db?sslmode=disable")
	ownerRepository := infrastructure.NewOwnerRepository(repository)
	petRepository := infrastructure.NewPetRepository(repository)

	ownerService := services.NewOwnerService(ownerRepository)
	petService := services.NewPetService(petRepository)

	ownerhandler := controllers.NewOwnerHandler(ownerService)
	pethandler := controllers.NewPetHandler(petService)

	e := echo.New()
	ownerhandler.RegisterOwnerHandler(e)
	pethandler.RegisterPetHandler(e)
	err := e.Start("8080")
	if err != nil {
		panic(err)
	}
}
