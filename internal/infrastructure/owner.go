package infrastructure

import (
	"time"
)

type Owner struct {
	Id        int
	Name      string
	BirthDate time.Time
	PetId     []int
}
