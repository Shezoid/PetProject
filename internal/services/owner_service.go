package services

import (
	"PetProject/internal/infrastructure"
	"PetProject/internal/services/models"
	"errors"
)

type OwnerService struct {
	repository *infrastructure.OwnerRepository
	counter    int
}

func NewOwnerService(repository *infrastructure.OwnerRepository) *OwnerService {
	return &OwnerService{repository: repository, counter: 0}
}

func (service *OwnerService) CreateOwner(request *models.CreateOwnerRequest) error {
	if request.Validate() {
		return errors.New("invalid Owner")
	}
	Owner := &infrastructure.Owner{
		Id:        service.counter,
		Name:      request.Name,
		BirthDate: request.BirthDate,
		PetId:     request.PetId}
	service.counter++
	service.repository.SaveOwner(Owner)
	return nil
}

func (service *OwnerService) GetOwner(id int) *models.GetOwnerResponse {
	Owner := service.repository.FindOwnerById(id)
	response := &models.GetOwnerResponse{
		Id:        Owner.Id,
		Name:      Owner.Name,
		BirthDate: Owner.BirthDate,
		PetId:     Owner.PetId}
	return response
}

func (service *OwnerService) GetOwners() (response []models.GetOwnerResponse) {
	Owners := service.repository.FindAllOwner()
	for _, Owner := range Owners {
		model := models.GetOwnerResponse{
			Id:        Owner.Id,
			Name:      Owner.Name,
			BirthDate: Owner.BirthDate,
			PetId:     Owner.PetId}
		response = append(response, model)
	}
	return response
}

func (service *OwnerService) UpdateOwner(request *models.UpdateOwnerRequest) error {
	if request.Validate() {
		return errors.New("invalid Owner")
	}
	Owner := &infrastructure.Owner{
		Id:        request.Id,
		Name:      request.Name,
		BirthDate: request.BirthDate,
		PetId:     request.PetId}

	service.repository.SaveOwner(Owner)
	return nil
}

func (service *OwnerService) DeleteOwner(id int) {
	service.repository.Delete(id)
}
