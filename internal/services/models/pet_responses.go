package models

import (
	"time"
)

type GetPetResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	OwnerID   int       `json:"ownerId"`
	FriendIDs []int     `json:"friendIds,omitempty"`
}
