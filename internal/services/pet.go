package services

import (
	"PetProject/internal/infrastructure"
	"PetProject/internal/services/models"
)

type Pet struct {
	repository *infrastructure.PetRepository
}

func NewPet(repository *infrastructure.PetRepository) *Pet {
	return &Pet{repository: repository}
}

func (s *Pet) Create(request *models.CreatePetRequest) error {
	if err := request.Validate(); err != nil {
		return err
	}
	pet := &infrastructure.Pet{
		Name:      request.Name,
		BirthDate: request.BirthDate,
		Breed:     request.Breed,
		Color:     request.Color,
		OwnerID:   request.OwnerID,
		FriendIDs: request.FriendIDs}
	err := s.repository.Create(pet)
	return err
}

func (s *Pet) GetByID(id int) (*models.GetPetResponse, error) {
	pet, err := s.repository.FindByID(id)
	response := &models.GetPetResponse{
		ID:        pet.ID,
		Name:      pet.Name,
		BirthDate: pet.BirthDate,
		Breed:     pet.Breed,
		Color:     pet.Color,
		OwnerID:   pet.OwnerID,
		FriendIDs: pet.FriendIDs}
	return response, err
}

func (s *Pet) GetAll() ([]models.GetPetResponse, error) {
	pets, err := s.repository.FindAllPet()
	response := make([]models.GetPetResponse, 0, len(pets))
	for _, pet := range pets {
		model := models.GetPetResponse{
			ID:        pet.ID,
			Name:      pet.Name,
			BirthDate: pet.BirthDate,
			Breed:     pet.Breed,
			Color:     pet.Color,
			OwnerID:   pet.OwnerID,
			FriendIDs: pet.FriendIDs}
		response = append(response, model)
	}
	return response, err
}

func (s *Pet) Update(request *models.UpdatePetRequest) error {
	if err := request.Validate(); err != nil {
		return err
	}
	pet := &infrastructure.Pet{
		ID:        request.ID,
		Name:      request.Name,
		BirthDate: request.BirthDate,
		Breed:     request.Breed,
		Color:     request.Color,
		OwnerID:   request.OwnerID,
		FriendIDs: request.FriendIDs}

	err := s.repository.Update(pet)
	return err
}

func (s *Pet) Delete(id int) error {
	err := s.repository.Delete(id)
	return err
}
