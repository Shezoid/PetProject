package models

import (
	"time"
)

type GetPetResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	OwnerId   int       `json:"ownerId"`
	FriendIds []int     `json:"petId,omitempty"`
}
