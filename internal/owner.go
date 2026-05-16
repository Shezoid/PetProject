package internal

import (
	"time"
)

type Owner struct {
	id        int
	name      string
	birthDate time.Time
	petId     []int
}
