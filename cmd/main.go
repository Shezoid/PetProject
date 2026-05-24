package main

import (
	"PetProject/internal/http"
	"PetProject/internal/infrastructure"
	"PetProject/internal/services"
	"PetProject/pkg"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	dBPort := "5432"
	dbName := "pet_db"
	dbUser := "admin"
	dbPassword := "admin"
	dbHost := "localhost"
	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dBPort + "/" + dbName + "?sslmode=disable"

	println(url)

	repository := pkg.NewRepository(url)
	ownerRepository := infrastructure.NewOwner(repository)
	petRepository := infrastructure.NewPet(repository)

	ownerService := services.NewOwner(ownerRepository)
	petService := services.NewPet(petRepository)

	ownerHandler := http.NewOwnerHandler(ownerService)
	petHandler := http.NewPetHandler(petService)

	appPort := os.Getenv("APPLICATION_PORT")
	e := echo.New()
	ownerHandler.RegisterOwnerHandler(e)
	petHandler.RegisterPetHandler(e)
	err := e.Start(":" + appPort)
	if err != nil {
		panic(err)
	}
}
