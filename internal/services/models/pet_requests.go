package models

import (
	"time"
)

type CreatePetRequest struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	OwnerId   int       `json:"ownerId"`
	FriendIds []int     `json:"petId,omitempty"`
}

type UpdatePetRequest struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	OwnerId   int       `json:"ownerId"`
	FriendIds []int     `json:"petId,omitempty"`
}

func (request *CreatePetRequest) Validate() bool {
	return request.Color == "" ||
		request.Name == "" ||
		request.Breed == "" ||
		request.BirthDate.IsZero() ||
		request.OwnerId == 0
}

func (request *UpdatePetRequest) Validate() bool {
	return request.Color == "" ||
		request.Name == "" ||
		request.Breed == "" ||
		request.BirthDate.IsZero() ||
		request.OwnerId == 0
}
