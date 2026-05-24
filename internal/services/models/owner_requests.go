package models

import (
	"errors"
	"time"
)

type CreateOwnerRequest struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	PetIDs    []int     `json:"pet_ids,omitempty"`
}

type UpdateOwnerRequest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	PetIDs    []int     `json:"pet_ids,omitempty"`
}

func (r *CreateOwnerRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.BirthDate.IsZero() {
		return errors.New("birthDate is required")
	}
	return nil
}

func (r *UpdateOwnerRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.BirthDate.IsZero() {
		return errors.New("birthDate is required")
	}
	return nil
}
