package models

import "time"

type GetOwnerResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	PetId     []int     `json:"petId,omitempty"`
}
