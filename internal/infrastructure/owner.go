package infrastructure

import (
	"time"
)

type Owner struct {
	ID        int
	Name      string
	BirthDate time.Time
	PetIDs    []int
}
