package models

import "time"

type CreateOwnerRequest struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	PetId     []int     `json:"petId,omitempty"`
}

type UpdateOwnerRequest struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	PetId     []int     `json:"petId,omitempty"`
}

func (request *CreateOwnerRequest) Validate() bool {
	return request.Name == "" ||
		request.BirthDate.IsZero()
}

func (request *UpdateOwnerRequest) Validate() bool {
	return request.Name == "" ||
		request.BirthDate.IsZero()
}
