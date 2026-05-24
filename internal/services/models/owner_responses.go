package models

import "time"

type GetOwnerResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	PetIDs    []int     `json:"petIds,omitempty"`
}
