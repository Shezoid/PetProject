package services

import (
	"PetProject/internal/infrastructure"
	"PetProject/internal/services/models"
	"errors"
)

type PetService struct {
	repository *infrastructure.PetRepository
	counter    int
}

func NewPetService(repository *infrastructure.PetRepository) *PetService {
	return &PetService{repository: repository, counter: 0}
}

func (service *PetService) CreatePet(request *models.CreatePetRequest) error {
	if request.Validate() {
		return errors.New("invalid pet")
	}
	pet := &infrastructure.Pet{
		Id:        service.counter,
		Name:      request.Name,
		BirthDate: request.BirthDate,
		Breed:     request.Breed,
		Color:     request.Color,
		OwnerId:   request.OwnerId,
		FriendIds: request.FriendIds}
	service.counter++
	service.repository.SavePet(pet)
	return nil
}

func (service *PetService) GetPet(id int) *models.GetPetResponse {
	pet := service.repository.FindPetById(id)
	response := &models.GetPetResponse{
		Id:        pet.Id,
		Name:      pet.Name,
		BirthDate: pet.BirthDate,
		Breed:     pet.Breed,
		Color:     pet.Color,
		OwnerId:   pet.OwnerId,
		FriendIds: pet.FriendIds}
	return response
}

func (service *PetService) GetPets() (response []models.GetPetResponse) {
	pets := service.repository.FindAllPet()
	for _, pet := range pets {
		model := models.GetPetResponse{
			Id:        pet.Id,
			Name:      pet.Name,
			BirthDate: pet.BirthDate,
			Breed:     pet.Breed,
			Color:     pet.Color,
			OwnerId:   pet.OwnerId,
			FriendIds: pet.FriendIds}
		response = append(response, model)
	}
	return response
}

func (service *PetService) UpdatePet(request *models.UpdatePetRequest) error {
	if request.Validate() {
		return errors.New("invalid pet")
	}
	pet := &infrastructure.Pet{
		Id:        request.Id,
		Name:      request.Name,
		BirthDate: request.BirthDate,
		Breed:     request.Breed,
		Color:     request.Color,
		OwnerId:   request.OwnerId,
		FriendIds: request.FriendIds}

	service.repository.SavePet(pet)
	return nil
}

func (service *PetService) DeletePet(id int) {
	service.repository.Delete(id)
}
