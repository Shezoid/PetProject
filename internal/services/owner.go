package services

import (
	"PetProject/internal/infrastructure"
	"PetProject/internal/services/models"
)

type Owner struct {
	repository *infrastructure.OwnerRepository
}

func NewOwner(repository *infrastructure.OwnerRepository) *Owner {
	return &Owner{repository: repository}
}

func (s *Owner) Create(request *models.CreateOwnerRequest) error {
	if err := request.Validate(); err != nil {
		return err
	}
	owner := &infrastructure.Owner{
		Name:      request.Name,
		BirthDate: request.BirthDate,
		PetIDs:    request.PetIDs}
	err := s.repository.Create(owner)
	return err
}

func (s *Owner) GetByID(id int) (*models.GetOwnerResponse, error) {
	owner, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &models.GetOwnerResponse{
		ID:        owner.ID,
		Name:      owner.Name,
		BirthDate: owner.BirthDate,
		PetIDs:    owner.PetIDs}
	return response, nil
}

func (s *Owner) GetAll() (response []models.GetOwnerResponse, err error) {
	Owners, err := s.repository.FindAll()
	if err != nil {
		return response, err
	}
	for _, owner := range Owners {
		model := models.GetOwnerResponse{
			ID:        owner.ID,
			Name:      owner.Name,
			BirthDate: owner.BirthDate,
			PetIDs:    owner.PetIDs}
		response = append(response, model)
	}
	return response, nil
}

func (s *Owner) Update(request *models.UpdateOwnerRequest) error {
	if err := request.Validate(); err != nil {
		return err
	}

	owner := &infrastructure.Owner{
		ID:        request.ID,
		Name:      request.Name,
		BirthDate: request.BirthDate,
		PetIDs:    request.PetIDs}

	err := s.repository.Update(owner)
	return err
}

func (s *Owner) Delete(id int) error {
	err := s.repository.Delete(id)
	return err
}
