package models

import (
	"errors"
	"time"
)

type CreatePetRequest struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	OwnerID   int       `json:"owner_id"`
	FriendIDs []int     `json:"friend_ids,omitempty"`
}

type UpdatePetRequest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	OwnerID   int       `json:"ownerId"`
	FriendIDs []int     `json:"friendIds,omitempty"`
}

func (r *CreatePetRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Color == "" {
		return errors.New("color is required")
	}
	if r.BirthDate.IsZero() {
		return errors.New("birthDate is required")
	}
	return nil
}

func (r *UpdatePetRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Color == "" {
		return errors.New("color is required")
	}
	if r.BirthDate.IsZero() {
		return errors.New("birthDate is required")
	}
	return nil
}
